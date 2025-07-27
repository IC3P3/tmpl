package git

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
)

func UpdateRepository(name string) error {
	repository, err := openRepository(name)
	if err != nil {
		return err
	}

	worktree, err := repository.Worktree()
	if err != nil {
		fmt.Println(err)

		return fmt.Errorf("Could not access the worktree of %s.", name)
	}

	if err := worktree.Pull(&git.PullOptions{}); err != nil {
		fmt.Println(err)

		return fmt.Errorf("Could not update the repository.")
	}

	return nil
}

func openRepository(name string) (*git.Repository, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err.Error())

		return nil, errors.New("Could not find a home directory.")
	}

	repository, err := git.PlainOpen(homeDir + "/.local/share/tmpl/" + name)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Could not open the repository %s", name)
	}

	return repository, nil
}
