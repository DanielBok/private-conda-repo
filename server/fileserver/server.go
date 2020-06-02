package fileserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/api/interfaces"
	"private-conda-repo/config"
)

type MasterHandler struct {
	DB interfaces.DataAccessLayer
	*chi.Mux
}

func New(conf *config.AppConfig, db interfaces.DataAccessLayer) (*http.Server, error) {
	addr := fmt.Sprintf(":%d", conf.FileServer.Port)
	log.WithField("Address", addr).Info("Server details")

	m := MasterHandler{
		Mux: chi.NewRouter(),
		DB:  db,
	}

	m.attachMiddleware()
	m.addFileServer(conf)

	return &http.Server{
		Addr:    addr,
		Handler: m,
	}, nil
}

func (m *MasterHandler) attachMiddleware() {
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)
}

func (m *MasterHandler) addFileServer(conf *config.AppConfig) {
	handler := &FileHandler{
		DB: m.DB,
	}

	m.Get("/*", handler.Server(conf.Indexer.MountFolder))
}
