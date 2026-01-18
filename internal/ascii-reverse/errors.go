package asciireverse

import "fmt"

// Usage message for the reverse feature
const UsageReverse = `Usage: go run ./cmd [OPTION]

EX: go run ./cmd --reverse=<fileName>`

// Error messages
var (
	// ErrInvalidReverseFormat is returned when the --reverse flag format is incorrect
	ErrInvalidReverseFormat = fmt.Errorf("invalid reverse flag format\n%s", UsageReverse)

	// ErrEmptyFilename is returned when the filename is empty or whitespace
	ErrEmptyFilename = fmt.Errorf("filename cannot be empty\n%s", UsageReverse)

	// ErrMissingFilename is returned when the --reverse= flag has no filename
	ErrMissingFilename = fmt.Errorf("missing filename after --reverse=\n%s", UsageReverse)
)

// WrapFileReadError wraps file reading errors with additional context
func WrapFileReadError(filename string, err error) error {
	return fmt.Errorf("failed to read file %q: %w", filename, err)
}

// WrapFileNotFoundError wraps file not found errors
func WrapFileNotFoundError(filename string) error {
	return fmt.Errorf("file not found: %q", filename)
}
