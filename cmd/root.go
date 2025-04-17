package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "dfiles",
	Short:   "Better dotfiles management",
	Long:    `Header Header Header Header`,
	Version: "0.0.1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var commands = []*cobra.Command{
	statusCmd,
	addCmd,
	removeCmd,
	// commitCmd,
	// pushCmd,
	// listCmd,
	// configCmd,
}

func init() {
	cobra.OnInitialize(initConfig)

	for _, cmd := range commands {
		rootCmd.AddCommand(cmd)
	}
}

func initConfig() {}
