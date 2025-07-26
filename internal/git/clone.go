package git

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func GetTemplateRepository(name string, remote string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not find a home directory.")
	}

	repository, err := git.PlainClone(homeDir+"/.local/share/tmpl/"+name, false, &git.CloneOptions{
		URL:   remote,
		Depth: 1,
	})
	if err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not clone the repository.")
	}

	fmt.Println(repository)

	return nil
}
