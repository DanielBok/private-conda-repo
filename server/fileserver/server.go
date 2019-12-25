package fileserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/config"
)

type router struct {
	*chi.Mux
}

func New() (*http.Server, error) {
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not start repository server due to issue with config")
	}
	addr := fmt.Sprintf(":%d", conf.FileServer.Port)
	log.WithField("Address", addr).Info("Server details")

	r := router{chi.NewRouter()}
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
	root := http.Dir(conf.Conda.MountFolder)
	fs := http.FileServer(root)

	log.WithField("Repository mount folder used to serve packages", root).Info()

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}
