package application

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/spf13/viper"

	_ "private-conda-repo/conda/condamocks"
	_ "private-conda-repo/store/storemock"
)

const (
	ApplicationJson = "application/json"
)

func init() {
	viper.Set("db.type", "mock")
	viper.Set("conda.type", "mock")

	if err := initStore(); err != nil {
		log.Fatalln(err)
	}
}

func newTestServer(f http.HandlerFunc) *httptest.Server {
	ts := httptest.NewServer(f)
	return ts
}
