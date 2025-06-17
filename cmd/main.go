package main

import (
	"ascii-art/internal/ascii"
	"fmt"
)

func main() {
	input, banner, err := ascii.GetUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User Input: ", input)
	fmt.Println("Banner Choice: ", banner)
	fmt.Println()

	result, err := ascii.LoadBannerFile(banner)
	if err != nil {
		fmt.Println(err)
	}

	for r := rune(32); r <= 126; r++ {
		fmt.Printf("Character: %q (ASCII %d)\n", r, r)
		if lines, ok := result[r]; ok {
			for _, line := range lines {
				fmt.Println(line)
			}
		} else {
			fmt.Println("Missing or corrupted")
		}
		fmt.Println()
	}
}
