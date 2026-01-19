package asciireverse

import (
	"strings"
)

// CharacterTemplate represents an ASCII art character pattern
type CharacterTemplate struct {
	Char    rune     // The character this template represents
	Pattern []string // 8 lines representing the ASCII art pattern
}

// LoadBannerTemplates loads the banner file and creates a map of character patterns
// Returns a map where key is the character abd value is its 8-line ASCII pattern
func LoadBannerTemplates(bannerContent string) (map[rune][]string, error) {
	templates := make(map[rune][]string)

	// Split banner content into lines
	lines := strings.Split(bannerContent, "\n")

	// ASCII printable characters range from 32 (space) to 126 (~)
	// Each character has 9 lines (1 header + 8 pattern lines)
	for char := ' '; char <= '~'; char++ {
		// Calculate starting line for this character
		// First line (char 32, space) starts at line 1 (line 0 is usually empty/header)
		startLine := int(char-32)*9 + 1

		// Check if we have enough lines
		if startLine+8 > len(lines) {
			break
		}

		// Extract 8 lines for this character
		pattern := make([]string, 8)
		for i := 0; i < 8; i++ {
			if startLine+i < len(lines) {
				pattern[i] = lines[startLine+i]
			} else {
				pattern[i] = ""
			}
		}

		templates[char] = pattern
	}

	return templates, nil
}

func GetCharacterWidth(pattern []string) int {
	for _, line := range pattern {
		if len(line) > 0 {
			return len(line)
		}
	}
	return 0
}

// ComparePatterns compares two patterns and returns true if they match
func ComparePatterns(pattern1, pattern2 []string) bool {
	if len(pattern1) != len(pattern2) {
		return false
	}

	for i := range pattern1 {
		if pattern1[i] != pattern2[i] {
			return false
		}
	}

	return true
}
