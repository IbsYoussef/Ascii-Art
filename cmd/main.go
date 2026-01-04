package main

import (
	"ascii-art/internal/ascii"
	"fmt"
)

func main() {
	// Get user input and banner choice
	input, banner, err := ascii.GetUserInput()

	// If there was an error getting user input, print the error and exit
	if err != nil {
		fmt.Println(err)
		return
	}

	// Load the ASCII banner from the specified file
	result, err := ascii.LoadBannerFile(banner)

	// If the banner file is not found, print an error message and exit
	if err != nil {
		fmt.Println(err)
		return
	}

	// Render the input in banner style as ASCII art
	ascii.RenderAscii(input, result)
}
