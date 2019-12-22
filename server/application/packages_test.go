package application

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"private-conda-repo/conda/condatypes"
)

func TestListPackagesByUser(t *testing.T) {
	assert := assert.New(t)

	ts := newTestServerWithRouteContext("GET", "/{user}", ListPackagesByUser)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/pikachu")
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var output []*condatypes.ChannelMetaPackageOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.Len(output, 1) // hard-coded from the mock interface
}

func TestListPackageDetails(t *testing.T) {
	assert := assert.New(t)

	ts := newTestServerWithRouteContext("GET", "/{user}/{pkg}", ListPackageDetails)
	defer ts.Close()

	tests := []struct {
		input       string
		statusCode  int
		expectedLen int
	}{
		{
			input:       "bad-package-name",
			statusCode:  500,
			expectedLen: 0,
		}, {
			input:       "perfana",
			statusCode:  200,
			expectedLen: 1,
		},
	}

	for _, test := range tests {
		resp, err := http.Get(fmt.Sprintf("%s/pikachu/%s", ts.URL, test.input))
		assert.NoError(err)
		assert.EqualValues(test.statusCode, resp.StatusCode)

		if test.statusCode == 200 {
			defer func() { _ = resp.Body.Close() }()
			var output []*condatypes.Package
			err = json.NewDecoder(resp.Body).Decode(&output)
			assert.NoError(err)

			assert.Len(output, test.expectedLen)
		}
	}
}
