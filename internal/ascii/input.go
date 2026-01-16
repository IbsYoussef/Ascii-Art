package ascii

import (
	"errors"
	"os"
	"strings"
)

var (
	ErrMissingInput = errors.New(`missing input. 
expected format: go run . "text input" <banner-style>
hint: wrap multi-word input in quotes`)

	ErrEmptyInput = errors.New(`text input cannot be empty.
example: go run . "Hello World" shadow`)

	ErrInvalidBanner = errors.New(`invalid banner style.
valid banners: standard, shadow, thinkertoy
note: if omitted, 'standard' will be used by default`)
)

var ValidBanners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

// GetUserInput validates and returns user input and banner
func GetUserInput() (string, string, error) {
	// Get user input in terminal
	args := os.Args[1:]

	// Check if we have at least one argument (text input)
	if len(args) < 1 {
		return "", "", ErrMissingInput
	}

	// First argument is the text input
	input := strings.TrimSpace(args[0])
	input = strings.ReplaceAll(input, "\\n", "\n")

	if input == "" {
		return "", "", ErrEmptyInput
	}

	// Default banner is the standard banner
	banner := "standard"

	// If second argument exists, it's the banner
	if len(args) >= 2 {
		banner = args[1]
		if !ValidBanners[banner] {
			return "", "", ErrInvalidBanner
		}
	}

	return input, banner, nil
}
