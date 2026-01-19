package unit

import (
	reverse "ascii-art/internal/ascii-reverse"
	"testing"
)

func TestParseAsciiArt(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		wantChunks int
		checkChunk int
		wantLines  []string
	}{
		{
			name:       "single line of ASCII art (8 lines)",
			content:    " _              _   _          \n| |            | | | |         \n| |__     ___  | | | |   ___   \n|  _ \\   / _ \\ | | | |  / _ \\  \n| | | | |  __/ | | | | | (_) | \n|_| |_|  \\___| |_| |_|  \\___/  \n                               \n                               \n",
			wantChunks: 1,
			checkChunk: 0,
			wantLines: []string{
				" _              _   _          ",
				"| |            | | | |         ",
				"| |__     ___  | | | |   ___   ",
				"|  _ \\   / _ \\ | | | |  / _ \\  ",
				"| | | | |  __/ | | | | | (_) | ",
				"|_| |_|  \\___| |_| |_|  \\___/  ",
				"                               ",
				"                               ",
			},
		},
		{
			name:       "two lines of ASCII art (16 lines)",
			content:    " _              _   _          \n| |            | | | |         \n| |__     ___  | | | |   ___   \n|  _ \\   / _ \\ | | | |  / _ \\  \n| | | | |  __/ | | | | | (_) | \n|_| |_|  \\___| |_| |_|  \\___/  \n                               \n                               \n _   _   _                 _       _  \n| | | | | |               | |     | | \n| | | | | |   ___    ___  | |   __| | \n| | | | | |  / _ \\  / __| | |  / _` | \n| | | | | | | (_) | |     | | | (_| | \n|_| |_| |_|  \\___/  |     |_|  \\__,_| \n                                      \n                                      \n",
			wantChunks: 2,
			checkChunk: 1,
			wantLines: []string{
				" _   _   _                 _       _  ",
				"| | | | | |               | |     | | ",
				"| | | | | |   ___    ___  | |   __| | ",
				"| | | | | |  / _ \\  / __| | |  / _` | ",
				"| | | | | | | (_) | |     | | | (_| | ",
				"|_| |_| |_|  \\___/  |     |_|  \\__,_| ",
				"                                      ",
				"                                      ",
			},
		},
		{
			name:       "empty content",
			content:    "",
			wantChunks: 0,
		},
		{
			name:       "single character (8 lines)",
			content:    " _  \n| | \n| | \n| | \n|_| \n(_) \n    \n    \n",
			wantChunks: 1,
			checkChunk: 0,
			wantLines: []string{
				" _  ",
				"| | ",
				"| | ",
				"| | ",
				"|_| ",
				"(_) ",
				"    ",
				"    ",
			},
		},
		{
			name:       "incomplete chunk (less than 8 lines)",
			content:    " _  \n| | \n| | \n",
			wantChunks: 1,
			checkChunk: 0,
			wantLines: []string{
				" _  ",
				"| | ",
				"| | ",
				"",
				"",
				"",
				"",
				"",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := reverse.ParseAsciiArt(tt.content)
			if err != nil {
				t.Errorf("ParseAsciiArt() error = %v", err)
				return
			}

			if len(got) != tt.wantChunks {
				t.Errorf("ParseAsciiArt() returned %d chunks, want %d", len(got), tt.wantChunks)
				return
			}

			// Check specific chunk if specified
			if tt.checkChunk >= 0 && tt.checkChunk < len(got) {
				chunk := got[tt.checkChunk]

				if len(chunk) != 8 {
					t.Errorf("ParseAsciiArt() chunk %d has %d lines, want 8", tt.checkChunk, len(chunk))
					return
				}

				if tt.wantLines != nil {
					for i, line := range chunk {
						if i < len(tt.wantLines) && line != tt.wantLines[i] {
							t.Errorf("ParseAsciiArt() chunk %d line %d = %q, want %q", tt.checkChunk, i, line, tt.wantLines[i])
						}
					}
				}
			}
		})
	}
}

