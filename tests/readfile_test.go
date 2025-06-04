package tests

import (
	"ascii-art/util"
	"os"
	"testing"
)

func TestReadFile_ValidFile(t *testing.T) {
	tmpFile := "testfile.txt"
	content := []byte("Hello, world!")

	err := os.WriteFile(tmpFile, content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tmpFile)

	data, err := util.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if string(data) != string(content) {
		t.Errorf("Expected %q, got %q", content, data)
	}
}

func TestReadFile_NonExistent(t *testing.T) {
	_, err := util.ReadFile("no_such_file.txt")
	if err == nil {
		t.Error("Expected an error for non-existent file, got nil")
	}
}
