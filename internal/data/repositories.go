package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type TemplateConfig struct {
	Displayname string
	Name        string
	Description string
	Addons      []string
}

type AddonConfig struct {
	DisplayName string
	Name        string
	Description string
}

type HookConfig struct{}

type RepositoryConfig struct {
	Displayname string
	Description string
	Templates   []TemplateConfig
	Addons      []AddonConfig
	Hooks       []HookConfig
}

func DeleteRepository(name string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not find a home directory.")
	}

	if err := os.RemoveAll(homeDir + "/.local/share/tmpl/" + name); err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not delete repository")
	}

	fmt.Printf("Successfully deleted %s from the list.\n", name)

	return nil
}

func GetRepositoryDescription(name string) (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return "", errors.New("Could not find a home directory.")
	}

	file, err := os.ReadFile(homeDir + "/.local/share/tmpl/" + name + "/tmpl.json")
	if err != nil {
		fmt.Println(err.Error())

		return "", errors.New("No file or permission to read the file.")
	}

	repositoryContent := RepositoryConfig{}
	if err := json.Unmarshal(file, &repositoryContent); err != nil {
		fmt.Println(err.Error())

		return "", errors.New("Could not parse repository.json")
	}

	return repositoryContent.Description, nil
}

func GetTemplateData(repoName string) ([]TemplateConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return []TemplateConfig{}, errors.New("Could not find a home directory.")
	}

	file, err := os.ReadFile(homeDir + "/.local/share/tmpl/" + repoName + "/tmpl.json")
	if err != nil {
		fmt.Println(err.Error())

		return []TemplateConfig{}, errors.New("No file or permission to read the file.")
	}

	repositoryContent := RepositoryConfig{}
	if err := json.Unmarshal(file, &repositoryContent); err != nil {
		fmt.Println(err.Error())

		return []TemplateConfig{}, errors.New("Could not parse repository.json")
	}

	return repositoryContent.Templates, nil
}

func GetAddonData(repoName string) ([]AddonConfig, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return []AddonConfig{}, errors.New("Could not find a home directory.")
	}

	file, err := os.ReadFile(homeDir + "/.local/share/tmpl/" + repoName + "/tmpl.json")
	if err != nil {
		fmt.Println(err.Error())

		return []AddonConfig{}, errors.New("No file or permission to read the file.")
	}

	repositoryContent := RepositoryConfig{}
	if err := json.Unmarshal(file, &repositoryContent); err != nil {
		fmt.Println(err.Error())

		return []AddonConfig{}, errors.New("Could not parse repository.json")
	}

	return repositoryContent.Addons, nil
}
