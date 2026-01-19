package asciireverse

import (
	"fmt"
	"strings"
)

// RecogniseText takes ASCII art chunks and banner templates, and returns the recognized text
func RecogniseText(chunks [][]string, templates map[rune][]string) (string, error) {
	var result strings.Builder

	// Process each chunk (each chunk is one line of text)
	for chunkIdx, chunk := range chunks {
		if len(chunk) != 8 {
			return "", fmt.Errorf("invalid chunk at index %d: expected 8 lines, got %d", chunkIdx, len(chunk))
		}

		// Determine character width from templates (using space character as reference)
		charWidth := GetCharacterWidth(templates[' '])
		if charWidth == 0 {
			return "", fmt.Errorf("invalid character width: cannot be 0")
		}

		// Split chunk into individual character patterns
		characterPatterns := SplitChunkIntoCharacters(chunk, charWidth)

		// Recognize each character pattern
		for _, pattern := range characterPatterns {
			char, err := RecogniseCharacter(pattern, templates)
			if err != nil {
				return "", fmt.Errorf("failed to recognize character: %w", err)
			}
			result.WriteRune(char)
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
