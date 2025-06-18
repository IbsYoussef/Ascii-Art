package ascii

import (
	"fmt"
	"strings"
)

func RenderAscii(input string, banner map[rune][]string) {
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		// Build and print each of the 8 ASCII lines for this line of input
		for row := 0; row < 8; row++ {
			var outputLine string

			for _, ch := range line {
				if asciiChar, ok := banner[ch]; ok {
					outputLine += asciiChar[row]
				} else {
					outputLine += "        " // 8 spaces as placeholder
				}
			}

			fmt.Println(outputLine)
		}
	}
}
