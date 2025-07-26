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

	if _, err := git.PlainClone(homeDir+"/.local/share/tmpl/"+name, false, &git.CloneOptions{
		URL:   remote,
		Depth: 1,
	}); err != nil {
		fmt.Println(err.Error())

		return errors.New("Could not clone the repository.")
	}

	return nil
}
