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
	if len(args) < 1 {
		return InputData{}, fmt.Errorf(
			"Usage: go run . \"<text-input>\" [banner-name]\nNote: Enclose text in quotes if it contains spaces",
		)
	}

	textInput := strings.ReplaceAll(args[0], `\n`, "\n")
	if textInput == "" {
		return InputData{}, fmt.Errorf("error: Text input cannot be empty")
	}

	banner := "standard"
	if len(args) >= 2 {
		banner = args[1]
	}

	return InputData{
		Text:   textInput,
		Banner: banner,
	}, nil
}
