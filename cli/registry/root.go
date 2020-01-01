package registry

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"cli/config"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "registry",
	Short: "Logs the user into the system",
	Long: `Logs the user into the system. This will raise an
error if the private conda repository's url is not set.`,
	Run: func(cmd *cobra.Command, args []string) {
		conf := config.New()
		registry := conf.Registry
		if registry == "" {
			registry = "<undefined: Please set registry with 'pcr registry set'>"
		}

		channel := conf.Channel.Channel
		if channel == "" {
			channel = "<Not logged in: Please login with 'pcr registry login'>"
		}

		log.Println(strings.TrimSpace(fmt.Sprintf(`
CLI Registry details:
	Registry:     %s
	Channel :     %s
`, registry, channel)))
	},
}

func init() {
	RootCmd.AddCommand(setCmd, loginCmd, logoutCmd, registerCmd)
}
