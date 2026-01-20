package unit

import (
	reverse "ascii-art/internal/ascii-reverse"
	"testing"
)

func TestLoadBannerTemplates(t *testing.T) {
	tests := []struct {
		name          string
		bannerContent string
		wantChars     []rune
		checkPattern  rune
		wantPattern   []string
	}{
		{
			name:          "load simple banner with space character",
			bannerContent: "\n      \n      \n      \n      \n      \n      \n      \n      \n",
			wantChars:     []rune{' '},
			checkPattern:  ' ',
			wantPattern: []string{
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
		{
			name:          "load banner with multiple characters",
			bannerContent: "\n      \n      \n      \n      \n      \n      \n      \n      \n _  \n| | \n| | \n| | \n|_| \n(_) \n    \n    \n",
			wantChars:     []rune{' ', '!'},
			checkPattern:  '!',
			wantPattern: []string{
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
			name:          "load banner with space and exclamation",
			bannerContent: "\n      \n      \n      \n      \n      \n      \n      \n      \n _  \n| | \n| | \n| | \n|_| \n(_) \n    \n    \n",
			wantChars:     []rune{' ', '!'},
			checkPattern:  ' ',
			wantPattern: []string{
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
			got, err := reverse.LoadBannerTemplates(tt.bannerContent)
			if err != nil {
				t.Errorf("LoadBannerTemplates() error = %v", err)
				return
			}

			// Check that expected characters are present
			for _, char := range tt.wantChars {
				if _, exists := got[char]; !exists {
					t.Errorf("LoadBannerTemplates() missing character %q", char)
				}
			}

			// Check specific pattern if specified
			if tt.checkPattern != 0 {
				pattern, exists := got[tt.checkPattern]
				if !exists {
					t.Errorf("LoadBannerTemplates() missing pattern for character %q", tt.checkPattern)
					return
				}

				if len(pattern) != len(tt.wantPattern) {
					t.Errorf("LoadBannerTemplates() pattern length = %d, want %d", len(pattern), len(tt.wantPattern))
					return
				}

				for i := range pattern {
					if pattern[i] != tt.wantPattern[i] {
						t.Errorf("LoadBannerTemplates() pattern line %d = %q, want %q", i, pattern[i], tt.wantPattern[i])
					}
				}
			}
		})
	}
}

func TestGetCharacterWidth(t *testing.T) {
	tests := []struct {
		name    string
		pattern []string
		want    int
	}{
		{
			name: "simple pattern width",
			pattern: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			want: 4,
		},
		{
			name: "wider pattern",
			pattern: []string{
				" _    _ ",
				"| |  | |",
				"|_|  |_|",
			},
			want: 9,
		},
		{
			name: "empty pattern",
			pattern: []string{
				"",
				"",
				"",
			},
			want: 0,
		},
		{
			name: "pattern with leading empty lines",
			pattern: []string{
				"",
				"",
				"  __  ",
				" |  | ",
			},
			want: 6,
		},
		{
			name: "space character pattern (all spaces)",
			pattern: []string{
				"      ",
				"      ",
				"      ",
				"      ",
			},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse.GetCharacterWidth(tt.pattern)
			if got != tt.want {
				t.Errorf("GetCharacterWidth() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestComparePatterns(t *testing.T) {
	tests := []struct {
		name     string
		pattern1 []string
		pattern2 []string
		want     bool
	}{
		{
			name: "identical patterns",
			pattern1: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			pattern2: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			want: true,
		},
		{
			name: "different patterns",
			pattern1: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			pattern2: []string{
				" _  ",
				"| | ",
				"| | ",
			},
			want: false,
		},
		{
			name: "different lengths",
			pattern1: []string{
				" _  ",
				"| | ",
			},
			pattern2: []string{
				" _  ",
				"| | ",
				"|_| ",
			},
			want: false,
		},
		{
			name:     "both empty",
			pattern1: []string{},
			pattern2: []string{},
			want:     true,
		},
		{
			name: "whitespace difference",
			pattern1: []string{
				" _  ",
				"| | ",
			},
			pattern2: []string{
				" _ ",
				"| |",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse.ComparePatterns(tt.pattern1, tt.pattern2)
			if got != tt.want {
				t.Errorf("ComparePatterns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadBannerTemplatesCharacterRange(t *testing.T) {
	// Create a simple banner with space character (char 32)
	bannerContent := "\n      \n      \n      \n      \n      \n      \n      \n      \n"

	templates, err := reverse.LoadBannerTemplates(bannerContent)
	if err != nil {
		t.Fatalf("LoadBannerTemplates() error = %v", err)
	}

	// Should have at least the space character
	if len(templates) == 0 {
		t.Error("LoadBannerTemplates() returned empty map")
	}

	// Check that space character exists
	if _, exists := templates[' ']; !exists {
		t.Error("LoadBannerTemplates() missing space character")
	}

	// Check that each pattern has 8 lines
	for char, pattern := range templates {
		if len(pattern) != 8 {
			t.Errorf("LoadBannerTemplates() character %q has %d lines, want 8", char, len(pattern))
		}
	}
}