func TestSplitChunkIntoCharacters(t *testing.T) {
	tests := []struct {
		name      string
		chunk     []string
		charWidth int
		wantCount int
		checkChar int
		wantChar  []string
	}{
		{
			name: "split simple chunk with 2 characters",
			chunk: []string{
				" _   _  ",
				"| | | | ",
				"| | | | ",
				"| | | | ",
				"|_| |_| ",
				"(_) (_) ",
				"        ",
				"        ",
			},
			charWidth: 4,
			wantCount: 2,
			checkChar: 0,
			wantChar: []string{
				" _  ",
				"| | ",
				"| | ",
				"| | ",
				"|_| ",
				"(_) ",
				"    ",
				"    ",
			},
		},
		{
			name: "split chunk with 3 characters",
			chunk: []string{
				" _   _   _  ",
				"| | | | | | ",
				"| | | | | | ",
				"| | | | | | ",
				"|_| |_| |_| ",
				"(_) (_) (_) ",
				"            ",
				"            ",
			},
			charWidth: 4,
			wantCount: 3,
			checkChar: 1,
			wantChar: []string{
				" _  ",
				"| | ",
				"| | ",
				"| | ",
				"|_| ",
				"(_) ",
				"    ",
				"    ",
			},
		},
		{
			name: "empty chunk",
			chunk: []string{
				"",
				"",
				"",
				"",
				"",
				"",
				"",
				"",
			},
			charWidth: 4,
			wantCount: 0,
		},
		{
			name: "chunk with different character widths",
			chunk: []string{
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
			},
			charWidth: 6,
			wantCount: 1,
			checkChar: 0,
			wantChar: []string{
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
				"      ",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse.SplitChunkIntoCharacters(tt.chunk, tt.charWidth)

			if len(got) != tt.wantCount {
				t.Errorf("SplitChunkIntoCharacters() returned %d characters, want %d", len(got), tt.wantCount)
				return
			}

			// Check specific character if specified
			if tt.checkChar >= 0 && tt.checkChar < len(got) && tt.wantChar != nil {
				char := got[tt.checkChar]

				if len(char) != 8 {
					t.Errorf("SplitChunkIntoCharacters() character %d has %d lines, want 8", tt.checkChar, len(char))
					return
				}

				for i, line := range char {
					if i < len(tt.wantChar) && line != tt.wantChar[i] {
						t.Errorf("SplitChunkIntoCharacters() char %d line %d = %q, want %q", tt.checkChar, i, line, tt.wantChar[i])
					}
				}
			}
		})
	}
}

func TestSplitChunkIntoCharactersInvalidInput(t *testing.T) {
	// Test with wrong number of lines
	chunk := []string{"line1", "line2"}
	got := reverse.SplitChunkIntoCharacters(chunk, 4)

	if got != nil {
		t.Errorf("SplitChunkIntoCharacters() with invalid chunk = %v, want nil", got)
	}
}

func TestNormalizePattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern []string
		want    []string
	}{
		{
			name: "normalize uneven pattern",
			pattern: []string{
				" _  ",
				"| | ",
				"|_|",
			},
			want: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
		},
		{
			name: "already normalized pattern",
			pattern: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			want: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
		},
		{
			name:    "empty pattern",
			pattern: []string{},
			want:    []string{},
		},
		{
			name: "pattern with varying lengths",
			pattern: []string{
				"a",
				"abc",
				"ab",
				"abcd",
			},
			want: []string{
				"a   ",
				"abc ",
				"ab  ",
				"abcd",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse.NormalizePattern(tt.pattern)

			if len(got) != len(tt.want) {
				t.Errorf("NormalizePattern() length = %d, want %d", len(got), len(tt.want))
				return
			}

			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("NormalizePattern() line %d = %q, want %q", i, got[i], tt.want[i])
				}
			}
		})
	}
}
