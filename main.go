package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Combine the command-line arguments into a single string.
	str := strings.Join(os.Args[1:], " ")
	// Split the string into individual words using "\n" as the separator.
	words := strings.Split(str, `\n`)
	// Open the file containing the banner text.
	var filename string
	if len(os.Args[1:]) == 2 {
		filename = os.Args[2]
	} else {
		filename = "standard"
	}
	bannerFile, err := os.Open(filename + ".txt")
	if err != nil {
		// If there was an error opening the file, panic.
		panic(err)
	}
	// Make sure to close the file when the function returns.
	defer bannerFile.Close()
	// Create a scanner to read the contents of the file line by line.
	scanner := bufio.NewScanner(bannerFile)
	// Store each line of the banner text in a slice.
	var bannerLines []string
	for scanner.Scan() {
		bannerLines = append(bannerLines, scanner.Text())
	}
	// Iterate over each word in the input.
	for i, word := range words {
		// If the word is empty, print a newline and continue to the next word.
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}
		// For each "line" in the banner text (represented by a sequence of characters
		// on a particular row), print the appropriate characters for the current word.
		for row := 1; row < 9; row++ {
			for _, char := range word {
				for lineIndex, line := range bannerLines {
					// Determine which line corresponds to the current character and row.
					if lineIndex == (int(char)-32)*9+row {
						fmt.Print(line)
					}
				}
			}
			// Print a newline after each row of the word.
			fmt.Println()
		}
	}
}
