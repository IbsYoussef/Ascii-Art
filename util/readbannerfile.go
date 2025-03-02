package util

import (
	"bufio"
	"os"
)

func ReadBannerFile(filename string) ([]string, error) {
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
