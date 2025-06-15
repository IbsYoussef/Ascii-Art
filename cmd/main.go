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

	fmt.Println("Input:", input)
	fmt.Println("Banner:", banner)
}
