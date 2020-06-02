package config

import (
	"strings"

	"private-conda-repo/libs"
)

type TLSConfig struct {
	Cert string `mapstructure:"cert"`
	Key  string `mapstructure:"key"`
}

func (t *TLSConfig) HasCert() bool {
	t.Cert = strings.TrimSpace(t.Cert)
	t.Key = strings.TrimSpace(t.Key)
	if t.Cert == "" || t.Key == "" {
		return false
	}

	if !libs.PathExists(t.Cert) {
		return false
	}

	if !libs.PathExists(t.Key) {
		return false
	}

	return true
}
