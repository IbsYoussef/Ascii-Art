package asciijustify

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// HandleJustify orchestrates the justify alignment feature
// It captures rendered ASCII art output and applies alignment
func HandleJustify(renderFunc func(), alignType string) error {
	if alignType == "left" {
		// Left is default, no transformation needed
		renderFunc()
		return nil
	}

	// For other alignments, we need to capture the output
	// Save original stdout
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = w

	// Channel to capture output
	outputChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	// Execute render function
	renderFunc()

	// Close writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Get captured output
	output := <-outputChan

	// Split output into lines
	lines := strings.Split(output, "\n")

	// Remove trailing empty line if exists
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return nil
	}

	// Get terminal width
	termWidth := GetTerminalWidth()

	// Apply alignment
	alignedLines := ApplyAlignment(lines, alignType, termWidth)

	// Print aligned output
	for _, line := range alignedLines {
		fmt.Println(line)
	}

	return nil
}
