package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Uploads the built conda package",
	Args:    cobra.ExactArgs(1),
	Example: "pcr upload dist/noarch/numpy-0.1.1-py_0.tar.bz",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		handler := newUploadHandler()

		handler.verifyPackage(file)
		pkg := handler.uploadPackage()

		log.Printf(strings.TrimSpace(fmt.Sprintf(`
Uploaded file '%s' successfully
Details
	Name:         %s
	Version:      %s
	Build String: %s
	Build Number  %d
`, file, pkg.Name, pkg.Version, pkg.BuildString, pkg.BuildNumber)))
	},
}

type uploadHandler struct {
	cmd         *cobra.Command
	url         string
	channel     string
	password    string
	packagePath string
}

func newUploadHandler() uploadHandler {
	conf := config.New()
	h := uploadHandler{
		cmd:      nil,
		url:      conf.Registry + "/p",
		channel:  conf.Channel.Channel,
		password: conf.Channel.Password,
	}
	return h
}

func (h *uploadHandler) verifyPackage(file string) {
	if !strings.HasSuffix(file, ".tar.bz2") {
		log.Fatalln("expect conda package should have extension '.tar.bz2'")
	}

	if cwd, err := os.Getwd(); err == nil {
		// relative path
		path := filepath.Join(cwd, file)
		if _, err := os.Stat(path); err == nil {
			h.packagePath = path
			return
		}
	}

	// absolute path
	if _, err := os.Stat(file); err == nil {
		h.packagePath = file
		return
	}

	log.Fatal("package does not exist")
}

func (h *uploadHandler) uploadPackage() *Package {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	if err := writer.WriteField("channel", h.channel); err != nil {
		log.Fatal(err)
	}
	if err := writer.WriteField("password", h.password); err != nil {
		log.Fatal(err)
	}
	parts, err := writer.CreateFormFile("file", filepath.Base(h.packagePath))
	if err != nil {
		log.Fatal(err)
	}

	// Write package file into form field
	file, err := os.Open(h.packagePath)
	if err != nil {
		log.Fatalf("could not open file at %s", h.packagePath)
	}
	if _, err = io.Copy(parts, file); err != nil {
		log.Fatalf("could not copy file to form payload")
	}

	if err := writer.Close(); err != nil {
		log.Fatalf("could not close form for upload")
	}

	resp, err := http.Post(h.url, writer.FormDataContentType(), body)
	if err != nil {
		log.Fatal(errors.Wrap(err, "package upload failed"))
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == 200 {
		var output Package
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			log.Fatal(err)
		}
		return &output
	}

	errorMessage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Double whammy. Upload failed and cannot parse reason")
	}
	log.Fatalf("Upload failed: %s", string(errorMessage))

	return nil
}
