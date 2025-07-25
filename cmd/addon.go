package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addonCmd = &cobra.Command{
	Use:   "addon",
	Short: "Add smaller addons after the initial creation",
	Long: `Run this command to add some of the smaller extensions of the
template after it was already created.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(addonCmd)
}
