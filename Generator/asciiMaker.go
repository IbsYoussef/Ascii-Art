package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Get the command-line arguments excluding the program name
	Arg := os.Args[1:]

	// Check if there are exactly two command-line arguments
	if len(Arg) != 2 {
		// If not, return from the main function and terminate the program
		return
	}

	// Iterate over each character in the first command-line argument
	for _, r := range Arg[0] {
		// Check if the character is outside the printable ASCII range (from ' ' to '~')
		if r < ' ' || r > '~' {
			// If a non-printable character is found, return and terminate the program
			return
		}
	}

	// Read the contents of the "standard.txt" file located in the "banners" folder
	bytes, err := os.ReadFile("./banners/standard.txt")
	if err != nil {
		// If there was an error reading the file, exit the program with an error code (1)
		os.Exit(1)
	}

	// Split the contents of the file into lines and store them in the "lines" slice
	lines := strings.Split(string(bytes), "\n")

	// Initialize an empty slice of runes (characters) named "arr"
	var arr []rune

	// Initialize a boolean variable "Newline" to false
	Newline := false

	// Iterate over each character in the first command-line argument again
	for i, r := range Arg[0] {
		// Check if the "Newline" flag is true (indicating the need for a new line of art)
		if Newline {
			// Reset the "Newline" flag to false
			Newline = false

			// Call the "printArt" function to print the ASCII art stored in "arr"
			// The "true" argument indicates that a new line should be printed before the art
			printArt(arr, lines, Arg[1], true)

			// Reset the "arr" slice to an empty slice to store the next line of art characters
			arr = []rune{}

			// Continue to the next character in the loop
			continue
		}

		// Check if the current character is a backslash ('\') and not the last character
		if r == '\\' && len(Arg[0]) != i+1 {
			// Check if the next character after the backslash is 'n' (newline)
			if Arg[0][i+1] == 'n' {
				// If it is, set the "Newline" flag to true, indicating the need for a new line of art
				Newline = true

				// Continue to the next character in the loop
				continue
			}
		}

		// If none of the special conditions are met, add the current character to the "arr" slice
		arr = append(arr, r)
	}

	// After the loop, check if the "Newline" flag is still false
	if !Newline {
		// If "Newline" is false, call the "printArt" function to print the remaining art in "arr"
		// The "false" argument indicates that no new line should be printed before the art
		printArt(arr, lines, Arg[1], false)
	}
}

// The "printArt" function takes the ASCII art characters in "arr", the "lines" slice containing the ASCII art templates,
// the "outputFile" string representing the output file path, and the "Newline" boolean indicating whether to print a new line before the art.
func printArt(arr []rune, lines []string, outputFile string, Newline bool) {
	// Check if the "arr" slice is not empty (i.e., contains ASCII art characters)
	if len(arr) != 0 {
		// Open the output file in append mode and create it if it doesn't exist
		f, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			// If there was an error opening the output file, print an error message and exit the program
			fmt.Println("Error creating output file:")
			os.Exit(0)
		}
		defer f.Close()

		// Iterate over each line of the ASCII art (8 lines)
		for line := 1; line <= 8; line++ {
			var artLine string

			// Iterate over each ASCII art character in "arr"
			for _, r := range arr {
				// Calculate the appropriate offset to retrieve the corresponding ASCII art line from the "lines" slice
				skip := (r - 32) * 9
				// Concatenate the ASCII art line to "artLine"
				artLine += lines[line+int(skip)]
			}

			// Append a newline character at the end of the ASCII art line
			artLine += "\n"

			// Write the "artLine" to the output file
			_, err = f.WriteString(artLine)
			if err != nil {
				// If there was an error writing to the output file, print an error message and exit the program
				fmt.Println("Error writing output file:")
				os.Exit(0)
			}
		}

		// If the "Newline" flag is true, write an additional newline character to the output file
		if Newline {
			_, err = f.WriteString("\n")
			if err != nil {
				fmt.Println("Error writing output file:")
				os.Exit(0)
			}
		}
	} else {
		// If the "arr" slice is empty (no ASCII art characters), print an empty line to the console
		fmt.Println()
	}
}
