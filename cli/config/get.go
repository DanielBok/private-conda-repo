package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	getConfCmd.Flags().Bool("values", false, "Shows values of the keys when listing all keys")
}

var getConfCmd = &cobra.Command{
	Use:     "get",
	Short:   "Gets the configuration value",
	Long:    "Gets the configuration value. If name is not given, returns the name of all config keys",
	Args:    cobra.ArbitraryArgs,
	Example: "pcr config get ssl_verify",
	Run: func(cmd *cobra.Command, args []string) {
		handler := getHandler{cmd: cmd}
		if len(args) == 0 {
			handler.ListAllKeys()
		}

		for _, arg := range args {
			handler.Get(arg)
		}
	},
}

type getHandler struct {
	cmd *cobra.Command
}

func (g *getHandler) ListAllKeys() {
	options := []string{
		sslVerify,
	}

	if showValues, err := g.cmd.Flags().GetBool("values"); showValues && err == nil {
		for i, key := range options {
			value, err := g.getValue(key)
			if err != nil {
				log.Fatalln(err)
				return
			}
			options[i] = fmt.Sprintf("%-15s: %v", key, value)
		}
	}

	log.Printf("Available Keys\n%s", strings.Join(options, "\n"))
}

func (g *getHandler) Get(key string) {
	value, err := g.getValue(key)
	if err != nil {
		log.Fatalf("Invalid option: %s\n", key)
		return
	}
	log.Printf("%s: %v\n", sslVerify, value)
}

func (g *getHandler) getValue(key string) (interface{}, error) {
	key = strings.ToLower(strings.TrimSpace(key))
	conf := New()

	switch key {
	case sslVerify:
		return conf.SslVerify, nil
	default:
		return nil, errors.Errorf("Invalid option: %s\n", key)
	}
}
