package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the added repositories",
	Long: `Fetches the added repositories for any new changes and
pulls them again to apply changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
