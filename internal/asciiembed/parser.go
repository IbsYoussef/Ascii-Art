package asciiembed

import (
	"bufio"
	"io"
)

func parseBannerFromReader(reader io.Reader) (map[rune][]string, error) {
	scanner := bufio.NewScanner(reader)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	bannerMap := make(map[rune][]string)
	currentRune := rune(32)
	i := 0

	for i < len(lines) && lines[i] == "" {
		i++
	}

	for i+8 <= len(lines) && currentRune <= 126 {
		bannerMap[currentRune] = lines[i : i+8]
		i += 9
		currentRune++
	}

	return bannerMap, nil

}
