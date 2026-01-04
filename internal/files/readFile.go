package files

import (
	"bufio"
	"fmt"
	"os"
)

func ReadBannerLines(path string) ([]string, error) {
	// Attempt to open file
	file, err := os.Open(path)

	if err != nil {
		// Return nil slice and encountered error if opening the file fails
		return nil, err
	}
	defer file.Close() // Close the file when the function exits

	// Create a slice to hold the lines from the file
	var lines []string

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Append each line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Check for errors during file read
	if err := scanner.Err(); err != nil {
		// Return encountered error if there was an issue scanning the file
		return nil, err
	}

	if len(lines) == 0 {
		return nil, fmt.Errorf("file %s is empty", path)
	}

	// Return the slice of lines and nil error on success
	return lines, nil
}
