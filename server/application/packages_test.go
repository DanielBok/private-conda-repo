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
	"private-conda-repo/store/models"
	"private-conda-repo/testutils"
)

func TestListAllPackages(t *testing.T) {
	assert := require.New(t)

	ts := newTestServer(ListAllPackages)
	defer ts.Close()

	numChn := 10
	pkgPerChn := 4
	var channels []string
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("new-channel-%d", i)
		chn, _ := repo.CreateChannel(name)

		for j := 0; j < pkgPerChn; j++ {
			_, _ = chn.AddPackage(nil, &condatypes.Package{
				Name:        fmt.Sprintf("package-%d", j),
				Version:     fmt.Sprintf("%d.%d", j, j+1),
				BuildString: "py",
				BuildNumber: 0,
				Platform:    "noarch",
			})
		}
		channels = append(channels, name)
	}

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	defer func() { _ = resp.Body.Close() }()

	var output []*condatypes.ChannelMetaPackageOutput
	err = json.NewDecoder(resp.Body).Decode(&output)
	assert.NoError(err)
	assert.Len(output, numChn*pkgPerChn)

	for _, chn := range channels {
		err := repo.RemoveChannel(chn)
		assert.NoError(err)
	}
}

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
			statusCode:  http.StatusNotFound,
			expectedLen: 0,
		}, {
			input:       "perfana",
			statusCode:  http.StatusOK,
			expectedLen: 1,
		},
	}

	runTest := func(test TestRow) {
		resp, err := http.Get(fmt.Sprintf("%s/%s/%s", ts.URL, channel, test.input))
		assert.NoError(err)
		assert.EqualValues(test.statusCode, resp.StatusCode)

		if test.statusCode == http.StatusOK && resp.StatusCode == http.StatusOK {
			defer func() { _ = resp.Body.Close() }()
			var output ChannelPackageDetails
			err = json.NewDecoder(resp.Body).Decode(&output)
			assert.NoError(err)
			assert.Len(output.Details, test.expectedLen)
			assert.EqualValues(output.Channel, channel)
			assert.EqualValues(output.Package, test.input)

			assert.NotNil(output.Latest)
		}
	}

	for _, test := range tests {
		runTest(test)
	}
}

func TestUploadPackage(t *testing.T) {
	assert := require.New(t)
	channelName := "daniel-upload-package"

	ts := newTestServer(UploadPackage)
	defer ts.Close()
	_, err := db.AddUser(channelName, "pikachu", "daniel@gmail.com")
	assert.NoError(err)
	_, err = repo.CreateChannel(channelName)
	assert.NoError(err)

	formData := func(pkg testutils.TestPackage) (io.ReadWriter, string, error) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		if err := writer.WriteField("channel", channelName); err != nil {
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
	channelName := "daniel-remove-package"

	ts := newTestServer(RemovePackage)
	defer ts.Close()
	_, err := db.AddUser(channelName, "pikachu", "daniel@gmail.com")
	assert.NoError(err)
	chn, err := repo.CreateChannel(channelName)
	assert.NoError(err)

	pkg := &condatypes.Package{
		Name:        "test-package",
		Version:     "0.1",
		BuildString: "py12345",
		BuildNumber: 0,
		Platform:    "noarch",
	}

	_, _ = chn.AddPackage(nil, pkg)
	_, _ = db.CreatePackageCount(&models.PackageCount{
		Channel:     channelName,
		Package:     pkg.Name,
		BuildString: pkg.BuildString,
		BuildNumber: pkg.BuildNumber,
		Version:     pkg.Version,
		Platform:    pkg.Platform,
	})

	payload := fmt.Sprintf(`{
		"channel": "%s",
		"password": "pikachu",
		"package": {
			"name": "test-package",
			"version": "0.1",
			"build_string": "py12345",
			"build_number": 0,
			"platform": "noarch"
		}
	}`, channelName)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", ts.URL, strings.NewReader(payload))
	assert.NoError(err)

	resp, err := client.Do(req)
	assert.NoError(err)
	assert.EqualValues(http.StatusOK, resp.StatusCode)
}
