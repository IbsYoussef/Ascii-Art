package ascii

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadBannerFile(t *testing.T) {
	// Create a temporary file with test banner data
	tempDir := t.TempDir() // Creates a temporary directory for the test
	filepath := filepath.Join(tempDir, "test_banner.txt")

	content := `line1
line2
line3
line4
line5
line6
line7
line8

` // <- 1 char block with 8 lines + 1 separator (represents rune 32: ' ')

	// Write the file
	if err := os.WriteFile(filepath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write temp banner file: %v", err)
	}

	originalPath := "banners/teststyle.txt"
	if err := os.MkdirAll("banners", 0755); err != nil {
		t.Fatalf("Failed to create banners directory: %v", err)
	}
	defer os.RemoveAll("banners") // Clean up after test

	if err := os.WriteFile(originalPath, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to write teststyle.txt: %v", err)
	}

	// Call the function
	result, err := LoadBannerFile("teststyle")
	if err != nil {
		t.Fatalf("LoadBannerFile returned unexpected error: %v", err)
	}

	// Check the result
	if len(result) != 1 {
		t.Errorf("Expected 1 character in banner map, got %d", len(result))
	}

	r := rune(32) // ASCII 32 is space
	charLines, ok := result[r]
	if !ok {
		t.Errorf("Expected rune %q to be in banner map", r)
	}
	if len(charLines) != 8 {
		t.Errorf("Expected 8 lines for rune %q, got %d", r, len(charLines))
	}
}

func TestLoadBannerFile_FileNotFound(t *testing.T) {
	_, err := LoadBannerFile("nonexistent")
	if err == nil {
		t.Fatalf("Expected error for non-existent banner file, got nil")
	}
}
