package cmd

import (
	"fmt"
	"log"

	"github.com/IC3P3/tmpl/internal/data"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "remove <Name>",
	Args:  cobra.ExactArgs(1),
	Short: "Removes a repository from tmpl",
	Long: `This removes a not needed repository from the list of
available repositories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := data.DeleteRepository(args[0]); err != nil {
			log.Fatal(err)
		}

		if err := data.RemoveRepository(args[0]); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s was removed.\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(removeCmd)
}
