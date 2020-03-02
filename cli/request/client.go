package request

import (
	"crypto/tls"
	"io"
	"net/http"

	"cli/config"
)

func NewClient() *http.Client {
	conf := config.New()
	if conf.SslVerify {
		return http.DefaultClient
	} else {
		tr := http.DefaultTransport.(*http.Transport).Clone()
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

		return &http.Client{Transport: tr}
	}
}

func Get(url string) (*http.Response, error) {
	return NewClient().Get(url)
}

func Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return NewClient().Post(url, contentType, body)
}
