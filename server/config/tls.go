package config

import (
	"os"
	"strings"
)

type tls struct {
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

func (t *tls) HasCert() bool {
	t.Cert = strings.TrimSpace(t.Cert)
	t.Key = strings.TrimSpace(t.Key)
	if t.Cert == "" || t.Key == "" {
		return false
	}

	if _, err := os.Stat(t.Cert); os.IsNotExist(err) {
		return false
	}

	if _, err := os.Stat(t.Key); os.IsNotExist(err) {
		return false
	}

	return true
}
