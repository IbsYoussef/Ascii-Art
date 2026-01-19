package unit

import (
	reverse "ascii-art/internal/ascii-reverse"
	"testing"
)

func TestRecogniseCharacter(t *testing.T) {
	// Create simple templates for testing
	templates := map[rune][]string{
		' ': {
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
		},
		'!': {
			" _  ",
			"| | ",
			"| | ",
			"| | ",
			"|_| ",
			"(_) ",
			"    ",
			"    ",
		},
		'a': {
			"       ",
			"       ",
			"  __ _ ",
			" / _` |",
			"| (_| |",
			" \\__,_|",
			"       ",
			"       ",
		},
	}

	tests := []struct {
		name    string
		pattern []string
		want    rune
		wantErr bool
	}{
		{
			name: "recognize exclamation mark",
			pattern: []string{
				" _  ",
				"| | ",
				"| | ",
				"| | ",
				"|_| ",
				"(_) ",
				"    ",
				"    ",
			},
			want:    '!',
			wantErr: false,
		},
		{
			name: "recognize space",
			pattern: []string{
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
			},
			want:    ' ',
			wantErr: false,
		},
		{
			name: "recognize lowercase a",
			pattern: []string{
				"       ",
				"       ",
				"  __ _ ",
				" / _` |",
				"| (_| |",
				" \\__,_|",
				"       ",
				"       ",
			},
			want:    'a',
			wantErr: false,
		},
		{
			name: "unrecognized pattern",
			pattern: []string{
				"xxxxx",
				"xxxxx",
				"xxxxx",
				"xxxxx",
				"xxxxx",
				"xxxxx",
				"xxxxx",
				"xxxxx",
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverse.RecogniseCharacter(tt.pattern, templates)

			if tt.wantErr {
				if err == nil {
					t.Errorf("RecognizeCharacter() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("RecognizeCharacter() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("RecognizeCharacter() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRecogniseText(t *testing.T) {
	// Create templates for testing - ALL MUST HAVE SAME WIDTH (6 chars)
	templates := map[rune][]string{
		' ': {
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
			"      ",
		},
		'h': {
			" _    ",
			"| |   ",
			"| |__ ",
			"| '_ \\",
			"| | | ",
			"|_| |_",
			"      ",
			"      ",
		},
		'i': {
			" _    ",
			"(_)   ",
			" |    ",
			" |    ",
			" |    ",
			"|_|   ",
			"      ",
			"      ",
		},
	}

	tests := []struct {
		name    string
		chunks  [][]string
		want    string
		wantErr bool
	}{
		{
			name: "recognize single character",
			chunks: [][]string{
				{
					" _    ",
					"| |   ",
					"| |__ ",
					"| '_ \\",
					"| | | ",
					"|_| |_",
					"      ",
					"      ",
				},
			},
			want:    "h",
			wantErr: false,
		},
		{
			name: "recognize two characters",
			chunks: [][]string{
				{
					" _     _    ",
					"| |   (_)   ",
					"| |__  |    ",
					"| '_ \\ |    ",
					"| | |  |    ",
					"|_| |_|_|   ",
					"            ",
					"            ",
				},
			},
			want:    "hi",
			wantErr: false,
		},
		{
			name: "recognize multiple lines",
			chunks: [][]string{
				{
					" _    ",
					"| |   ",
					"| |__ ",
					"| '_ \\",
					"| | | ",
					"|_| |_",
					"      ",
					"      ",
				},
				{
					" _    ",
					"(_)   ",
					" |    ",
					" |    ",
					" |    ",
					"|_|   ",
					"      ",
					"      ",
				},
			},
			want:    "h\ni",
			wantErr: false,
		},
		{
			name: "invalid chunk length",
			chunks: [][]string{
				{
					"line1",
					"line2",
				},
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverse.RecogniseText(tt.chunks, templates)

			if tt.wantErr {
				if err == nil {
					t.Errorf("RecognizeText() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("RecogniseText() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("RecogniseText() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRecogniseTextWithBanner(t *testing.T) {
	// Banner content with proper structure:
	// Line 0: empty header
	// Lines 1-8: space character (8 lines)
	// Line 9: empty separator
	// Lines 10-17: exclamation character (8 lines)
	bannerContent := "\n" + // Line 0: header
		"      \n" + // Line 1: space line 1
		"      \n" + // Line 2: space line 2
		"      \n" + // Line 3: space line 3
		"      \n" + // Line 4: space line 4
		"      \n" + // Line 5: space line 5
		"      \n" + // Line 6: space line 6
		"      \n" + // Line 7: space line 7
		"      \n" + // Line 8: space line 8
		"\n" + // Line 9: separator
		" _    \n" + // Line 10: ! line 1
		"| |   \n" + // Line 11: ! line 2
		"| |   \n" + // Line 12: ! line 3
		"| |   \n" + // Line 13: ! line 4
		"|_|   \n" + // Line 14: ! line 5
		"(_)   \n" + // Line 15: ! line 6
		"      \n" + // Line 16: ! line 7
		"      \n" // Line 17: ! line 8

	tests := []struct {
		name     string
		asciiArt string
		want     string
		wantErr  bool
	}{
		{
			name:     "recognize exclamation mark",
			asciiArt: " _    \n| |   \n| |   \n| |   \n|_|   \n(_)   \n      \n      \n",
			want:     "!",
			wantErr:  false,
		},
		{
			name:     "recognize space",
			asciiArt: "      \n      \n      \n      \n      \n      \n      \n      \n",
			want:     " ",
			wantErr:  false,
		},
		{
			name:     "empty ASCII art",
			asciiArt: "",
			want:     "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverse.RecogniseTextWithBanner(tt.asciiArt, bannerContent)

			if tt.wantErr {
				if err == nil {
					t.Errorf("RecognizeTextWithBanner() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("RecogniseTextWithBanner() unexpected error = %v", err)
				return
			}

			if got != tt.want {
				t.Errorf("RecogniseTextWithBanner() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestRecogniseCharacterNormalization(t *testing.T) {
	// Test that normalization works - patterns with different trailing spaces should match
	templates := map[rune][]string{
		'a': {
			"  __  ",
			" /  \\ ",
			"|    |",
			"|____|",
			"      ",
			"      ",
			"      ",
			"      ",
		},
	}

	// Pattern without trailing spaces on some lines
	pattern := []string{
		"  __",
		" /  \\",
		"|    |",
		"|____|",
		"",
		"",
		"",
		"",
	}

	got, err := reverse.RecogniseCharacter(pattern, templates)
	if err != nil {
		t.Errorf("RecogniseCharacter() with normalization unexpected error = %v", err)
		return
	}

	if got != 'a' {
		t.Errorf("RecogniseCharacter() with normalization = %q, want 'a'", got)
	}
}
