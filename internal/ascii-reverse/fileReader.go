package asciireverse

import (
	output "ascii-art/internal/ascii-output"
	"os"
)

// ReadAsciiArtFile reads the ASCII art file and returns its content
// Returns the file content as a string or error if reading fails
func ReadAsciiArtFile(filename string) (string, error) {
	// Check if file exists
	if !output.FileExists(filename) {
		return "", WrapFileNotFoundError(filename)
	}

	// Read file contents
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", WrapFileReadError(filename, err)
	}

	return string(content), nil
}
