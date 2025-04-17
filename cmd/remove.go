package cmd

import (
	"github.com/kaputi/dfiles/core"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "remore remove remove",
	Long:  `Header Header Header Header`,

	Run: func(cmd *cobra.Command, args []string) {
		core.Remove(args[0])
	},
}
