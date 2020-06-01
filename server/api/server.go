package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"github.com/rs/cors"

	"private-conda-repo/api/interfaces"
	"private-conda-repo/config"
	_ "private-conda-repo/infrastructure/database/postgres"
)

type MasterHandler struct {
	*chi.Mux
	Config       *config.AppConfig
	db           interfaces.DataAccessLayer
	decompressor interfaces.Decompressor
	fileSys      interfaces.FileSys
}

func New(conf *config.AppConfig, db interfaces.DataAccessLayer, decompressor interfaces.Decompressor, fileSys interfaces.FileSys) (*http.Server, error) {
	addr := fmt.Sprintf(":%d", conf.AppServer.Port)
	log.WithField("Address", addr).Info("Server details")

	r := MasterHandler{
		chi.NewRouter(),
		conf,
		db,
		decompressor,
		fileSys,
	}

	r.attachMiddleware()
	r.registerRoutes()

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}, nil
}

func (m *MasterHandler) attachMiddleware() {
	m.Use(middleware.RequestID)
	m.Use(middleware.RealIP)
	m.Use(middleware.Logger)
	m.Use(middleware.Recoverer)

	m.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler)
}

func (m *MasterHandler) registerRoutes() {
	// channel routes
	m.Route("/channel", func(r chi.Router) {
		h := ChannelHandler{
			DB:      m.db,
			FileSys: m.fileSys,
		}

		r.Get("/", h.ListChannels())
		r.Get("/{channel}", h.GetChannelInfo())

		r.Post("/", h.CreateChannel())
		r.Post("/check", h.CheckChannel())

		r.Delete("/", h.RemoveChannel())
	})

	// package routes
	m.Route("/p", func(r chi.Router) {
		h := PackageHandler{
			DB:           m.db,
			Decompressor: m.decompressor,
			FileSys:      m.fileSys,
		}

		r.Get("/", h.ListAllPackages())
		r.Get("/{channel}", h.ListPackagesInChannel())
		r.Get("/{channel}/{pkg}", h.FetchPackageDetails())

		r.Post("/", h.UploadPackage())
		r.Delete("/", h.RemovePackage())
		r.Delete("/{pkg}", h.RemoveAllPackages())
	})

	// index level routes, these are the least specific
	m.Route("/", func(r chi.Router) {
		h := IndexHandler{Conf: m.Config}

		r.Get("/healthcheck", h.HealthCheck())
		r.Get("/meta", h.MetaInfo())

		r.HandleFunc("/*", h.NotFound())
	})
}
