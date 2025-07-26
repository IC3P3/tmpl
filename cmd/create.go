package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/IC3P3/tmpl/internal/commands"
	"github.com/IC3P3/tmpl/internal/data"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Args:  cobra.ExactArgs(0),
	Short: "Create a new project from a template",
	Long: `Start the process of creating a project from the
available template configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		projectName := commands.AskForText("Enter the name of your new project")

		repositories, err := data.GetRepositories()
		if err != nil {
			log.Fatal(err)
		}

		var useRepoOptions []commands.Answer
		for i, repository := range repositories {
			description, err := data.GetRepositoryDescription(repository.Name)
			if err != nil {
				log.Fatal(err)
			}

			useRepoOptions = append(useRepoOptions, commands.Answer{Id: i, Name: repository.Name, Description: description})
		}

		templateRepository, err := commands.AskFromListSingle("Choose the repository you want to use.", useRepoOptions)
		if err != nil {
			log.Fatal(err)
		}

		templates, err := data.GetTemplateData(templateRepository.Name)
		if err != nil {
			log.Fatal(err)
		}

		var useTempOptions []commands.Answer
		for i, template := range templates {
			useTempOptions = append(useTempOptions, commands.Answer{Id: i, Name: template.Name, Description: template.Description})
		}

		template, err := commands.AskFromListSingle("Choose the repository you want to use.", useTempOptions)
		if err != nil {
			log.Fatal(err)
		}

		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}

		data.CopyDirectory(homeDir+"/.local/share/tmpl/"+templateRepository.Name+"/templates/"+template.Name, currentDir+"/"+projectName)

		fmt.Println("The project was successfully created.")
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
