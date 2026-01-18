package unit

import (
	reverse "ascii-art/internal/ascii-reverse"
	"testing"
)

func TestValidateReverseFlag(t *testing.T) {
	tests := []struct {
		name        string
		flag        string
		wantFile    string
		wantErr     bool
		errContains string
	}{
		{
			name:     "valid flag with txt extension",
			flag:     "--reverse=file.txt",
			wantFile: "file.txt",
			wantErr:  false,
		},
		{
			name:     "valid flag with path",
			flag:     "--reverse=path/to/file.txt",
			wantFile: "path/to/file.txt",
			wantErr:  false,
		},
		{
			name:     "valid flag with different extension",
			flag:     "--reverse=output.out",
			wantFile: "output.out",
			wantErr:  false,
		},
		{
			name:        "empty filename",
			flag:        "--reverse=",
			wantFile:    "",
			wantErr:     true,
			errContains: "filename cannot be empty",
		},
		{
			name:        "whitespace filename",
			flag:        "--reverse=   ",
			wantFile:    "",
			wantErr:     true,
			errContains: "filename cannot be empty",
		},
		{
			name:        "missing equals sign",
			flag:        "--reverse",
			wantFile:    "",
			wantErr:     true,
			errContains: "invalid reverse flag format",
		},
		{
			name:        "wrong prefix",
			flag:        "-reverse=file.txt",
			wantFile:    "",
			wantErr:     true,
			errContains: "invalid reverse flag format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, err := reverse.ValidateReverseFlag(tt.flag)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ValidateReverseFlag() expected error containing %q, got nil", tt.errContains)
					return
				}
				if tt.errContains != "" && !contains(err.Error(), tt.errContains) {
					t.Errorf("ValidateReverseFlag() error = %v, want error containing %q", err, tt.errContains)
				}
			} else {
				if err != nil {
					t.Errorf("ValidateReverseFlag() unexpected error = %v", err)
					return
				}
				if gotFile != tt.wantFile {
					t.Errorf("ValidateReverseFlag() = %q, want %q", gotFile, tt.wantFile)
				}
			}
		})
	}
}

func TestParseReverseFlag(t *testing.T) {
	tests := []struct {
		name             string
		args             []string
		wantFilename     string
		wantRemainingLen int
		wantRemaining    []string
		wantErr          bool
	}{
		{
			name:             "valid reverse flag",
			args:             []string{"--reverse=file.txt"},
			wantFilename:     "file.txt",
			wantRemainingLen: 0,
			wantRemaining:    []string{},
			wantErr:          false,
		},
		{
			name:             "valid reverse flag with other args",
			args:             []string{"--reverse=art.txt", "extra", "args"},
			wantFilename:     "art.txt",
			wantRemainingLen: 2,
			wantRemaining:    []string{"extra", "args"},
			wantErr:          false,
		},
		{
			name:             "no reverse flag",
			args:             []string{"hello", "standard"},
			wantFilename:     "",
			wantRemainingLen: 2,
			wantRemaining:    []string{"hello", "standard"},
			wantErr:          false,
		},
		{
			name:             "reverse flag with color flag",
			args:             []string{"--reverse=file.txt", "--color=red"},
			wantFilename:     "file.txt",
			wantRemainingLen: 1,
			wantRemaining:    []string{"--color=red"},
			wantErr:          false,
		},
		{
			name:             "reverse flag at end",
			args:             []string{"some", "args", "--reverse=test.txt"},
			wantFilename:     "test.txt",
			wantRemainingLen: 2,
			wantRemaining:    []string{"some", "args"},
			wantErr:          false,
		},
		{
			name:             "invalid format - no equals",
			args:             []string{"--reverse", "file.txt"},
			wantFilename:     "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
		{
			name:             "invalid format - empty filename",
			args:             []string{"--reverse="},
			wantFilename:     "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
		{
			name:             "invalid format - single dash",
			args:             []string{"-reverse=file.txt"},
			wantFilename:     "",
			wantRemainingLen: 0,
			wantRemaining:    nil,
			wantErr:          true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFile, gotRemaining, err := reverse.ParseReverseFlag(tt.args)

			if tt.wantErr {
				if err == nil {
					t.Errorf("ParseReverseFlag() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("ParseReverseFlag() unexpected error = %v", err)
				return
			}

			if gotFile != tt.wantFilename {
				t.Errorf("ParseReverseFlag() filename = %q, want %q", gotFile, tt.wantFilename)
			}

			if len(gotRemaining) != tt.wantRemainingLen {
				t.Errorf("ParseReverseFlag() remaining args length = %d, want %d", len(gotRemaining), tt.wantRemainingLen)
			}

			if !equalSlices(gotRemaining, tt.wantRemaining) {
				t.Errorf("ParseReverseFlag() remaining args = %v, want %v", gotRemaining, tt.wantRemaining)
			}
		})
	}
}

func TestHasReverseFlag(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "has reverse flag",
			args: []string{"--reverse=file.txt"},
			want: true,
		},
		{
			name: "has reverse flag with other args",
			args: []string{"--reverse=file.txt", "hello"},
			want: true,
		},
		{
			name: "no reverse flag",
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
			got := reverse.HasReverseFlag(tt.args)
			if got != tt.want {
				t.Errorf("HasReverseFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
