package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Removes a repository from tmpl",
	Long: `This removes a not needed repository from the list of
available repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
