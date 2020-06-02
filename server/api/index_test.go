package api_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/api"
	"private-conda-repo/api/dto"
	"private-conda-repo/config"
)

func NewIndexHandler() *IndexHandler {
	return &IndexHandler{
		Conf: &config.AppConfig{
			Admin: nil,
			Indexer: &config.IndexerConfig{
				Type:      "shell",
				ImageName: "",
			},
			DB:         nil,
			FileServer: &config.ServerConfig{Port: 5050},
			AppServer:  &config.ServerConfig{Port: 5060},
		},
	}
}

func TestIndexHandler_HealthCheck(t *testing.T) {
	assert := require.New(t)
	handler := NewIndexHandler()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/healthcheck", nil)

	handler.HealthCheck()(w, r)
	assert.Equal(http.StatusOK, w.Code)
}

func TestIndexHandler_MetaInfo(t *testing.T) {
	assert := require.New(t)
	handler := NewIndexHandler()

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/meta", nil)

	handler.MetaInfo()(w, r)
	assert.Equal(http.StatusOK, w.Code)

	var meta dto.ApiMetaInfo
	err := json.NewDecoder(w.Body).Decode(&meta)
	assert.NoError(err)

	formURL := func(port int) string {
		return fmt.Sprintf("http://%s:%d", r.Host, port)
	}

	assert.Equal(dto.ApiMetaInfo{
		Indexer:    "shell",
		Image:      "",
		Registry:   formURL(5060),
		Repository: formURL(5050),
	}, meta)
}
