package unit

import (
	output "ascii-art/internal/ascii-output"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestCaptureStdout(t *testing.T) {
	tests := []struct {
		name       string
		renderFunc output.RenderFunc
		want       string
	}{
		{
			name: "capture simple output",
			renderFunc: func() {
				fmt.Print("Hello, World!")
			},
			want: "Hello, World!",
		},
		{
			name: "capture multiline output",
			renderFunc: func() {
				fmt.Println("Line 1")
				fmt.Println("Line 2")
				fmt.Println("Line 3")
			},
			want: "Line 1\nLine 2\nLine 3\n",
		},
		{
			name: "capture empty output",
			renderFunc: func() {
				// Do nothing
			},
			want: "",
		},
		{
			name: "capture ASCII art output",
			renderFunc: func() {
				fmt.Print(" _    _          _   _          \n")
				fmt.Print("| |  | |        | | | |         \n")
				fmt.Print("| |__| |   ___  | | | |   ___   \n")
			},
			want: " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := output.CaptureStdout(tt.renderFunc)
			if err != nil {
				t.Errorf("CaptureStdout() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("CaptureStdout() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestHandleOutput(t *testing.T) {
	tempDir := t.TempDir()

	tests := []struct {
		name       string
		outputFile string
		renderFunc output.RenderFunc
		wantOutput string
		wantErr    bool
		checkFile  bool
	}{
		{
			name:       "output to stdout (no file)",
			outputFile: "",
			renderFunc: func() {
				fmt.Print("Direct to stdout")
			},
			wantOutput: "Direct to stdout",
			wantErr:    false,
			checkFile:  false,
		},
		{
			name:       "output to file",
			outputFile: filepath.Join(tempDir, "output1.txt"),
			renderFunc: func() {
				fmt.Print("Hello to file")
			},
			wantOutput: "Hello to file",
			wantErr:    false,
			checkFile:  true,
		},
		{
			name:       "output multiline to file",
			outputFile: filepath.Join(tempDir, "output2.txt"),
			renderFunc: func() {
				fmt.Println("Line 1")
				fmt.Println("Line 2")
			},
			wantOutput: "Line 1\nLine 2\n",
			wantErr:    false,
			checkFile:  true,
		},
		{
			name:       "output ASCII art to file",
			outputFile: filepath.Join(tempDir, "banner.txt"),
			renderFunc: func() {
				fmt.Print(" _              _   _          \n")
				fmt.Print("| |            | | | |         \n")
			},
			wantOutput: " _              _   _          \n| |            | | | |         \n",
			wantErr:    false,
			checkFile:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := output.HandleOutput(tt.outputFile, tt.renderFunc)

			if tt.wantErr {
				if err == nil {
					t.Errorf("HandleOutput() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("HandleOutput() unexpected error = %v", err)
				return
			}

			// If we should check the file, verify its contents
			if tt.checkFile {
				if !output.FileExists(tt.outputFile) {
					t.Errorf("HandleOutput() file %q was not created", tt.outputFile)
					return
				}

				content, err := os.ReadFile(tt.outputFile)
				if err != nil {
					t.Errorf("Failed to read output file: %v", err)
					return
				}

				if string(content) != tt.wantOutput {
					t.Errorf("HandleOutput() file content = %q, want %q", string(content), tt.wantOutput)
				}
			}
		})
	}
}

func TestHandleOutputInvalidPath(t *testing.T) {
	// Test with invalid file path
	renderFunc := func() {
		fmt.Print("test")
	}

	err := output.HandleOutput("/invalid/path/file.txt", renderFunc)
	if err == nil {
		t.Error("HandleOutput() expected error for invalid path, got nil")
	}
}

func TestCaptureStdoutRestoresStdout(t *testing.T) {
	// Verify that stdout is properly restored after capture
	originalStdout := os.Stdout

	renderFunc := func() {
		fmt.Print("test output")
	}

	_, err := output.CaptureStdout(renderFunc)
	if err != nil {
		t.Errorf("CaptureStdout() error = %v", err)
	}

	// Check that stdout was restored
	if os.Stdout != originalStdout {
		t.Error("CaptureStdout() did not restore os.Stdout")
	}
}
