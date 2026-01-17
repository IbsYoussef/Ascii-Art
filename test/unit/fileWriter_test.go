package unit

import (
	output "ascii-art/internal/ascii-output"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	// Create temporary directory for test files
	tempDir := t.TempDir()

	tests := []struct {
		name        string
		filename    string
		content     string
		wantErr     bool
		errContains string
	}{
		{
			name:     "write simple content",
			filename: filepath.Join(tempDir, "test1.txt"),
			content:  "Hello, World!",
			wantErr:  false,
		},
		{
			name:     "write multiline content",
			filename: filepath.Join(tempDir, "test2.txt"),
			content:  "Line 1\nLine 2\nLine 3\n",
			wantErr:  false,
		},
		{
			name:     "write empty content",
			filename: filepath.Join(tempDir, "test3.txt"),
			content:  "",
			wantErr:  false,
		},
		{
			name:     "write ASCII art content",
			filename: filepath.Join(tempDir, "banner.txt"),
			content:  " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n",
			wantErr:  false,
		},
		{
			name:     "overwrite existing file",
			filename: filepath.Join(tempDir, "overwrite.txt"),
			content:  "New content",
			wantErr:  false,
		},
		{
			name:        "invalid directory path",
			filename:    "/nonexistent/directory/file.txt",
			content:     "test",
			wantErr:     true,
			errContains: "failed to create file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Special case: test overwriting by creating file first
			if tt.name == "overwrite existing file" {
				// Create file with initial content
				err := output.WriteToFile(tt.filename, "Initial content")
				if err != nil {
					t.Fatalf("Failed to create initial file: %v", err)
				}
			}

			// Execute WriteToFile
			err := output.WriteToFile(tt.filename, tt.content)

			// Check error expectations
			if tt.wantErr {
				if err == nil {
					t.Errorf("WriteToFile() expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !stringContains(err.Error(), tt.errContains) {
					t.Errorf("WriteToFile() error = %v, want error containing %q", err, tt.errContains)
				}
				return
			}

			// No error expected
			if err != nil {
				t.Errorf("WriteToFile() unexpected error = %v", err)
				return
			}

			// Verify file was created
			if !output.FileExists(tt.filename) {
				t.Errorf("WriteToFile() file %q was not created", tt.filename)
				return
			}

			// Read file content and verify
			gotContent, err := os.ReadFile(tt.filename)
			if err != nil {
				t.Errorf("Failed to read file %q: %v", tt.filename, err)
				return
			}

			if string(gotContent) != tt.content {
				t.Errorf("WriteToFile() file content = %q, want %q", string(gotContent), tt.content)
			}

			// Special case: verify overwrite worked
			if tt.name == "overwrite existing file" {
				if string(gotContent) == "Initial content" {
					t.Errorf("WriteToFile() did not overwrite file, still contains initial content")
				}
			}
		})
	}
}

func TestFileExists(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name     string
		filename string
		setup    func(string) error // Optional setup function to create file
		want     bool
	}{
		{
			name:     "file exists",
			filename: filepath.Join(tempDir, "exists.txt"),
			setup: func(path string) error {
				return os.WriteFile(path, []byte("test"), 0644)
			},
			want: true,
		},
		{
			name:     "file does not exist",
			filename: filepath.Join(tempDir, "notexists.txt"),
			setup:    nil,
			want:     false,
		},
		{
			name:     "directory exists (not a file)",
			filename: tempDir,
			setup:    nil,
			want:     true, // os.Stat returns nil for directories too
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run setup if provided
			if tt.setup != nil {
				err := tt.setup(tt.filename)
				if err != nil {
					t.Fatalf("Setup failed: %v", err)
				}
			}

			got := output.FileExists(tt.filename)
			if got != tt.want {
				t.Errorf("FileExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteToFilePermissions(t *testing.T) {
	tempDir := t.TempDir()
	filename := filepath.Join(tempDir, "permissions.txt")

	// Write file
	err := output.WriteToFile(filename, "test content")
	if err != nil {
		t.Fatalf("WriteToFile() failed: %v", err)
	}

	// Check file permissions
	info, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("Failed to stat file: %v", err)
	}

	// File should have 0644 permissions (owner: rw-, group: r--, others: r--)
	// On some systems this might be different due to umask
	mode := info.Mode()
	if mode.Perm()&0600 != 0600 {
		t.Errorf("File permissions = %v, want owner read/write permissions", mode.Perm())
	}
}
