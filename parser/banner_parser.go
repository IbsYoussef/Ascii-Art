package parser

import (
	"ascii-art/util"
	"strings"
)

func LoadBanner(filePath string) (map[rune][]string, error) {
	data, err := util.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	banner := make(map[rune][]string)

	ascii := 32 // starting from ' '
	for i := 0; i+7 < len(lines); i += 8 {
		charLines := lines[i : 1+8]
		banner[rune(ascii)] = charLines
		ascii++
	}

	return banner, nil
}
