package config

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	listConfCmd.Flags().Bool("show", false, "Shows values of the keys when listing all keys")
}

var listConfCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all configuration keys (and values)",
	Args:    cobra.NoArgs,
	Example: "pcr config get ssl_verify",
	Run: func(cmd *cobra.Command, args []string) {
		handler := ListHandler{cmd: cmd}

		show, err := handler.cmd.Flags().GetBool("show")
		if err != nil {
			cmd.PrintErr(err)
			return
		}
		handler.ListAllKeys(show)
	},
}

type ListHandler struct {
	cmd *cobra.Command
}

func (h *ListHandler) ListAllKeys(show bool) {
	options := []string{
		sslVerify,
	}

	for i, key := range options {
		value, err := getValue(key)
		if err != nil {
			h.cmd.PrintErr(err)
			return
		}

		if show {
			options[i] = fmt.Sprintf("%-15s: %v", key, value)
		} else {
			options[i] = fmt.Sprintf("%-15s", key)
		}

	}

	h.cmd.Printf("Available Keys\n%s", strings.Join(options, "\n"))
}
