package application

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"private-conda-repo/conda"
	_ "private-conda-repo/conda/filesys"
	"private-conda-repo/config"
	"private-conda-repo/store"
	_ "private-conda-repo/store/postgres"
)

var (
	db   store.Store
	repo conda.Conda
)

type router struct {
	*chi.Mux
}

func initStore() error {
	_db, err := store.New()
	if err != nil {
		return err
	}

	_repo, err := conda.New()
	if err != nil {
		return err
	}

	db = _db
	repo = _repo

	return nil
}

func New() (*http.Server, error) {
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, "could not start repository server due to issue with config")
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
}

func (r *router) registerRoutes() {
	r.Get("/HealthCheck", HealthCheck)

	// user routes
	r.Route("/user", func(r chi.Router) {
		r.Get("/", ListUsers)
		r.Post("/", CreateUser)
		r.Delete("/", RemoveUser)
	})

	// package routes
	r.Route("/p", func(r chi.Router) {
		r.Get("/{user}", ListPackagesByUser)
		r.Get("/{user}/{pkg}", ListPackageDetails)

		r.Post("/", UploadPackage)
		r.Delete("/", RemovePackage)
		r.Delete("/{pkg}", RemoveAllPackages)
	})
}

func toJson(w http.ResponseWriter, object interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(object); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func readJson(r *http.Request, object interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(object); err != nil {
		return err
	}
	return nil
}

func ok(w http.ResponseWriter) {
	if _, err := fmt.Fprint(w, "Okay"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
