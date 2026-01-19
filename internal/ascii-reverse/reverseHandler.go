package asciireverse

import (
	"fmt"
	"os"
)

// HandleReverse processes the --reverse flag and converts ASCII art back to text
// This is the main entry point for the reverse feature
func HandleReverse(args []string) {
	// Parse the reverse flag to get filename
	filename, _, err := ParseReverseFlag(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Read the ASCII art file
	asciiArtContent, err := ReadAsciiArtFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Try all available banners to find a match
	banners := []string{"standard", "shadow", "thinkertoy"}

	for _, bannerName := range banners {
		bannerPath := fmt.Sprintf("banners/%s.txt", bannerName)

		// Load banner template
		bannerBytes, err := os.ReadFile(bannerPath)
		if err != nil {
			// Skip this banner if it can't be read
			continue
		}
		bannerContent := string(bannerBytes)

		// Try to convert ASCII art back to text with this banner
		text, err := RecogniseTextWithBanner(asciiArtContent, bannerContent)
		if err == nil {
			// Successfully recognized! Print and exit
			fmt.Print(text)
			return
		}
	}

	// If we get here, no banner matched
	fmt.Println("Error: unable to recognize ASCII art with any available banner")
}
