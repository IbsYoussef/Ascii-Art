package unit_test

import (
	color "ascii-art/internal/ascii-color"
	"os"
	"testing"
)

func TestGetUserInputWithColor(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name              string
		args              []string
		expectedInput     string
		expectedBanner    string
		expectedColorFlag bool
		expectedColor     string
		expectedSubstring string
		expectedError     bool
	}{
		{
			name:              "Color entire string",
			args:              []string{"cmd", "--color=red", "hello"},
			expectedInput:     "hello",
			expectedBanner:    "standard",
			expectedColorFlag: true,
			expectedColor:     "red",
			expectedSubstring: "",
		},
		{
			name:              "Color substring",
			args:              []string{"cmd", "--color=blue", "B", "RGB()"},
			expectedInput:     "RGB()",
			expectedBanner:    "standard",
			expectedColorFlag: true,
			expectedColor:     "blue",
			expectedSubstring: "B",
		},
		{
			name:              "Color with banner",
			args:              []string{"cmd", "--color=green", "hello", "shadow"},
			expectedInput:     "hello",
			expectedBanner:    "shadow",
			expectedColorFlag: true,
			expectedColor:     "green",
			expectedSubstring: "",
		},
		{
			name:              "Color substring with banner",
			args:              []string{"cmd", "--color=yellow", "kit", "a king kitten", "thinkertoy"},
			expectedInput:     "a king kitten",
			expectedBanner:    "thinkertoy",
			expectedColorFlag: true,
			expectedColor:     "yellow",
			expectedSubstring: "kit",
		},
		{
			name:              "No color flag (backwards compatible)",
			args:              []string{"cmd", "hello", "shadow"},
			expectedInput:     "hello",
			expectedBanner:    "shadow",
			expectedColorFlag: false,
		},
		{
			name:          "Invalid color flag format (no equals)",
			args:          []string{"cmd", "--color", "red", "hello"},
			expectedError: true,
		},
		{
			name:          "Missing text after color flag",
			args:          []string{"cmd", "--color=red"},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			input, banner, colorConfig, err := color.GetUserInputWithColor()

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if input != tt.expectedInput {
				t.Errorf("Input = %q, want %q", input, tt.expectedInput)
			}

			if banner != tt.expectedBanner {
				t.Errorf("Banner = %q, want %q", banner, tt.expectedBanner)
			}

			if colorConfig.Enabled != tt.expectedColorFlag {
				t.Errorf("ColorEnabled = %v, want %v", colorConfig.Enabled, tt.expectedColorFlag)
			}

			if tt.expectedColorFlag {
				if colorConfig.Color != tt.expectedColor {
					t.Errorf("Color = %q, want %q", colorConfig.Color, tt.expectedColor)
				}

				if colorConfig.Substring != tt.expectedSubstring {
					t.Errorf("Substring = %q, want %q", colorConfig.Substring, tt.expectedSubstring)
				}
			}
		})
	}
}
