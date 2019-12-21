package application

import (
	"net/http"
)

func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	ok(w)
}
