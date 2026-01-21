package asciijustify

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// HandleJustify orchestrates the justify alignment feature
// For justify, it renders words separately. For other alignments, it captures output.
func HandleJustify(renderFunc func(), alignType string, input string, banner map[rune][]string) error {
	// Special handling for justify - render words separately
	if alignType == "justify" {
		termWidth := GetTerminalWidth()
		justifiedLines := RenderWithJustify(input, banner, termWidth)
		for _, line := range justifiedLines {
			fmt.Println(line)
		}
		return nil
	}

	// For left alignment, just render normally
	if alignType == "left" {
		renderFunc()
		return nil
	}

	// For center and right, capture output and apply alignment
	// Save original stdout
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		return err
	}
	os.Stdout = w

	// Channel to capture output
	outputChan := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outputChan <- buf.String()
	}()

	// Execute render function
	renderFunc()

	// Close writer and restore stdout
	w.Close()
	os.Stdout = oldStdout

	// Get captured output
	output := <-outputChan

	// Split output into lines
	lines := strings.Split(output, "\n")

	// Remove trailing empty line if exists
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	if len(lines) == 0 {
		return nil
	}

	// Get terminal width
	termWidth := GetTerminalWidth()

	// Apply alignment
	alignedLines := ApplyAlignment(lines, alignType, termWidth)

	// Print aligned output
	for _, line := range alignedLines {
		fmt.Println(line)
	}

	return nil
}

// RenderWithJustify renders text with justify alignment by rendering words separately
// This is necessary because justify needs to control spacing between words
func RenderWithJustify(input string, banner map[rune][]string, termWidth int) []string {
	result := make([]string, 0)

	// Handle empty input or only newlines
	if len(strings.ReplaceAll(input, "\n", "")) == 0 {
		lineCount := len(strings.Split(input, "\n")) - 1
		for i := 0; i < lineCount; i++ {
			result = append(result, "")
		}
		return result
	}

	// Split input by newlines
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if line == "" {
			result = append(result, "")
			continue
		}

		// Split line into words
		words := strings.Fields(line) // Fields splits by whitespace

		if len(words) == 0 {
			result = append(result, "")
			continue
		}

		// Render each word separately
		renderedWords := make([][]string, len(words))
		for i, word := range words {
			renderedWords[i] = renderWord(word, banner)
		}

		// Apply justify spacing between rendered words
		justifiedLines := justifyRenderedWords(renderedWords, termWidth)
		result = append(result, justifiedLines...)
	}

	return result
}

// renderWord renders a single word as ASCII art
func renderWord(word string, banner map[rune][]string) []string {
	result := make([]string, 8)

	for _, ch := range word {
		charArt, ok := banner[ch]
		if !ok {
			continue
		}

		// Append each line of the character to the result
		for i := 0; i < 8; i++ {
			if i < len(charArt) {
				result[i] += charArt[i]
			}
		}
	}

	return result
}

// justifyRenderedWords applies justify spacing between already-rendered words
func justifyRenderedWords(renderedWords [][]string, termWidth int) []string {
	if len(renderedWords) == 0 {
		return []string{}
	}

	// If only one word, center it
	if len(renderedWords) == 1 {
		return CenterAlign(renderedWords[0], termWidth)
	}

	// Calculate total width of all words
	totalWordWidth := 0
	for _, word := range renderedWords {
		if len(word) > 0 {
			totalWordWidth += len(word[0])
		}
	}

	// Tuning variable
	const marginPercent = 12

	// Apply margin
	margin := (termWidth * marginPercent) / 100
	usableWidth := termWidth - (2 * margin)

	// Check if content fits
	if totalWordWidth >= usableWidth {
		// Too wide, center instead
		combined := combineWords(renderedWords, 0)
		return CenterAlign(combined, termWidth)
	}

	// Calculate spacing
	availableSpace := usableWidth - totalWordWidth
	gaps := len(renderedWords) - 1

	if gaps <= 0 || availableSpace <= 0 {
		combined := combineWords(renderedWords, 0)
		return CenterAlign(combined, termWidth)
	}

	spacePerGap := availableSpace / gaps

	extraSpace := availableSpace % gaps

	// Build justified output
	result := make([]string, 8)
	for row := 0; row < 8; row++ {
		// Start with left margin
		line := strings.Repeat(" ", margin)

		for wordIdx, word := range renderedWords {
			// Add the word
			if row < len(word) {
				line += word[row]
			}

			// Add spacing after word (except last word)
			if wordIdx < len(renderedWords)-1 {
				spacing := spacePerGap
				if wordIdx < extraSpace {
					spacing++
				}
				line += strings.Repeat(" ", spacing)
			}
		}
		result[row] = line
	}

	return result
}

// combineWords combines rendered words with fixed spacing between them
func combineWords(renderedWords [][]string, spacing int) []string {
	if len(renderedWords) == 0 {
		return []string{}
	}

	result := make([]string, 8)
	spacer := strings.Repeat(" ", spacing)

	for row := 0; row < 8; row++ {
		for wordIdx, word := range renderedWords {
			if row < len(word) {
				result[row] += word[row]
			}
			// Add spacing between words (except after last word)
			if wordIdx < len(renderedWords)-1 {
				result[row] += spacer
			}
		}
	}

	return result
}
