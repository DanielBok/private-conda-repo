package config

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var setConfCmd = &cobra.Command{
	Use:     "set",
	Short:   "Sets configuration values",
	Args:    cobra.RangeArgs(1, 2),
	Example: "pcr config set ssl_verify true",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			args = strings.SplitN(args[0], "=", 2)
		}

		if len(args) != 2 {
			cmd.PrintErr("set values must be pairs or separated by '='. Example, 'ssl_verify=true' or 'ssl_verify true'")
			return
		}

		handler := SetHandler{cmd: cmd}
		conf, err := handler.Set(args[0], args[1])
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		conf.Save()
	},
}

type SetHandler struct {
	cmd *cobra.Command
}

func (s *SetHandler) Set(key, value string) (*Config, error) {
	key = strings.ToLower(strings.TrimSpace(key))
	value = strings.ToLower(strings.TrimSpace(value))
	conf := New()

	err := func() (*Config, error) {
		return nil, errors.Errorf("Invalid option for %s: %s", key, value)
	}

	switch key {
	case sslVerify:
		switch value {
		case "true", "t", "1":
			conf.SslVerify = true
		case "false", "f", "0":
			conf.SslVerify = false
		default:
			return err()
		}
	default:
		return nil, errors.Errorf("Invalid option: %s", key)
	}
	return conf, nil
}
