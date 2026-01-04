package ascii

import (
	"errors"
	"os"
	"strings"
)

var (
	errMissingInput = errors.New(`missing input. 
expected format: go run . "text input" <banner-style>
hint: wrap multi-word input in quotes`)

	errEmptyInput = errors.New(`text input cannot be empty.
example: go run . "Hello World" shadow`)

	errInvalidBanner = errors.New(`invalid banner style.
valid banners: standard, shadow, thinkertoy
note: if omitted, 'standard' will be used by default`)
)

var validBanners = map[string]bool{
	"standard":   true,
	"shadow":     true,
	"thinkertoy": true,
}

// GetUserInput validates and returns user input and banner
func GetUserInput() (string, string, error) {
	args := os.Args

	if len(args) < 2 {
		return "", "", errMissingInput
	}

	input := strings.TrimSpace(args[1])
	input = strings.ReplaceAll(input, "\\n", "\n")

	if input == "" {
		return "", "", errEmptyInput
	}

	banner := "standard"
	if len(args) >= 3 {
		banner = args[2]
		if !validBanners[banner] {
			return "", "", errInvalidBanner
		}
	}

	return input, banner, nil
}
