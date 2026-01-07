package e2e

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestAsciiArtEndToEnd(t *testing.T) {
	// Set up paths
	cmdPath := filepath.Join("..", "..", "cmd", "main.go")
	outputFile := "test_output.txt"
	defer os.Remove(outputFile) // Clean up after test

	cmd := exec.Command("go", "run", cmdPath, "A")

	outFile, err := os.Create(outputFile)
	if err != nil {
		t.Fatal("Failed to create output file:", err)
	}

	var stderr bytes.Buffer
	cmd.Stdout = outFile
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		t.Fatalf("Program failed: %v\nStderr: %s", err, stderr.String())
	}

	// Readback and check the output
	content, err := os.ReadFile(outputFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}

	if len(content) == 0 {
		t.Error("Output file is empty, expected ASCII art output")
	}
}
