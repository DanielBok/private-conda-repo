package main

import (
	"log"

	"github.com/spf13/cobra"

	"cli/config"
	"cli/registry"
	"cli/upload"
)

var rootCmd = &cobra.Command{
	Use:   "pcr",
	Short: "Private Conda Repository Command Line Tool",
	Long: `Private Conda Repository command line tool.
Aids in various aspect of using the Private Conda Repository
application. This tool is catered for package contributors.
`,
	Version: "3.0",
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	rootCmd.AddCommand(config.RootCmd, registry.RootCmd, upload.RootCmd, versionCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is Private Conda Repo CLI's`,
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Printf("Private Conda Repo CLI: %s", rootCmd.Version)
	},
}
