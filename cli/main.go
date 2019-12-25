package main

import (
	"log"

	"github.com/spf13/cobra"

	"cli/registry"
)

var rootCmd = &cobra.Command{
	Use:   "pcr",
	Short: "Private Conda Repository Command Line Tool",
	Long: `Private Conda Repository command line tool.
Aids in various aspect of using the Private Conda Repository
application. This tool is catered for package contributors.
`,
}

func main() {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	rootCmd.AddCommand(registry.RootCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
