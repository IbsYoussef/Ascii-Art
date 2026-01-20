package asciijustify

import (
	"strings"
)

// ValidAlignTypes defines allowed alignment types
var ValidAlignTypes = map[string]bool{
	"left":    true,
	"center":  true,
	"right":   true,
	"justify": true,
}

// ParseAlignFlag extracts and validates the --align flag
// Returns: alignType, remainingArgs, error
func ParseAlignFlag(args []string) (string, []string, error) {
	for i, arg := range args {
		// Check for malformed flag (missing =)
		if arg == "--align" {
			return "", nil, ErrInvalidAlignFormat
		}

		// Check for properly formatted flag
		if strings.HasPrefix(arg, "--align=") {
			// Extract alignment type
			alignType := strings.TrimPrefix(arg, "--align=")

			// Check if empty
			if alignType == "" {
				return "", nil, ErrInvalidAlignFormat
			}

			// Validate alignment type
			if !ValidAlignTypes[alignType] {
				return "", nil, WrapAlignTypeError(alignType)
			}

			// Remove this arg from the list
			remaining := append([]string{}, args[:i]...)
			remaining = append(remaining, args[i+1:]...)

			return alignType, remaining, nil
		}
	}

	// No align flag found, default to left
	return "left", args, nil
}

// HasAlignFlag checks if --align flag exists in args
func HasAlignFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--align=") || arg == "--align" {
			return true
		}
	}

	return false
}

// ValidateAlignType checks if the alignment type is valid
func ValidateAlignType(alignType string) bool {
	return ValidAlignTypes[alignType]
}
