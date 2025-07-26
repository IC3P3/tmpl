package commands

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Answer struct {
	Id          int
	Name        string
	Description string
}

func AskFromListMultiple(description string, options []Answer) ([]Answer, error) {
	fmt.Println(description)

	for _, option := range options {
		fmt.Printf("\n%d: \t%s\t\t\t%s\n\n", option.Id, option.Name, option.Description)
	}

	fmt.Printf("Enter a selection space-seperated (%d-%d): ", 1, len(options)+1)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	inputStrings := strings.Fields(input)

	var answer []Answer

	for _, s := range inputStrings {
		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Printf("Warning: '%s' is not a valid number, skipping.\n", s)
			continue
		}

		answer = append(answer, options[num-1])
	}

	return answer, nil
}

func AskFromListSingle(description string, options []Answer) (Answer, error) {
	fmt.Println(description)

	for _, option := range options {
		fmt.Printf("\n%d: \t%s\t\t\t%s\n\n", option.Id+1, option.Name, option.Description)
	}

	fmt.Printf("Enter a one number (%d-%d): ", 1, len(options))

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	inputStrings := strings.Fields(input)

	num, err := strconv.Atoi(inputStrings[0])
	if err != nil {
		return Answer{}, fmt.Errorf("Warning: '%s' is not a valid number, skipping.\n", inputStrings[0])
	}

	return options[num-1], nil
}

func AskForText(description string) string {
	fmt.Printf("%s: ", description)

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	return input
}
