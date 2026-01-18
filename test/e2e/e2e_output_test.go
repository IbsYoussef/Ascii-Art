package e2e

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestOutputCLI tests that --output flag works correctly via CLI
func TestOutputCLI(t *testing.T) {
	// Create temporary directory for test output files
	tempDir := t.TempDir()

	tests := []struct {
		name         string
		args         []string
		outputFile   string
		expectError  bool
		validateFile bool
		description  string
	}{
		{
			name:         "Invalid format - missing equals",
			args:         []string{"--output", "test00.txt", "banana", "standard"},
			outputFile:   "",
			expectError:  true,
			validateFile: false,
			description:  "Should show usage when --output missing '='",
		},
		{
			name:         "Valid - First\\nTest with shadow",
			args:         []string{"--output=" + filepath.Join(tempDir, "test00.txt"), "First\\nTest", "shadow"},
			outputFile:   filepath.Join(tempDir, "test00.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Multiline text with shadow banner",
		},
		{
			name:         "Valid - hello with standard",
			args:         []string{"--output=" + filepath.Join(tempDir, "test01.txt"), "hello", "standard"},
			outputFile:   filepath.Join(tempDir, "test01.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Simple text with standard banner",
		},
		{
			name:         "Valid - special characters",
			args:         []string{"--output=" + filepath.Join(tempDir, "test02.txt"), "123 -> #$%", "standard"},
			outputFile:   filepath.Join(tempDir, "test02.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Numbers and special characters",
		},
		{
			name:         "Valid - shadow with special chars",
			args:         []string{"--output=" + filepath.Join(tempDir, "test03.txt"), "432 -> #$%&@", "shadow"},
			outputFile:   filepath.Join(tempDir, "test03.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Complex special characters with shadow",
		},
		{
			name:         "Valid - There with shadow",
			args:         []string{"--output=" + filepath.Join(tempDir, "test04.txt"), "There", "shadow"},
			outputFile:   filepath.Join(tempDir, "test04.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Single word with shadow banner",
		},
		{
			name:         "Valid - thinkertoy with quotes",
			args:         []string{"--output=" + filepath.Join(tempDir, "test05.txt"), "123 -> \"#$%@", "thinkertoy"},
			outputFile:   filepath.Join(tempDir, "test05.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Special chars with quotes - thinkertoy",
		},
		{
			name:         "Valid - 2 you with thinkertoy",
			args:         []string{"--output=" + filepath.Join(tempDir, "test06.txt"), "2 you", "thinkertoy"},
			outputFile:   filepath.Join(tempDir, "test06.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Number and text with thinkertoy",
		},
		{
			name:         "Valid - long output",
			args:         []string{"--output=" + filepath.Join(tempDir, "test07.txt"), "Testing long output!", "standard"},
			outputFile:   filepath.Join(tempDir, "test07.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Long text string with standard",
		},
		{
			name:         "Valid - with color flag",
			args:         []string{"--output=" + filepath.Join(tempDir, "test08.txt"), "--color=red", "Colored", "standard"},
			outputFile:   filepath.Join(tempDir, "test08.txt"),
			expectError:  false,
			validateFile: true,
			description:  "Output flag combined with color flag",
		},
		{
			name:         "Invalid - empty filename",
			args:         []string{"--output=", "test", "standard"},
			outputFile:   "",
			expectError:  true,
			validateFile: false,
			description:  "Should show usage for empty filename",
		},
		{
			name:         "Invalid - single dash",
			args:         []string{"-output=test.txt", "test", "standard"},
			outputFile:   "",
			expectError:  true,
			validateFile: false,
			description:  "Should show usage for single dash",
		},
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("ASCII ART OUTPUT - E2E TESTS")
	fmt.Println(strings.Repeat("=", 80) + "\n")

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build command
			cmd := exec.Command("go", append([]string{"run", "./cmd/main.go"}, tt.args...)...)
			cmd.Dir = "../.."

			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			// Display test header
			fmt.Printf("\n┌─ Test %d: %s\n", i+1, tt.name)
			fmt.Printf("│  %s\n", tt.description)
			fmt.Printf("│  Command: go run . %s\n", strings.Join(tt.args, " "))
			fmt.Println("└" + strings.Repeat("─", 78))

			// Handle error cases
			if tt.expectError {
				if err == nil && !strings.Contains(outputStr, "Usage") {
					fmt.Println("❌ FAIL: Expected error or usage message")
					t.Errorf("Expected error, but got success")
					return
				}
				fmt.Println("\n📋 Output (Usage Message):")
				fmt.Println(strings.TrimSpace(outputStr))
				fmt.Println("\n✅ PASS: Correct error handling")
				return
			}

			// Check for unexpected errors
			if err != nil {
				fmt.Printf("\n❌ FAIL: Unexpected error: %v\n", err)
				fmt.Println("Output:", outputStr)
				t.Errorf("Unexpected error: %v", err)
				return
			}

			// Validate file creation and content
			if tt.validateFile {
				// Check file exists
				if !fileExists(tt.outputFile) {
					fmt.Printf("\n❌ FAIL: Output file not created: %s\n", tt.outputFile)
					t.Errorf("Output file %q was not created", tt.outputFile)
					return
				}

				// Read file content
				content, err := os.ReadFile(tt.outputFile)
				if err != nil {
					fmt.Printf("\n❌ FAIL: Could not read output file: %v\n", err)
					t.Errorf("Failed to read output file: %v", err)
					return
				}

				contentStr := string(content)

				// Display file content preview
				fmt.Println("\n📄 File Created:", filepath.Base(tt.outputFile))
				fmt.Println("📏 File Size:", len(contentStr), "bytes")
				fmt.Println("\n📝 Content Preview (first 10 lines):")
				fmt.Println(strings.Repeat("-", 78))

				lines := strings.Split(contentStr, "\n")
				previewLines := 10
				if len(lines) < previewLines {
					previewLines = len(lines)
				}
				for j := 0; j < previewLines; j++ {
					fmt.Println(lines[j])
				}

				if len(lines) > previewLines {
					fmt.Printf("... (%d more lines)\n", len(lines)-previewLines)
				}
				fmt.Println(strings.Repeat("-", 78))

				// Validation checks
				if len(contentStr) == 0 {
					fmt.Println("\n❌ FAIL: Output file is empty")
					t.Errorf("Output file is empty")
					return
				}

				// Check if it contains ASCII art (should have multiple lines)
				if len(lines) < 5 {
					fmt.Printf("\n❌ FAIL: Expected ASCII art with multiple lines, got %d lines\n", len(lines))
					t.Errorf("Expected ASCII art output with multiple lines, got %d lines", len(lines))
					return
				}

				// Check for ANSI codes if color flag was used
				if strings.Contains(strings.Join(tt.args, " "), "--color") {
					if !strings.Contains(contentStr, "\033[") && !strings.Contains(contentStr, "[38;2;") {
						fmt.Println("\n⚠️  WARNING: Color flag used but no ANSI codes found in output")
					} else {
						fmt.Println("\n🎨 Color codes detected in output file")
					}
				}

				fmt.Println("\n✅ PASS: File created with valid ASCII art content")
			}

			fmt.Println()
		})
	}

	fmt.Println(strings.Repeat("=", 80))
	fmt.Println("✅ ALL OUTPUT TESTS COMPLETE")
	fmt.Println(strings.Repeat("=", 80) + "\n")
}

// fileExists checks if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

// TestOutputBackwardsCompatibility ensures program still works without --output flag
func TestOutputBackwardsCompatibility(t *testing.T) {
	tests := []struct {
		name string
		args []string
	}{
		{
			name: "Simple text - no output flag",
			args: []string{"Hello"},
		},
		{
			name: "Text with banner - no output flag",
			args: []string{"World", "shadow"},
		},
		{
			name: "With color - no output flag",
			args: []string{"--color=blue", "Test"},
		},
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("ASCII ART OUTPUT - BACKWARDS COMPATIBILITY TESTS")
	fmt.Println(strings.Repeat("=", 80) + "\n")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "./cmd/main.go"}, tt.args...)...)
			cmd.Dir = "../.."

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Printf("❌ FAIL: %s\n", tt.name)
				t.Errorf("Command failed: %v", err)
				return
			}

			outputStr := string(output)
			if len(strings.Split(outputStr, "\n")) < 5 {
				fmt.Printf("❌ FAIL: %s - insufficient output\n", tt.name)
				t.Errorf("Expected ASCII art output")
				return
			}

			fmt.Printf("✅ PASS: %s\n", tt.name)
		})
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("✅ BACKWARDS COMPATIBILITY VERIFIED")
	fmt.Println(strings.Repeat("=", 80) + "\n")
}
