package cmd

import (
	"fmt"
	"log"

	"github.com/IC3P3/tmpl/internal/data"
	"github.com/IC3P3/tmpl/internal/git"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <Name> <Repository Remote>",
	Args:  cobra.ExactArgs(2),
	Short: "Add a new repository containing templates",
	Long: `The add command adds a new repository to the list of
available templates.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := data.AddRepository(args[0], args[1]); err != nil {
			log.Fatal(err)
		}

		if err := git.GetTemplateRepository(args[0], args[1]); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("The repository '%s' was successfully cloned and added.\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
