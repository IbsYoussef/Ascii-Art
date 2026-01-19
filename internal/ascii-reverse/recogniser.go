package asciireverse

import (
	"fmt"
	"strings"
)

// RecogniseText takes ASCII art chunks and banner templates, and returns the recognized text
// Handles variable-width characters by trying to match each character from the templates
func RecogniseText(chunks [][]string, templates map[rune][]string) (string, error) {
	var result strings.Builder

	// Process each chunk (each chunk is one line of text)
	for chunkIdx, chunk := range chunks {
		if len(chunk) != 8 {
			return "", fmt.Errorf("invalid chunk at index %d: expected 8 lines, got %d", chunkIdx, len(chunk))
		}

		// Find the maximum line length in the chunk
		maxLen := 0
		for _, line := range chunk {
			if len(line) > maxLen {
				maxLen = len(line)
			}
		}

		// If chunk is empty, skip it
		if maxLen == 0 {
			continue
		}

		// Process characters by trying to match at each position
		pos := 0
		for pos < maxLen {
			// Try to match each character in templates
			matched := false
			var matchedChar rune
			matchLen := 0

			// Try all characters from templates
			for char, template := range templates {
				charWidth := GetCharacterWidth(template)
				if charWidth == 0 {
					continue
				}

				// Extract pattern at current position with this width
				if pos+charWidth > maxLen {
					// This character would extend beyond the line, try smaller widths
					continue
				}

				pattern := make([]string, 8)
				for i := 0; i < 8; i++ {
					if pos < len(chunk[i]) {
						endPos := pos + charWidth
						if endPos <= len(chunk[i]) {
							pattern[i] = chunk[i][pos:endPos]
						} else {
							// Pad with spaces
							pattern[i] = chunk[i][pos:]
							for len(pattern[i]) < charWidth {
								pattern[i] += " "
							}
						}
					} else {
						pattern[i] = strings.Repeat(" ", charWidth)
					}
				}

				// Try to match this pattern
				normalizedPattern := NormalizePattern(pattern)
				normalizedTemplate := NormalizePattern(template)

				if ComparePatterns(normalizedPattern, normalizedTemplate) {
					// Found a match!
					matched = true
					matchedChar = char
					matchLen = charWidth
					break
				}
			}

			if matched {
				result.WriteRune(matchedChar)
				pos += matchLen
			} else {
				// No match found - this is an error
				return "", fmt.Errorf("failed to recognise character at position %d", pos)
			}
		}

		// Add newline between chunks (except for last chunk)
		if chunkIdx < len(chunks)-1 {
			result.WriteRune('\n')
		}
	}

	return result.String(), nil
}

// RecogniseCharacter matches a pattern against templates and returns the recognized character
func RecogniseCharacter(pattern []string, templates map[rune][]string) (rune, error) {
	// Normalize the pattern to ensure consistent comparison
	normalizedPattern := NormalizePattern(pattern)

	// Try to match against all templates
	for char, template := range templates {
		normalizedTemplate := NormalizePattern(template)

		if ComparePatterns(normalizedPattern, normalizedTemplate) {
			return char, nil
		}
	}

	// If no match found, return error
	return 0, fmt.Errorf("no matching character found for pattern")
}

// RecogniseTextWithBanner is a convenience function that loads banner content and recognizes text
func RecogniseTextWithBanner(asciiArt string, bannerContent string) (string, error) {
	// Parse ASCII art into chunks
	chunks, err := ParseAsciiArt(asciiArt)
	if err != nil {
		return "", fmt.Errorf("failed to parse ASCII art: %w", err)
	}

	// Load banner templates
	templates, err := LoadBannerTemplates(bannerContent)
	if err != nil {
		return "", fmt.Errorf("failed to load banner templates: %w", err)
	}

	// Recognise text
	text, err := RecogniseText(chunks, templates)
	if err != nil {
		return "", fmt.Errorf("failed to recognise text: %w", err)
	}

	return text, nil
}
