package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the syntax of the template repository",
	Long: `Looks up all the configuration files and checks if every one of these
has the correct syntax.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
