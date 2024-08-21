package userInput

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(label string) (string, error) {
	if label == "" {
		return "", errors.New("input label cannot be empty")
	}

	var userInput string

	fmt.Println(label)
	// fmt.Scanln(&userInput)
	reader := bufio.NewReader(os.Stdin)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.TrimSuffix(userInput, "\n")
	userInput = strings.TrimSuffix(userInput, "\r")

	if userInput == "" {
		return "", fmt.Errorf("invalid input received for %v", label)
	}

	return userInput, nil
}
