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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return nil, errors.New("Could not find a home directory.")
	}

	file, err := os.ReadFile(homeDir + "/.local/share/tmpl/repositories.json")
	if err != nil {
		fmt.Println(err.Error())

		return nil, errors.New("No file or permission to read the file.")
	}

	repositories := []RepositoryData{}
	if err := json.Unmarshal(file, &repositories); err != nil {
		fmt.Println(err.Error())

		return nil, errors.New("Could not parse repository.json")
	}

	return repositories, nil
}
