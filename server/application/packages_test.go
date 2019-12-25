package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"private-conda-repo/conda/condatypes"
	"private-conda-repo/testutils"
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
	type TestRow struct {
		input       string
		statusCode  int
		expectedLen int
	}

	assert := assert.New(t)

	ts := newTestServerWithRouteContext("GET", "/{user}/{pkg}", ListPackageDetails)
	defer ts.Close()

	tests := []TestRow{
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

	runTest := func(test TestRow) {
		resp, err := http.Get(fmt.Sprintf("%s/pikachu/%s", ts.URL, test.input))
		assert.NoError(err)
		assert.EqualValues(test.statusCode, resp.StatusCode)

		if test.statusCode == 200 && resp.StatusCode == 200 {
			defer func() { _ = resp.Body.Close() }()
			var output []*condatypes.Package
			if err := json.NewDecoder(resp.Body).Decode(&output); assert.NoError(err) {
				assert.Len(output, test.expectedLen)
			}
		}
	}

	for _, test := range tests {
		runTest(test)
	}
}

func TestUploadPackage(t *testing.T) {
	assert := assert.New(t)

	ts := newTestServer(UploadPackage)
	defer ts.Close()
	_, err := db.AddUser("daniel", "pikachu")
	assert.NoError(err)

	formData := func(pkg testutils.TestPackage) (io.ReadWriter, string, error) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		if err := writer.WriteField("channel", "daniel"); err != nil {
			return nil, "", err
		}
		if err := writer.WriteField("password", "pikachu"); err != nil {
			return nil, "", err
		}
		parts, err := writer.CreateFormFile("file", pkg.Filename)
		if err != nil {
			return nil, "", err
		}

		// Write package file into form field
		file, err := os.Open(pkg.Path)
		if err != nil {
			return nil, "", err
		}
		if _, err = io.Copy(parts, file); err != nil {
			return nil, "", err
		}

		if err := writer.Close(); err != nil {
			return nil, "", err
		}

		return body, writer.FormDataContentType(), nil
	}

	pkg, err := testutils.GetPackageN(0)
	assert.NoError(err)
	body, contentType, err := formData(pkg)
	assert.NoError(err)

	resp, err := http.Post(ts.URL, contentType, body)
	assert.NoError(err)
	assert.EqualValues(200, resp.StatusCode)
	defer func() { _ = resp.Body.Close() }()

	var output condatypes.Package
	if err = json.NewDecoder(resp.Body).Decode(&output); assert.NoError(err) {
		assert.EqualValues(*pkg.ToPackage(), output)
	}
}
