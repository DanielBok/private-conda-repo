package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"private-conda-repo/conda/condatypes"
	"private-conda-repo/testutils"
)

func TestListPackagesByUser(t *testing.T) {
	assert := require.New(t)

	channel := "daniel-list-by-user"
	err := createChannelAndAddPackages(channel, condatypes.Package{
		Name:        "perfana",
		Version:     "0.0.1",
		BuildString: "py",
		BuildNumber: 0,
		Platform:    "noarch",
	})
	assert.NoError(err)

	ts := newTestServerWithRouteContext("GET", "/{user}", ListPackagesByUser)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/%s", ts.URL, channel))
	assert.NoError(err)
	assert.EqualValues(resp.StatusCode, 200)

	var output []*condatypes.ChannelMetaPackageOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	assert.Len(output, 1)
}

func TestListPackageDetails(t *testing.T) {
	type TestRow struct {
		input       string
		statusCode  int
		expectedLen int
	}

	assert := require.New(t)
	channel := "daniel-list"
	err := createChannelAndAddPackages(channel, condatypes.Package{
		Name:        "perfana",
		Version:     "0.0.1",
		BuildString: "py",
		BuildNumber: 0,
		Platform:    "noarch",
	})
	assert.NoError(err)

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
		resp, err := http.Get(fmt.Sprintf("%s/%s/%s", ts.URL, channel, test.input))
		assert.NoError(err)
		assert.EqualValues(test.statusCode, resp.StatusCode)

		if test.statusCode == 200 && resp.StatusCode == 200 {
			defer func() { _ = resp.Body.Close() }()
			var output []*condatypes.Package
			err := json.NewDecoder(resp.Body).Decode(&output)
			assert.NoError(err)
			assert.Len(output, test.expectedLen)
		}
	}

	for _, test := range tests {
		runTest(test)
	}
}

func TestUploadPackage(t *testing.T) {
	assert := require.New(t)

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
	err = json.NewDecoder(resp.Body).Decode(&output)
	assert.NoError(err)
	assert.EqualValues(*pkg.ToPackage(), output)
}

func TestRemovePackage(t *testing.T) {
	assert := require.New(t)

	ts := newTestServer(RemovePackage)
	defer ts.Close()
	_, err := db.AddUser("daniel", "pikachu")
	assert.NoError(err)
	chn, _ := repo.GetChannel("remove-package-channel")
	_, _ = chn.AddPackage(nil, &condatypes.Package{
		Name:        "test-package",
		Version:     "0.1",
		BuildString: "py12345",
		BuildNumber: 0,
		Platform:    "noarch",
	})

	payload := `{
		"channel": "daniel",
		"password": "pikachu",
		"package": {
			"name": "test-package",
			"version": "0.1",
			"build_string": "py12345",
			"build_number": 0,
			"platform": "noarch"
		}
	}`

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", ts.URL, strings.NewReader(payload))
	assert.NoError(err)

	resp, err := client.Do(req)
	assert.NoError(err)
	assert.EqualValues(http.StatusOK, resp.StatusCode)
}
