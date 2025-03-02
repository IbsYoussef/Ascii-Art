package util

import "fmt"

func PrintWordBanner(word string, bannerLines []string) {
	for row := 1; row < 9; row++ {
		for _, char := range word {
			lineIndex := (int(char)-32)*9 + row
			if lineIndex >= 0 && lineIndex < len(bannerLines) {
				fmt.Print(bannerLines[lineIndex])
			}
		}
		fmt.Println()
	}
}
