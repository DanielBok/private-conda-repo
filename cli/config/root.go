package config

import (
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "config",
	Short: "PCR cli configuration",
}

func init() {
	RootCmd.AddCommand(setConfCmd, getConfCmd, listConfCmd)
}
