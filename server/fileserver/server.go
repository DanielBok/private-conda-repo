package fileserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
)

type router struct {
	*chi.Mux
}

func New(conf *config.AppConfig) (*http.Server, error) {
	addr := fmt.Sprintf(":%d", conf.FileServer.Port)
	log.WithField("Address", addr).Info("Server details")

	if err := initStore(conf); err != nil {
		return nil, err
	}

	r := router{
		Mux: chi.NewRouter(),
	}

	r.attachMiddleware()
	r.addFileServer(conf)

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}, nil
}

func (r *router) attachMiddleware() {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
}

func (r *router) addFileServer(conf *config.AppConfig) {
	r.Get("/*", FileHandler(conf.Conda.MountFolder))
}
