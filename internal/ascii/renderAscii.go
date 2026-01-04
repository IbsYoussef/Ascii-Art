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

// This version of RenderAscii does handles terminal wrapping more efficiently.
// func RenderAsciiWithTerminalWrapping(input string, banner map[rune][]string) {
// 	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
// 	if err != nil {
// 		termWidth = 80 // Fallback to a default width if detection fails
// 	}

// 	// Handle the case where the input string is just newlines
// 	if len(strings.ReplaceAll(input, "\n", "")) == 0 {
// 		fmt.Print(strings.Repeat("\n", len(strings.Split(input, "\n"))-1))
// 	}

// 	// Split the user's input by the "\n" delimiter
// 	inputFields := strings.Split(input, "\n")

// 	for _, line := range inputFields {
// 		if line == "" {
// 			fmt.Println()
// 			continue
// 		}

// 		// outputRows will store the 8 lines of the final, wrapped output
// 		var outputRows [8]string

// 		for _, ch := range line {
// 			// Look up the character art from the map
// 			charArt, ok := banner[ch]

// 			// If character is not in the banner map, we can't process it.
// 			// We'll skip it to avoid breaking the alignment.
// 			if !ok {
// 				// You could also print a placeholder or log a warning here.
// 				continue
// 			}

// 			// Get the width of this specific character's art
// 			charWidth := 0
// 			if len(charArt) > 0 {
// 				charWidth = len(charArt[0])
// 			}

// 			// DYNAMIC WRAPPING LOGIC:
// 			// Before adding the new character, check if it will exceed the terminal width.
// 			if len(outputRows[0])+charWidth > termWidth {
// 				// If it exceeds, print the lines we've built so far...
// 				for _, row := range outputRows {
// 					fmt.Println(row)
// 				}
// 				// ...and reset for the next line of wrapped output.
// 				outputRows = [8]string{}
// 			}

// 			// Append the new character's art to our output rows
// 			for i := 0; i < 8; i++ {
// 				if i < len(charArt) {
// 					outputRows[i] += charArt[i]
// 				}
// 			}
// 		}

// 		// After processing a line, print any remaining content in outputRows
// 		for _, row := range outputRows {
// 			fmt.Println(row)
// 		}
// 	}
// }
