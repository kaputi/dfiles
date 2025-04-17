package cmd

import (
	"fmt"

	"github.com/kaputi/dfiles/core"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show the status of the dotfiles",
	Long:  `Header Header Header Header`,

	Run: func(cmd *cobra.Command, args []string) {
		tracked := core.GetTraked()

		fmt.Println("Tracked files:")
		core.Separator()
		core.PrintList(tracked)
		core.Separator()

		changed := core.GetChanged()
		fmt.Println("Changes:")
		core.Separator()
		core.PrintList(changed)
	},
}
