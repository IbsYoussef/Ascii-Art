package main

import (
	"fmt"
	"os"
)

func main() {
	// Gather input from user
	input := os.Args[1:]

	// Check if arguements were provided
	if len(input) < 1 {
		fmt.Println("Usage go run . \"<text-input>\" [banner-name]")
		fmt.Println("Note: Text input should be enclosed in quotes if it contains spaces")
		os.Exit(1)
	}

	// First arguement should be the text input (in quotes when running from command line)
	textInput := input[0]

	// Default banner
	banner := "standard"

	// If there's a second argument, it's the banner name
	if len(input) >= 2 {
		banner = input[1]
	}

	// Check if text input is empty
	if textInput == "" {
		fmt.Println("Error: Text input cannot be empty")
		os.Exit(1)
	}

	// Output for testing
	fmt.Printf("Text to convert: '%s'\n", textInput)
	fmt.Printf("Using Banner: '%s'\n", banner)

}
