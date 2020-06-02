package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var getConfCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the configuration value",
	Args:    cobra.MinimumNArgs(1),
	Example: "pcr config get ssl_verify",
	Run: func(cmd *cobra.Command, args []string) {
		handler := GetHandler{cmd: cmd}

		for _, arg := range args {
			handler.Get(arg)
		}
	},
}

type GetHandler struct {
	cmd *cobra.Command
}

func (g *GetHandler) Get(key string) {
	value, err := getValue(key)
	if err != nil {
		g.cmd.PrintErrf("Invalid option: %s\n", key)
		return
	}
	g.cmd.Printf("%s: %v\n", sslVerify, value)
}

func getValue(key string) (interface{}, error) {
	key = strings.ToLower(strings.TrimSpace(key))
	conf := New()

	switch key {
	case sslVerify:
		return conf.SslVerify, nil
	default:
		return nil, errors.Errorf("Invalid option: %s\n", key)
	}
}
