package asciioutput

import (
	"fmt"
	"os"
)

// WriteToFile writes the given content to the specified file
// If the file exists, it will be overwritten
// Returns error if file creation or writing fails
func WriteToFile(filename, content string) error {
	// Create orr truncate the file
	// 0644 permissions: owner can read/write, others can read
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %q: %w", filename, err)
	}
	defer file.Close()

	// Write to content
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("failed to write to file %q: %w", filename, err)
	}

	return nil
}

// FileExists checks if a file exists at the given path
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
