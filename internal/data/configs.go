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

// TODO: Check if name is unique
func AddRepository(name string, remote string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not find a home directory.")
	}

	if err := createMissing(homeDir + "/.local/share/tmpl/"); err != nil {
		return err
	}

	repositories, _ := GetRepositoryData()
	if repositories == nil {
		repositories = []RepositoryData{}
	}
	repositories = append(repositories, RepositoryData{Name: name, Remote: remote})

	jsonData, err := json.Marshal(repositories)
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not parse data to json")
	}

	if err := os.WriteFile(homeDir+"/.local/share/tmpl/repositories.json", jsonData, 0750); err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not write to file")
	}

	return nil
}

func RemoveRepository(name string) error {
	repositories, _ := GetRepositoryData()
	if repositories == nil {
		return errors.New("No repository available")
	}

	deleted := false
	for i, repository := range repositories {
		if repository.Name == name {
			repositories = append(repositories[:i], repositories[i+1:]...)
			deleted = true
		}
	}

	if !deleted {
		return errors.New("Could not find repository with that name.")
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not find a home directory.")
	}

	jsonData, err := json.Marshal(repositories)
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not parse data to json")
	}

	if err := os.WriteFile(homeDir+"/.local/share/tmpl/repositories.json", jsonData, 0750); err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not write to file")
	}

	return nil

}

func createMissing(path string) error {
	if err := os.MkdirAll(path, 0750); err != nil {
		fmt.Println(err.Error())
		return errors.New("Could not create missing directories.")
	}

	return nil
}
