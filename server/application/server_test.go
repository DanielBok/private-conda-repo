package application

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/go-chi/chi"
	"github.com/spf13/viper"

	_ "private-conda-repo/conda/condamocks"
	"private-conda-repo/conda/condatypes"
	_ "private-conda-repo/store/storemock"
)

const (
	ApplicationJson = "application/json"
)

func init() {
	viper.Set("db.type", "mock")
	viper.Set("conda.type", "mock")
	viper.Set("decompressor.type", "mock")

	if err := initStore(); err != nil {
		log.Fatalln(err)
	}
}

func newTestServer(f http.HandlerFunc) *httptest.Server {
	ts := httptest.NewServer(f)
	return ts
}

func newTestServerWithRouteContext(method, pattern string, f http.HandlerFunc) *httptest.Server {
	m := chi.NewRouter()
	m.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	})

	m.MethodFunc(method, pattern, f)
	ts := httptest.NewServer(m)
	return ts
}

func createChannelAndAddPackages(channel string, packages ...condatypes.Package) error {
	chn, err := repo.GetChannel(channel)
	if err != nil {
		return err
	}

	for _, p := range packages {
		_, err := chn.AddPackage(bytes.NewBufferString("asd"), &p)
		if err != nil {
			return err
		}
	}

	return nil
}
