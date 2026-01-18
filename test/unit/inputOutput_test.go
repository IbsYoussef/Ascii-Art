package unit

import (
	output "ascii-art/internal/ascii-output"
	"testing"
)

func TestValidateOutputFlag(t *testing.T) {
	tests := []struct {
		name        string
		flag        string
		wantFile    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid flag with txt extension",
			flag:     "--output=banner.txt",
			wantFile: "banner.txt",
			wantErr:  false,
		},
		{
			name:     "valid flag with different extension",
			flag:     "--output=result.out",
			wantFile: "result.out",
			wantErr:  false,
		},
		{
			name:     "valid flag with path",
			flag:     "--output=outputs/banner.txt",
			wantFile: "outputs/banner.txt",
			wantErr:  false,
		},
		{
			name:        "empty filename",
			flag:        "--output=",
			wantFile:    "",
			wantErr:     true,
			errContains: "filename cannot be empty",
		},
		{
			name:        "whitespace filename",
			flag:        "--output=   ",
			wantFile:    "",
			wantErr:     true,
			errContains: "filename cannot be empty",
		},
		{
			name:        "missing equals sign",
			flag:        "--output",
			wantFile:    "",
			wantErr:     true,
			errContains: "invalid output flag format",
		},
		{
			name:        "wrong prefix",
			flag:        "-output=file.txt",
			wantFile:    "",
			wantErr:     true,
			errContains: "invalid output flag format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, err := output.ValidateOutputFlag(tt.flag)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidateOutputFlag() expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("ValidateOutputFlag() error = %v, want error containing %q", err, tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateOutputFlag() unexpected error = %v", err)
					return
				}
				if gotFile != tt.wantFile {
					t.Errorf("ValidateOutputFlag() = %q, want %q", gotFile, tt.wantFile)
				}
			}
		})
	}
}

func TestParseOutputFlag(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		wantOutputFile   string
		wantRemainingLen int
		wantRemaining    []string
		wantErr          bool
	}{
		{
			name:             "valid output flag with text and banner",
			args:             []string{"--output=banner.txt", "hello", "standard"},
			wantOutputFile:   "banner.txt",
			wantRemainingLen: 2,
			wantRemaining:    []string{"hello", "standard"},
			wantErr:          false,
		},
		{
			name:             "no output flag",
			args:             []string{"hello", "standard"},
			wantOutputFile:   "",
			wantRemainingLen: 2,
			wantRemaining:    []string{"hello", "standard"},
			wantErr:          false,
		},
		{
			name:             "output flag with color flag",
			args:             []string{"--output=result.txt", "--color=red", "hello", "standard"},
			wantOutputFile:   "result.txt",
			wantRemainingLen: 3,
			wantRemaining:    []string{"--color=red", "hello", "standard"},
			wantErr:          false,
		},
		{
			name:             "output flag at end",
			args:             []string{"hello", "standard", "--output=banner.txt"},
			wantOutputFile:   "banner.txt",
			wantRemainingLen: 2,
			wantRemaining:    []string{"hello", "standard"},
			wantErr:          false,
		},
		{
			name:             "output flag only",
			args:             []string{"--output=test.txt"},
			wantOutputFile:   "test.txt",
			wantRemainingLen: 0,
			wantRemaining:    []string{},
			wantErr:          false,
		},
		{
			name:             "invalid output flag format - empty filename",
			args:             []string{"--output=", "hello"},
			wantOutputFile:   "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
		{
			name:             "invalid output flag format - no equals sign",
			args:             []string{"--output", "test.txt", "hello"},
			wantOutputFile:   "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
		{
			name:             "invalid output flag format - space instead of equals",
			args:             []string{"--output test.txt", "hello"},
			wantOutputFile:   "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
		{
			name:             "invalid output flag format - single dash",
			args:             []string{"-output=test.txt", "hello"},
			wantOutputFile:   "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, gotRemaining, err := output.ParseOutputFlag(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseOutputFlag() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("ParseOutputFlag() unexpected error = %v", err)
				return
			}

			if gotFile != tt.wantOutputFile {
				t.Errorf("ParseOutputFlag() outputFile = %q, want %q", gotFile, tt.wantOutputFile)
			}

			if len(gotRemaining) != tt.wantRemainingLen {
				t.Errorf("ParseOutputFlag() remaining args length = %d, want %d", len(gotRemaining), tt.wantRemainingLen)
			}

			if !equalSlices(gotRemaining, tt.wantRemaining) {
				t.Errorf("ParseOutputFlag() remaining args = %v, want %v", gotRemaining, tt.wantRemaining)
			}
		})
	}
}

func TestHasOutputFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "has output flag",
			args: []string{"--output=file.txt", "hello"},
			want: true,
		},
		{
			name: "no output flag",
			args: []string{"hello", "standard"},
			want: false,
		},
		{
			name: "has color flag only",
			args: []string{"--color=red", "hello"},
			want: false,
		},
		{
			name: "empty args",
			args: []string{},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := output.HasOutputFlag(tt.args)
			if got != tt.want {
				t.Errorf("HasOutputFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && stringContains(s, substr)))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// Helper function to compare slices
func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
