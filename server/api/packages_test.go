package api_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	. "private-conda-repo/api"
	"private-conda-repo/api/dto"
	"private-conda-repo/api/interfaces"
	"private-conda-repo/domain/entity"
	"private-conda-repo/libs"
	"private-conda-repo/testutils"
)

const (
	pkgName    = "numpy" // there is only 1 unique package per channel, it just has different versions
	nSeedChn   = 2
	nVerPerPkg = 3 // num of different versions per pkg
)

func NewPackageHandler() *PackageHandler {
	return &PackageHandler{
		DB:           NewMockDb(),
		Decompressor: NewMockDecompressor(),
		FileSys:      NewMockFileSys(),
	}
}

func seedPackages(db interfaces.DataAccessLayer, fs interfaces.FileSys, chnPrefix string) error {
	for i := 0; i < nSeedChn; i++ {
		chn, err := db.CreateChannel(fmt.Sprintf("%s-channel-%d", chnPrefix, i), password, email)
		if err != nil {
			return err
		}

		_, err = fs.CreateChannel(chn.Channel)
		if err != nil {
			return err
		}

		for j := 0; j < nVerPerPkg; j++ {
			pkg, err := db.CreatePackageCount(&entity.PackageCount{
				ChannelId:   chn.Id,
				Package:     pkgName,
				BuildString: pkgName,
				BuildNumber: 0,
				Version:     strconv.Itoa(j + 1),
				Platform:    "noarch",
			})
			if err != nil {
				return err
			}

			for k := 0; k < 5; k++ {
				_, err := db.IncreasePackageCount(pkg)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func TestPackageHandler_ListAllPackages(t *testing.T) {
	handler := NewPackageHandler()
	assert := require.New(t)

	err := seedPackages(handler.DB, handler.FileSys, "list-packages")
	assert.NoError(err)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)

	handler.ListAllPackages()(w, r)
	assert.Equal(http.StatusOK, w.Code)

	var result []*dto.ChannelData
	err = json.NewDecoder(w.Body).Decode(&result)
	assert.NoError(err)
	assert.Len(result, nSeedChn) // only 2 since each channel only has 1 unique package (so 2 x 1 = 2)
}

func TestPackageHandler_ListPackagesInChannel(t *testing.T) {
	handler := NewPackageHandler()
	assert := require.New(t)

	prefix := "list-packages-in-channel"
	err := seedPackages(handler.DB, handler.FileSys, prefix)
	assert.NoError(err)

	for _, test := range []struct {
		Channel      string
		ExpectedCode int
	}{
		{fmt.Sprintf("%s-channel-0", prefix), http.StatusOK},
		{"channel-that-does-not-exist", http.StatusBadRequest},
	} {
		w := httptest.NewRecorder()
		r := NewTestRequest("GET", "/", nil, map[string]string{
			"channel": test.Channel,
		})

		handler.ListPackagesInChannel()(w, r)
		assert.Equal(test.ExpectedCode, w.Code)

		if test.ExpectedCode == http.StatusOK {
			var result []*dto.ChannelData
			err = json.NewDecoder(w.Body).Decode(&result)
			assert.NoError(err)
			assert.Len(result, 1) // only 1 unique package per channel
		}
	}
}

func TestPackageHandler_FetchPackageDetails(t *testing.T) {
	handler := NewPackageHandler()
	assert := require.New(t)

	prefix := "fetch-package-details"
	err := seedPackages(handler.DB, handler.FileSys, prefix)
	assert.NoError(err)

	w := httptest.NewRecorder()
	r := NewTestRequest("GET", "/", nil, map[string]string{
		"channel": fmt.Sprintf("%s-channel-0", prefix),
		"pkg":     pkgName,
	})

	handler.FetchPackageDetails()(w, r)
	assert.Equal(http.StatusOK, w.Code)
}

func TestPackageHandler_RemoveAllPackages(t *testing.T) {
	handler := NewPackageHandler()
	assert := require.New(t)

	prefix := "remove-all-packages"
	err := seedPackages(handler.DB, handler.FileSys, prefix)
	assert.NoError(err)

	channel := fmt.Sprintf("%s-channel-0", prefix)
	for _, test := range []struct {
		Channel      string
		Password     string
		ExpectedCode int
	}{
		{channel, password, http.StatusOK},
		{"wrong" + channel, password, http.StatusBadRequest},
		{channel, "wrong" + password, http.StatusForbidden},
	} {
		w := httptest.NewRecorder()

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(&dto.ChannelPackage{
			Channel:  test.Channel,
			Password: test.Password,
		})
		assert.NoError(err)

		r := NewTestRequest("DELETE", "/", &buf, map[string]string{
			"pkg": pkgName,
		})
		handler.RemoveAllPackages()(w, r)
		assert.Equal(test.ExpectedCode, w.Code)
	}

}

func TestPackageHandler_RemovePackage(t *testing.T) {
	handler := NewPackageHandler()
	assert := require.New(t)

	prefix := "remove-package"
	err := seedPackages(handler.DB, handler.FileSys, prefix)
	assert.NoError(err)

	channel := fmt.Sprintf("%s-channel-0", prefix)
	pkg := &dto.PackageDto{
		Name:        pkgName,
		BuildString: pkgName,
		BuildNumber: 0,
		Version:     "1",
		Platform:    "noarch",
	}

	for _, test := range []struct {
		Channel      string
		Password     string
		Package      *dto.PackageDto
		ExpectedCode int
	}{
		{channel, password, pkg, http.StatusOK},
		{channel, password, nil, http.StatusBadRequest},
		{"wrong" + channel, password, pkg, http.StatusBadRequest},
		{channel, "wrong" + password, pkg, http.StatusForbidden},
	} {
		w := httptest.NewRecorder()

		var buf bytes.Buffer
		err = json.NewEncoder(&buf).Encode(&dto.ChannelPackage{
			Channel:  test.Channel,
			Password: test.Password,
			Package:  test.Package,
		})
		assert.NoError(err)

		r := NewTestRequest("DELETE", "/", &buf, nil)
		handler.RemovePackage()(w, r)
		assert.Equal(test.ExpectedCode, w.Code)
	}
}

func TestPackageHandler_UploadPackage(t *testing.T) {
	assert := require.New(t)

	testPkg, err := testutils.GetPackageN(0)
	assert.NoError(err)

	handler := NewPackageHandler()
	handler.Decompressor = &MockDecompressor{
		Filepath: testPkg.Path,
	}

	prefix := "upload-package"
	err = seedPackages(handler.DB, handler.FileSys, prefix)
	assert.NoError(err)

	channel := fmt.Sprintf("%s-channel-0", prefix)

	createFormData := func(pkg testutils.TestPackage) (io.ReadWriter, string, func()) {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		err := writer.WriteField("channel", channel)
		assert.NoError(err)

		err = writer.WriteField("password", password)
		assert.NoError(err)

		err = writer.WriteField("fixes", "no-abi")
		assert.NoError(err)

		parts, err := writer.CreateFormFile("file", pkg.Filename)
		assert.NoError(err)

		// Write package file into form field
		file, err := os.Open(pkg.Path)
		assert.NoError(err)

		_, err = io.Copy(parts, file)
		assert.NoError(err)

		err = writer.Close()
		assert.NoError(err)

		return body, writer.FormDataContentType(), func() {
			libs.IOCloser(file)
		}
	}

	body, contentType, closer := createFormData(testPkg)

	w := httptest.NewRecorder()
	r := NewTestRequest("POST", "/", body, nil)
	r.Header.Set("Content-Type", contentType)

	handler.UploadPackage()(w, r)
	assert.Equal(http.StatusOK, w.Code)

	closer()
}
