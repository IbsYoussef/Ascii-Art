package asciioutput

import "fmt"

// Usage message for the output feature
const UsageOutput = `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

// Error messages
var (
	// ErrInvalidOutputFormat is returned when the --output flag format is incorrect
	ErrInvalidOutputFormat = fmt.Errorf("invalid output flag format\n%s", UsageOutput)

	// ErrEmptyFilename is returned when the filename is empty or whitespace
	ErrEmptyFilename = fmt.Errorf("filename cannot be empty\n%s", UsageOutput)

	// ErrMissingFilename is returned when the --output= flag has no filename
	ErrMissingFilename = fmt.Errorf("missing filename after --output=\n%s", UsageOutput)
)

// WrapFileWriteError wraps file writing errors with additional context
func WrapFileWriteError(filename string, err error) error {
	return fmt.Errorf("failed to write to file %q: %w", filename, err)
}

// WrapFileCreateError wraps file creation errors with additional context
func WrapFileCreateError(filename string, err error) error {
	return fmt.Errorf("failed to create file %q: %w", filename, err)
}
