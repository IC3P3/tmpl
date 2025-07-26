package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmpl",
	Short: "Generate the boilerplate code to your new project.",
	Long: `tmpl is a CLI tool used to quickly create all kinds of
different boilderplate code. These templates are defined by a
Git repository to create different base templates and addons
for these to your needs.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
