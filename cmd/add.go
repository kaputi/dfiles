package cmd

import (
	"github.com/kaputi/dfiles/core"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "ADD ADD ADD",
	Long:  `Header Header Header Header`,

	Run: func(cmd *cobra.Command, args []string) {
		core.Add(args[0])
	},
}
