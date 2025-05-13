package util

import (
	"fmt"
	"strings"
)

// InputData holds the parsed command-line arguments
type InputData struct {
	Text   string
	Banner string
}

// ParseInput processes and validates command-line arguements
func ParseInput(args []string) (InputData, error) {
	// Check length of text input
	if len(args) < 1 {
		// Return empty struct and error if minimum length not met
		return InputData{}, fmt.Errorf(
			"Usage: go run . \"<text-input>\" [banner-name]\nNote: Enclose text in quotes if it contains spaces",
		)
	}

	// Replace all newline inputs as newline characters
	textInput := strings.ReplaceAll(args[0], `\n`, "\n")
	
	// Check if string input is empty
	if textInput == "" {
		return InputData{}, fmt.Errorf("error: Text input cannot be empty")
	}

	// Set default banner style for formatting
	banner := "standard"

	// Check if there is a second arguement input for banner choice
	if len(args) >= 2 {
		banner = args[1]
	}

	// Return user input along with nil error code
	return InputData{
		Text:   textInput,
		Banner: banner,
	}, nil
}
