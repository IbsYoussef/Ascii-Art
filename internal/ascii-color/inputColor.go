package ascii

import (
	"ascii-art/internal/ascii"
	"errors"
	"os"
	"strings"
)

var errInvalidColorFlag = errors.New(`Usage: go run ./cmd [OPTION] [STRING]
EX: go run ./cmd --color=<color> <substring to be colored> "something"`)

// ColorConfig holds color configuration
type ColorConfig struct {
	Enabled   bool   // Whether color is enabled
	Color     string // The color (parsed later by ParseColor)
	Substring string // Substring to color (empty = color entire string)
}

// GetUserInputWithColor parses arguments including color flags
// Returns: input text, banner name, color config, error
func GetUserInputWithColor() (string, string, ColorConfig, error) {
	args := os.Args[1:]

	// No arguments at all
	if len(args) == 0 {
		return "", "", ColorConfig{}, ascii.ErrMissingInput
	}

	// Check if first argument is a color flag (with or without =)
	if strings.HasPrefix(args[0], "--color") {
		// Check if it has the correct format (with =)
		if !strings.HasPrefix(args[0], "--color=") {
			// Wrong format: --color red instead of --color=red
			return "", "", ColorConfig{}, errInvalidColorFlag
		}
		return parseWithColorFlag(args)
	}

	// No color flag - use original parser
	input, banner, err := ascii.GetUserInput()
	return input, banner, ColorConfig{Enabled: false}, err
}

// parseWithColorFlag handles arguments when --color flag is present
func parseWithColorFlag(args []string) (string, string, ColorConfig, error) {
	// Minimum: --color=<color> "text"
	if len(args) < 2 {
		return "", "", ColorConfig{}, errInvalidColorFlag
	}

	// Extract color from flag
	colorFlag := args[0]
	if !strings.HasPrefix(colorFlag, "--color=") {
		return "", "", ColorConfig{}, errInvalidColorFlag
	}

	color := strings.TrimPrefix(colorFlag, "--color=")
	if color == "" {
		return "", "", ColorConfig{}, errInvalidColorFlag
	}

	// Determine argument structure
	// Case 1: --color=red "text"
	// Case 2: --color=red "text" banner
	// Case 3: --color=red substring "text"
	// Case 4: --color=red substring "text" banner

	var input, substring, banner string

	// Check if last argument is a valid banner
	lastArg := args[len(args)-1]
	hasBanner := ascii.ValidBanners[lastArg]

	if hasBanner {
		// Banner is specified
		banner = lastArg
		args = args[:len(args)-1] // Remove banner from args
	} else {
		banner = "standard"
	}

	// Now args is: [--color=X, ...other args without banner]
	// We need at least the text input (args[1])
	if len(args) < 2 {
		return "", "", ColorConfig{}, errInvalidColorFlag
	}

	// Check if we have substring or not
	if len(args) == 2 {
		// --color=red "text" [banner]
		input = args[1]
		substring = "" // Color entire string
	} else if len(args) == 3 {
		// --color=red substring "text" [banner]
		substring = args[1]
		input = args[2]
	} else {
		// Too many arguments
		return "", "", ColorConfig{}, errInvalidColorFlag
	}

	// Process input (handle \n and trim)
	input = strings.TrimSpace(input)
	input = strings.ReplaceAll(input, "\\n", "\n")

	if input == "" {
		return "", "", ColorConfig{}, ascii.ErrMissingInput
	}

	// Validate color format (try to parse it)
	_, err := ParseColor(color)
	if err != nil {
		return "", "", ColorConfig{}, err
	}

	colorConfig := ColorConfig{
		Enabled:   true,
		Color:     color,
		Substring: substring,
	}

	return input, banner, colorConfig, nil
}
