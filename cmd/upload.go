package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload your own template repository",
	Long: `Uploads the repository to the remote and checks for any
syntax errors.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet!")
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
