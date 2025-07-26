package cmd

import (
	"fmt"
	"log"

	"github.com/IC3P3/tmpl/internal/data"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists every available local repository",
	Long: `This command lists all the available/added repository usable
by the tmpl tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		repositories, err := data.GetRepositoryData()
		if err != nil {
			log.Fatal(err)
		}

		for i, repository := range repositories {
			fmt.Printf("%d: %s => %s\n", i, repository.Name, repository.Remote)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
