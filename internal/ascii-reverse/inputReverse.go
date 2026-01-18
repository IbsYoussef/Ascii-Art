package asciireverse

import (
	"strings"
)

// ParseReverseFlag checks for --reverse flag and extracts filename
// Returns: filename (if flag present, empty otherwise),
//
//	remainingArgs (args without the reverse flag),
//	error (if flag format is invalid)
func ParseReverseFlag(args []string) (string, []string, error) {
	var filename string
	var remainingArgs []string

	for _, arg := range args {
		// Check for valid --reverse= format
		if strings.HasPrefix(arg, "--reverse=") {
			// Extract filename from flag
			file, err := ValidateReverseFlag(arg)
			if err != nil {
				return "", nil, err
			}
			filename = file
		} else if arg == "--reverse" || (strings.HasPrefix(arg, "--reverse") && !strings.HasPrefix(arg, "--reverse=")) {
			// Catches: --reverse (no value) or --reverse<anything-without-equals>
			return "", nil, ErrInvalidReverseFormat
		} else if strings.HasPrefix(arg, "-reverse=") {
			// Catches: -reverse= (single dash)
			return "", nil, ErrInvalidReverseFormat
		} else {
			// Keep non-reverse args
			remainingArgs = append(remainingArgs, arg)
		}
	}

	return filename, remainingArgs, nil
}

// ValidateReverseFlag ensures the flag is in correct format: --reverse=<fileName>
// Returns the extracted filename or error if format is invalid
func ValidateReverseFlag(flag string) (string, error) {
	// Check prefix
	if !strings.HasPrefix(flag, "--reverse=") {
		return "", ErrInvalidReverseFormat
	}

	// Extract filename after "="
	parts := strings.SplitN(flag, "=", 2)
	if len(parts) != 2 {
		return "", ErrMissingFilename
	}

	filename := parts[1]

	// Validate filename is not empty
	if strings.TrimSpace(filename) == "" {
		return "", ErrEmptyFilename
	}

	return filename, nil
}

// HasReverseFlag checks if --reverse flag exists in args
func HasReverseFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--reverse=") {
			return true
		}
	}
	return false
}
