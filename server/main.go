package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/application"
	"private-conda-repo/fileserver"
	"private-conda-repo/store"
)

func main() {
	setLogger()
	initStore()
	<-runServers()
}

func setLogger() {
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func initStore() {
	s, err := store.New()
	if err != nil {
		log.Fatalln(err)
	}
	if err := s.Migrate(); err != nil {
		log.Fatalln(err)
	}
}

// Runs the servers. Returns a channel for graceful shutdown
func runServers() <-chan struct{} {
	var _error error
	fileSrv, err := fileserver.New()
	if err != nil {
		_error = multierror.Append(_error, errors.Wrap(err, "Could not create file server"))
	}
	appSrv, err := application.New()
	if err != nil {
		_error = multierror.Append(_error, errors.Wrap(err, "Could not create application server"))
	}

	if _error != nil {
		log.Fatalln(_error)
	}

	idleConnClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		sig := <-sigint

		log.WithField("signal", sig.String()).Info("Shutting down servers")

		if err := fileSrv.Shutdown(context.Background()); err != nil {
			log.WithField("cause", err).Error("error shutting down file server")
		}
		if err := appSrv.Shutdown(context.Background()); err != nil {
			log.WithField("cause", err).Error("error shutting down application server")
		}
		close(idleConnClosed)
	}()

	runServer := func(srv *http.Server) {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.WithField("cause", err).Error("server error")
		}
	}

	go runServer(fileSrv)
	go runServer(appSrv)

	return idleConnClosed
}
