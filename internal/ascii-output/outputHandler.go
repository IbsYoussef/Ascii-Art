package asciioutput

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// RenderFunc is a function type that renders ASCII art to stdout
type RenderFunc func()

// HandleOutput manages output routing - either to file or stdout
// If outputFile is empty, renders directly to stdout
// If outputFile is provided, captures stdout and writes to file
func HandleOutput(outputFile string, renderFunc RenderFunc) error {
	if outputFile == "" {
		// No output file specified, render directly to stdout
		renderFunc()
		return nil
	}

	// Capture output and write to file
	output, err := CaptureStdout(renderFunc)
	if err != nil {
		return fmt.Errorf("failed to capture output: %w", err)
	}

	// Write captured output to file
	err = WriteToFile(outputFile, output)
	if err != nil {
		return err
	}

	return nil
}

// CaptureStdout redirects stdout, executes the function, and returns captured output
func CaptureStdout(fn RenderFunc) (string, error) {
	// Save original stdout
	old := os.Stdout

	// Create pipe to capture output
	r, w, err := os.Pipe()
	if err != nil {
		return "", fmt.Errorf("failed to create pipe: %w", err)
	}

	// Redirect stdout to pipe writer
	os.Stdout = w

	// Execute the render function
	fn()

	// Close writer and restore stdout
	w.Close()
	os.Stdout = old

	// Read captured output from pipe
	var buf bytes.Buffer
	_, err = io.Copy(&buf, r)
	if err != nil {
		return "", fmt.Errorf("failed to read captured output: %w", err)
	}

	return buf.String(), nil
}
