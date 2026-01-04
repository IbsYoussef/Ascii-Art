package ascii

import (
	"os"
	"testing"
)

func TestGetUserInput(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	testCases := []struct {
		name           string
		args           []string // Simulates full os.Args (includes program name)
		expectedInput  string
		expectedBanner string
		expectedError  error
	}{
		{
			name:          "No input provided",
			args:          []string{"cmd"}, // Only program name, no arguments
			expectedError: errMissingInput,
		},
		{
			name:          "Empty input string",
			args:          []string{"cmd", ""}, // Program name + empty string
			expectedError: errEmptyInput,
		},
		{
			name:           "Valid input with default banner",
			args:           []string{"cmd", "Hello"}, // Program name + text input
			expectedInput:  "Hello",
			expectedBanner: "standard",
		},
		{
			name:           "Valid input with shadow banner",
			args:           []string{"cmd", "Hi", "shadow"}, // Program name + text + banner
			expectedInput:  "Hi",
			expectedBanner: "shadow",
		},
		{
			name:          "Invalid banner style",
			args:          []string{"cmd", "Hey", "unknown"}, // Invalid banner name
			expectedError: errInvalidBanner,
		},
		{
			name:           "Escaped newline in input",
			args:           []string{"cmd", "Hello\\nWorld"}, // \\n should become \n
			expectedInput:  "Hello\nWorld",
			expectedBanner: "standard",
		},
		{
			name:           "Input with leading and trailing spaces",
			args:           []string{"cmd", "   Hello   "}, // Spaces should be trimmed
			expectedInput:  "Hello",
			expectedBanner: "standard",
		},
		{
			name:           "Input without banner defaults to standard",
			args:           []string{"cmd", "Hello World"},
			expectedInput:  "Hello World",
			expectedBanner: "standard",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Override os.Args with simulated CLI input (includes program name)
			os.Args = tc.args

			// Call the function under test
			input, banner, err := GetUserInput()

			// Check expected error
			if tc.expectedError != nil {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("Expected error %q, got %v", tc.expectedError, err)
				}
				return // Skip further checks if error was expected
			}

			// If no error expected, check values
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if input != tc.expectedInput {
				t.Errorf("Expected input %q, got %q", tc.expectedInput, input)
			}
			if banner != tc.expectedBanner {
				t.Errorf("Expected banner %q, got %q", tc.expectedBanner, banner)
			}
		})
	}
}
