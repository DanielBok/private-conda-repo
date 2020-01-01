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
	"private-conda-repo/config"
	"private-conda-repo/fileserver"
	"private-conda-repo/indexer"
	_ "private-conda-repo/indexer/docker"
	_ "private-conda-repo/indexer/shell"
)

type App struct {
	conf           *config.AppConfig
	idleConnClosed chan struct{}
}

func NewApp() *App {
	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	return &App{conf: conf}
}

func (a *App) updateIndexer() *App {
	idx, err := indexer.New()
	if err != nil {
		log.Fatalln(err)
	}

	if err := idx.Check(); err != nil {
		log.Fatalln("indexer does not exist. ", err)
	}

	if err := idx.Update(); err != nil {
		log.Fatalln("Could not update indexer")
	}

	return a
}

// Runs the servers. Returns a channel for graceful shutdown
func (a *App) runServers() <-chan struct{} {
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

	a.idleConnClosed = make(chan struct{})
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
		close(a.idleConnClosed)
	}()

	runServer := func(srv *http.Server) {
		tls := a.conf.TLS
		if tls.HasCert() {
			log.Info("Running server in HTTPS mode")
			if err := srv.ListenAndServeTLS(tls.Cert, tls.Key); err != http.ErrServerClosed {
				log.WithField("cause", err).Error("server error")
			}
		} else {
			log.Info("Running server in HTTP mode")
			if err := srv.ListenAndServe(); err != http.ErrServerClosed {
				log.WithField("cause", err).Error("server error")
			}
		}
	}

	go runServer(fileSrv)
	go runServer(appSrv)

	return a.idleConnClosed
}
