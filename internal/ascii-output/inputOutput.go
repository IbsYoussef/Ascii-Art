package asciioutput

import (
	"errors"
	"strings"
)

// ParseOutputFlag checks for --output flag and extracts filename
// Returns: outputFile (filename if flag present, empty otherwise),
//
//	remainingArgs (args without the output flag),
//	error (if flag format is invalid)
func ParseOutputFlag(args []string) (string, []string, error) {
	var outputFile string
	var remainingArgs []string

	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			// Extract filename from flag
			filename, err := ValidateOutputFlag(arg)
			if err != nil {
				return "", nil, err
			}
			outputFile = filename
		} else {
			// Keep non-output args
			remainingArgs = append(remainingArgs, arg)
		}
	}

	return outputFile, remainingArgs, nil
}

// ValidateOutputFlag ensures the flag is in correct format: --output=<fileName.txt>
// Returns the extracted filename or error if format is invalid
func ValidateOutputFlag(flag string) (string, error) {
	// Check prefix
	if !strings.HasPrefix(flag, "--output=") {
		return "", errors.New("invalid output flag format")
	}

	// Extract filename after "="
	parts := strings.SplitN(flag, "=", 2)
	if len(parts) != 2 {
		return "", errors.New("invalid output flag format: missing filename")
	}

	filename := parts[1]

	// Validate filename is not empty
	if strings.TrimSpace(filename) == "" {
		return "", errors.New("invalid output flag format: filename cannot be empty")
	}

	return filename, nil
}

// HasOutputFlag checks if --output flag exists in args
func HasOutputFlag(args []string) bool {
	for _, arg := range args {
		if strings.HasPrefix(arg, "--output=") {
			return true
		}
	}
	return false
}
