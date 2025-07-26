package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/IC3P3/tmpl/internal/data"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <Template Repository> <Template Name> <Project Name>",
	Args:  cobra.ExactArgs(3),
	Short: "Create a new project from a template",
	Long: `Start the process of creating a project from the
available template configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		data.CopyDirectory(homeDir+"/.local/share/tmpl/"+args[0]+"/templates/"+args[1], currentDir+"/"+args[2])

		fmt.Println("The project was successfully created.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
