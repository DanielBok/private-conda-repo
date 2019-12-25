package application

import (
	"fmt"
	"net/http"
	"strings"

	"private-conda-repo/config"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	ok(w)
}

func MetaInfo(w http.ResponseWriter, r *http.Request) {
	domain := strings.Split(r.Host, ":")[0]
	schema := "http"
	if r.TLS != nil {
		schema = "https"
	}

	conf, err := config.New()
	if err != nil {
		http.Error(w, "cannot read in application config", http.StatusInternalServerError)
		return
	}

	meta := struct {
		Image      string `json:"image"`
		Registry   string `json:"registry"`
		Repository string `json:"repository"`
	}{
		Image:      conf.Conda.ImageName,
		Registry:   fmt.Sprintf("%s://%s:%d", schema, domain, conf.AppServer.Port),
		Repository: fmt.Sprintf("%s://%s:%d", schema, domain, conf.FileServer.Port),
	}

	toJson(w, &meta)
}

func missing404Handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("'%s' request for '%s' does not exist", r.Method, r.URL.String()), http.StatusNotFound)
	return
}
