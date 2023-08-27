package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	str := "Hello"
	split := strings.Split(str, "\n")
	filename := "standard"

	bannerLines, err := readBannerFile("Banners/" + filename + ".txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(GenerateASCIIArt(split, bannerLines))
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

func GenerateASCIIArt(words []string, bannerLines []string) string {
	var asciiArt string

	for i, word := range words {
		if word == "" {
			if i < len(words)-1 {
				asciiArt += "\n"
			}
			continue
		}
		asciiArt += generateWordBanner(word, bannerLines)
	}

	return asciiArt
}

func generateWordBanner(word string, bannerLines []string) string {
	var wordBanner string

	for row := 1; row < 9; row++ {
		for _, char := range word {
			lineIndex := (int(char)-32)*9 + row
			if lineIndex >= 0 && lineIndex < len(bannerLines) {
				wordBanner += bannerLines[lineIndex]
			}
		}
		wordBanner += "\n"
	}

	return wordBanner
}
