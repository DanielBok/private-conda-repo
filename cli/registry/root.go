package registry

import (
	"fmt"
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

		if len(args) > 0 {
			c := strings.Join(args, " ")
			cmd.Printf("%s is not a valid command", c)
			return
		}

		if registry == "" {
			registry = "<undefined: Please set registry with 'pcr registry set'>"
		}

		channel := conf.Channel.Channel
		if channel == "" {
			channel = "<Not logged in: Please login with 'pcr registry login'>"
		}

		cmd.Println(strings.TrimSpace(fmt.Sprintf(`
CLI Registry details:
	Registry:     %s
	Channel :     %s

Use "%s registry --help" for more information.

Registry should be the api server that is used to create and add channel. 
Usually this is https://<host>:5060. Remember to set it with the "set"
command. 
`, registry, channel, strings.Split(cmd.CommandPath(), " ")[0])))
	},
}

func init() {
	RootCmd.AddCommand(setCmd, loginCmd, logoutCmd, registerCmd)
}
