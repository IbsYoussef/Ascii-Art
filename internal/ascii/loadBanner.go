package ascii

import (
	"ascii-art/internal/files"
)

func LoadBannerFile(stylename string) (map[rune][]string, error) {
	// Set the banner file path
	path := "banners/" + stylename + ".txt"

	// Read all lines from the file
	lines, err := files.ReadBannerLines(path)
	if err != nil {
		return nil, err
	}

	// Create a map of runes to lines
	bannerMap := make(map[rune][]string)
	currentRune := rune(32) // Start at ASCII 32(' ')
	i := 0

	// Skip leading empty lines (padding at top)
	for i < len(lines) && lines[i] == "" {
		i++
	}

	// Loop through file in blocks of 9 (8 content + 1 seperator)
	for i+8 <= len(lines) && currentRune <= 126 {
		charLines := lines[i : i+8] // Get 8 block lines
		bannerMap[currentRune] = charLines // Store those lines under currentRune count
		i += 9 // Skip the seperator line // Skip the seperator after each character
		currentRune++ // Move to next character rune count
	}

	return bannerMap, nil

}
