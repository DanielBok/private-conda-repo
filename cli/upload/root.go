package upload

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"cli/config"
	"cli/request"
)

func init() {
	RootCmd.Flags().Bool("no-abi", false, "If true, removes the 'python_abi' dependency for the channel")
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "upload",
	Short:   "Uploads the built conda package",
	Args:    cobra.ExactArgs(1),
	Example: "pcr upload dist/noarch/numpy-0.1.1-py_0.tar.bz",
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		handler := newUploadHandler(cmd)

		err := handler.verifyPackage(file)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		payload, err := handler.createPayload()
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		pkg, err := handler.upload(payload)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Printf(strings.TrimSpace(fmt.Sprintf(`
Uploaded file '%s' successfully
Details
	Name:         %s
	Version:      %s
	Build String: %s
	Build Number  %d
`, file, pkg.Name, pkg.Version, pkg.BuildString, pkg.BuildNumber)))
	},
}

type Handler struct {
	cmd         *cobra.Command
	url         string
	channel     string
	password    string
	packagePath string
}

type Payload struct {
	Body        *bytes.Buffer
	ContentType string
}

func newUploadHandler(cmd *cobra.Command) Handler {
	conf := config.New()
	h := Handler{
		cmd:      cmd,
		url:      conf.Registry + "/p",
		channel:  conf.Channel.Channel,
		password: conf.Channel.Password,
	}
	return h
}

func (h *Handler) verifyPackage(file string) error {
	if !strings.HasSuffix(file, ".tar.bz2") {
		return errors.New("expect conda package should have extension '.tar.bz2'")
	}

	if cwd, err := os.Getwd(); err == nil {
		// relative path
		path := filepath.Join(cwd, file)
		if _, err := os.Stat(path); err == nil {
			h.packagePath = path
			return nil
		}
	}

	// absolute path
	if _, err := os.Stat(file); err == nil {
		h.packagePath = file
		return nil
	}

	return errors.New("package does not exist")
}

func (h *Handler) createPayload() (*Payload, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)

	err := w.WriteField("channel", h.channel)
	if err != nil {
		return nil, err
	}

	err = w.WriteField("password", h.password)
	if err != nil {
		return nil, err
	}

	fixes, err := h.getFixFlags()
	if err != nil {
		return nil, err
	}
	err = w.WriteField("fixes", fixes)
	if err != nil {
		return nil, err
	}

	parts, err := w.CreateFormFile("file", filepath.Base(h.packagePath))
	if err != nil {
		return nil, err
	}

	// Write package file into form field
	file, err := os.Open(h.packagePath)
	if err != nil {
		return nil, errors.Errorf("could not open file at %s", h.packagePath)
	}
	if _, err = io.Copy(parts, file); err != nil {
		return nil, errors.New("could not copy file to form payload")
	}

	if err := w.Close(); err != nil {
		return nil, errors.New("could not close form for upload")
	}

	return &Payload{
		Body:        &body,
		ContentType: w.FormDataContentType(),
	}, nil
}

func (h *Handler) getFixFlags() (string, error) {
	var flags []string

	if noAbi, err := h.cmd.Flags().GetBool("no-abi"); err != nil {
		return "", err
	} else if noAbi {
		flags = append(flags, "no-abi")
	}

	return strings.Join(flags, ","), nil
}

func (h *Handler) upload(payload *Payload) (*Package, error) {
	resp, err := request.Post(h.url, payload.ContentType, payload.Body)
	if err != nil {
		return nil, errors.Wrap(err, "package upload failed")
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode == 200 {
		var output Package
		if err := json.NewDecoder(resp.Body).Decode(&output); err != nil {
			return nil, err
		}
		return &output, nil
	}

	errorMessage, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Double whammy. Upload failed and cannot parse reason")
	}

	return nil, errors.Errorf("Upload failed: %s", string(errorMessage))
}
