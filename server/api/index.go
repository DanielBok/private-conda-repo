package api

import (
	"fmt"
	"net/http"
	"strings"

	"private-conda-repo/api/dto"
	"private-conda-repo/config"
)

type IndexHandler struct {
	Conf *config.AppConfig
}

func (h *IndexHandler) HealthCheck() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok(w)
	}
}

func (h *IndexHandler) MetaInfo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := strings.Split(r.Host, ":")[0]
		schema := "http"
		if r.TLS != nil {
			schema = "https"
		}

		meta := dto.ApiMetaInfo{
			Indexer:    h.Conf.Indexer.Type,
			Image:      h.Conf.Indexer.ImageName,
			Registry:   fmt.Sprintf("%s://%s:%d", schema, domain, h.Conf.AppServer.Port),
			Repository: fmt.Sprintf("%s://%s:%d", schema, domain, h.Conf.FileServer.Port),
		}

		toJson(w, &meta)
	}
}

func (h *IndexHandler) NotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("'%s' request for '%s' does not exist", r.Method, r.URL.String()), http.StatusNotFound)
	}
}
