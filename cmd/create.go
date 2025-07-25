package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project from a template",
	Long: `Start the process of creating a project from the
available template configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented for now!")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
