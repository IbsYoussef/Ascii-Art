package asciioutput

import (
	"os"
)

// WriteToFile writes the given content to the specified file
// If the file exists, it will be overwritten
// Returns error if file creation or writing fails
func WriteToFile(filename, content string) error {
	// Create or truncate the file
	// 0644 permissions: owner can read/write, others can read
	file, err := os.Create(filename)
	if err != nil {
		return WrapFileCreateError(filename, err)
	}
	defer file.Close()

	// Write content to file
	_, err = file.WriteString(content)
	if err != nil {
		return WrapFileWriteError(filename, err)
	}

	return nil
}

// FileExists checks if a file exists at the given path
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}
