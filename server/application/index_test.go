package application

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHealthCheck(t *testing.T) {
	assert := require.New(t)
	ts := newTestServer(HealthCheck)
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.EqualValues(200, resp.StatusCode)
}
