package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/api"
	"private-conda-repo/api/interfaces"
	"private-conda-repo/config"
	"private-conda-repo/fileserver"
	"private-conda-repo/infrastructure/conda/filesys"
	"private-conda-repo/infrastructure/conda/index"
	"private-conda-repo/infrastructure/database/postgres"
	"private-conda-repo/infrastructure/decompressor"
)

type App struct {
	ApiServer      *http.Server
	FileServer     *http.Server
	conf           *config.AppConfig
	idleConnClosed chan struct{}
}

func NewApp() (*App, error) {
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, "error getting config")
	}

	db, err := postgres.New(conf.DB)
	if err != nil {
		return nil, err
	}
	if conf.DB.AutoMigrate {
		err = db.Migrate()
		if err != nil {
			return nil, err
		}
		log.Info("Database is at the latest migration")
	} else {
		log.Info("Migrations were not applied onto the database")
	}

	var indexer interfaces.Indexer
	switch conf.Indexer.Type {
	case "docker":
		indexer, err = index.NewDockerIndex(conf.Indexer.ImageName)
		log.Infof("Using DockerIndex on %s", conf.Indexer.MountFolder)
	case "shell":
		indexer, err = index.NewShellIndex()
		log.Infof("Using ShellIndex on %s", conf.Indexer.MountFolder)
	}
	if err != nil {
		return nil, err
	}

	if conf.Indexer.Update {
		err = indexer.Update()
		if err != nil {
			return nil, err
		}
		log.Info("The conda folder indexer is at the latest version")
	} else {
		log.Info("The conda folder indexer was not updated")
	}

	fs := filesys.New(conf.Indexer.MountFolder, indexer)

	app := &App{
		conf:           conf,
		idleConnClosed: make(chan struct{}),
	}

	app.FileServer, err = fileserver.New(conf, db)
	if err != nil {
		return nil, err
	}

	app.ApiServer, err = api.New(conf, db, decompressor.New(), fs)
	if err != nil {
		return nil, err
	}

	return app, nil
}

// Runs the servers. Returns a channel for graceful shutdown
func (a *App) RunServers() <-chan struct{} {
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt, syscall.SIGTERM)
		sig := <-sigint

		log.WithField("signal", sig.String()).Info("Shutting down servers")

		if err := a.FileServer.Shutdown(context.Background()); err != nil {
			log.WithField("cause", err).Error("error shutting down file server")
		}
		if err := a.ApiServer.Shutdown(context.Background()); err != nil {
			log.WithField("cause", err).Error("error shutting down api server")
		}
		close(a.idleConnClosed)
	}()

	runServer := func(srv *http.Server) {
		tls := a.conf.TLS
		log.Info(tls.HasCert())
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

	go runServer(a.FileServer)
	go runServer(a.ApiServer)

	return a.idleConnClosed
}
