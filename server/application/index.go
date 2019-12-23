package application

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	ok(w)
}

func missing404Handler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, fmt.Sprintf("'%s' request for '%s' does not exist", r.Method, r.URL.String()), http.StatusNotFound)
	return
}
