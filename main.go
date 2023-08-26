package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	str := strings.Join(os.Args[1:], " ")
	words := strings.Split(str, `\n`)
	filename := "standard"
	if len(os.Args[1:]) == 2 {
		filename = os.Args[2]
	}

	bannerLines, err := readBannerFile("Banners/" + filename + ".txt")
	if err != nil {
		panic(err)
	}

	GenerateASCIIArt(words, bannerLines)
}

func readBannerFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var bannerLines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bannerLines = append(bannerLines, scanner.Text())
	}
	return bannerLines, scanner.Err()
}

func GenerateASCIIArt(words []string, bannerLines []string) {
	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				fmt.Println()
			}
			continue
		}
		printWordBanner(word, bannerLines)
	}
}

func printWordBanner(word string, bannerLines []string) {
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
