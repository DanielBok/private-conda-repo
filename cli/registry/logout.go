package registry

import (
	"log"

	"github.com/spf13/cobra"

	"cli/config"
)

var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out of the registry",
	Long:  `Removes the user's credentials from the cli tool.`,
	Args:  cobra.NoArgs,
	Run: func(_ *cobra.Command, _ []string) {
		conf := config.New()
		user := conf.User.Username
		conf.User.Username = ""
		conf.User.Password = ""

		conf.Save()
		log.Printf("logged out of '%s'", user)
	},
}
