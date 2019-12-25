package application

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/rs/cors"

	_ "private-conda-repo/conda/filesys"
	"private-conda-repo/config"
	_ "private-conda-repo/store/postgres"
)

type router struct {
	*chi.Mux
}

func New() (*http.Server, error) {
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not start application server due to issue with config")
	}
	addr := fmt.Sprintf(":%d", conf.AppServer.Port)
	log.WithField("Address", addr).Info("Server details")

	r := router{chi.NewRouter()}
	r.attachMiddleware()
	r.registerRoutes()

	if err := initStore(); err != nil {
		return nil, err
	}

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

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler)
}

func (r *router) registerRoutes() {
	r.Get("/healthcheck", HealthCheck)
	r.Get("/meta", MetaInfo)

	// user routes
	r.Route("/user", func(r chi.Router) {
		r.Get("/", ListUsers)
		r.Post("/", CreateUser)
		r.Delete("/", RemoveUser)
		r.Post("/check", CheckUser)
	})

	// package routes
	r.Route("/p", func(r chi.Router) {
		r.Get("/{user}", ListPackagesByUser)
		r.Get("/{user}/{pkg}", ListPackageDetails)

		r.Post("/", UploadPackage)
		r.Delete("/", RemovePackage)
		r.Delete("/{pkg}", RemoveAllPackages)
	})

	r.Route("/image", func(r chi.Router) {
		r.Get("/", GetImageManagerVersion)
		r.Post("/", UpdateImageManagerVersion)
	})

	r.Route("/", func(r chi.Router) {
		r.Get("/*", missing404Handler)
		r.Put("/*", missing404Handler)
		r.Post("/*", missing404Handler)
		r.Delete("/*", missing404Handler)
		r.Patch("/*", missing404Handler)
	})
}
