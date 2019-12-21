package application

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	assert := assert.New(t)
	ts := newTestServer(HealthCheck)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if assert.NoError(err) {
		assert.EqualValues(200, resp.StatusCode)
	}
}
