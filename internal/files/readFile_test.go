package files

import (
	"fmt"
	"os"
	"testing"
)

func TestBannerLines(t *testing.T) {
	// Create temporary file with test content
	tmpFile, err := os.CreateTemp("", "test_banner_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name()) // Clean up after test

	// Write known content to the file
	testContent := "Hello\nASCII\nArt"
	if _, err := tmpFile.WriteString(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close() // Close the file so we can read it

	// Run function call to read the file
	lines, err := ReadBannerLines(tmpFile.Name())
	if err != nil {
		t.Fatalf("ReadBannerLines failed: %v", err)
	}

	// Check results
	expectedLines := []string{"Hello", "ASCII", "Art"}

	if len(lines) != len(expectedLines) {
		t.Errorf("Expected %d lines, got %d", len(expectedLines), len(lines))
	}

	for i, line := range expectedLines {
		if lines[i] != line {
			t.Errorf("Expectd line %d to be %q, got %q", i, line, lines[i])
		}
	}
}

func TestReadBannerLines_EmptyFile(t *testing.T) {
	// Create empty temporary file
	tmpFile, err := os.CreateTemp("", "empty_banner_*.txt")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.Close()

	// Attempt to read file
	_, err = ReadBannerLines(tmpFile.Name())
	if err == nil {
		t.Fatalf("Expected error when reading empty file, got nil")
	}

	// Check error message
	expectedErr := fmt.Sprintf("file %s is empty", tmpFile.Name())
	if err.Error() != expectedErr {
		t.Errorf("Expected error %q, got %q", expectedErr, err.Error())
	}
}
