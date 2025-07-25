package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version",
	Long:  `This command shows the currently installed version of tmpl`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented for now!")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
