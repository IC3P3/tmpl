package cmd

import (
	"fmt"
	"log"

	"github.com/IC3P3/tmpl/internal/data"
	"github.com/IC3P3/tmpl/internal/git"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the added repositories",
	Long: `Fetches the added repositories for any new changes and
pulls them again to apply changes.`,
	Run: func(cmd *cobra.Command, args []string) {
		repositories, err := data.GetRepositories()
		if err != nil {
			log.Fatal(err)
		}

		for _, repository := range repositories {
			if err := git.UpdateRepository(repository.Name); err != nil {
				log.Fatal(err)
			}
		}

		fmt.Println("Updated all template repositories.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
