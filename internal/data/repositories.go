package data

import (
	"errors"
	"fmt"
	"os"
)

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
