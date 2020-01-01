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
		channel := conf.Channel.Channel
		conf.Channel.Channel = ""
		conf.Channel.Password = ""

		conf.Save()
		if channel == "" {
			log.Fatalln("You're not logged in")
		} else {
			log.Printf("logged out of '%s'", channel)
		}

	},
}
