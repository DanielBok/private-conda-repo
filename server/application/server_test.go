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
	"private-conda-repo/config"
	_ "private-conda-repo/store/storemock"
)

const (
	ApplicationJson = "application/json"
)

func init() {
	viper.Set("db.type", "mock")
	viper.Set("conda.type", "mock")
	viper.Set("decompressor.type", "mock")

	conf, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}
	if err := initStore(conf); err != nil {
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
	_, err := repo.CreateChannel(channel)
	if err != nil {
		return err
	}

	chn, err := repo.GetChannel(channel)
	if err != nil {
		return err
	}

	for _, p := range packages {
		_, err := chn.AddPackage(bytes.NewBufferString(""), &p)
		if err != nil {
			return err
		}
		_, err = db.CreatePackageCount(p.ToPackageCount(channel))
		if err != nil {
			return err
		}
	}

	return nil
}
