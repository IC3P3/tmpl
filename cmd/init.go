package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new repository for templates",
	Long: `This creates a new repository with the needed configuration files
to add your own templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
