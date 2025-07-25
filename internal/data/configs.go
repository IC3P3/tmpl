package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type RepositoryData struct {
	Name   string
	Remote string
}

func GetRepositoryData() ([]RepositoryData, error) {
	configFile, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		noHome := errors.New("Could not find a home directory.")
		return nil, noHome
	}

	file, err := os.ReadFile(configFile + "/.local/share/tmpl/repositories.json")
	if err != nil {
		fmt.Println(err.Error())

		noFile := errors.New("No file or permission to read the file.")
		return nil, noFile
	}

	repositories := []RepositoryData{}
	if err := json.Unmarshal(file, &repositories); err != nil {
		fmt.Println(err.Error())

		parsingError := errors.New("Could not parse repository.json")
		return nil, parsingError
	}

	return repositories, nil
}
