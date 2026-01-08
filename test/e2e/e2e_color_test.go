package e2e

import (
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

// TestColorCLI tests that color flags work correctly via CLI
func TestColorCLI(t *testing.T) {
	tests := []struct {
		name        string
		args        []string
		expectError bool
		hasColor    bool
	}{
		{
			name:        "Color entire string - Red",
			args:        []string{"--color=red", "Hello"},
			expectError: false,
			hasColor:    true,
		},
		{
			name:        "Color substring - Blue",
			args:        []string{"--color=blue", "kit", "kitten"},
			expectError: false,
			hasColor:    true,
		},
		{
			name:        "Color with banner - Green",
			args:        []string{"--color=green", "Test", "shadow"},
			expectError: false,
			hasColor:    true,
		},
		{
			name:        "Hex color - Yellow",
			args:        []string{"--color=#FFFF00", "Go"},
			expectError: false,
			hasColor:    true,
		},
		{
			name:        "RGB color - Cyan",
			args:        []string{"--color=rgb(0,255,255)", "Art"},
			expectError: false,
			hasColor:    true,
		},
		{
			name:        "No color (backwards compatible)",
			args:        []string{"Plain"},
			expectError: false,
			hasColor:    false,
		},
		{
			name:        "Invalid flag format",
			args:        []string{"--color", "red", "test"},
			expectError: true,
			hasColor:    false,
		},
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("ASCII ART COLOR - E2E TESTS")
	fmt.Println(strings.Repeat("=", 80) + "\n")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := exec.Command("go", append([]string{"run", "./cmd/main.go"}, tt.args...)...)
			cmd.Dir = "../.."

			output, err := cmd.CombinedOutput()
			outputStr := string(output)

			// Display the test name and output
			fmt.Printf("\n▶ %s\n", tt.name)
			fmt.Printf("  Command: go run ./cmd %s\n\n", strings.Join(tt.args, " "))

			// Print the actual colored output (ANSI codes will render!)
			fmt.Print(outputStr)

			if !strings.HasSuffix(outputStr, "\n") {
				fmt.Println()
			}

			// Validation checks
			if tt.expectError {
				if err == nil && !strings.Contains(outputStr, "Usage") {
					t.Errorf("Expected error, but got success")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			hasANSI := strings.Contains(outputStr, "\033[")
			if tt.hasColor && !hasANSI {
				t.Errorf("Expected colored output, but no ANSI codes found")
				return
			}
			if !tt.hasColor && hasANSI {
				t.Errorf("Expected no color, but ANSI codes found")
				return
			}

			lines := strings.Split(outputStr, "\n")
			if !tt.expectError && len(lines) < 5 {
				t.Errorf("Expected ASCII art output with multiple lines, got %d lines", len(lines))
			}
		})
	}

	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("✅ ALL TESTS COMPLETE")
	fmt.Println(strings.Repeat("=", 80) + "\n")
}
