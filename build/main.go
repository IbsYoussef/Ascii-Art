package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	syncBanners()

	targets := []struct {
		GOOS   string
		GOARCH string
		OUT    string
	}{
		{"windows", "amd64", "ascii-art-windows.exe"},
		{"linux", "amd64", "ascii-art-linux"},
		{"darwin", "amd64", "ascii-art-macos-intel"},
		{"darwin", "arm64", "ascii-art-macos-arm64"},
	}

	for _, t := range targets {
		fmt.Printf("📦 Building for %s/%s...\n", t.GOOS, t.GOARCH)

		cmd := exec.Command("go", "build", "-o", t.OUT, "./build/embedmain/main.go")
		cmd.Env = append(os.Environ(),
			"GOOS="+t.GOOS,
			"GOARCH="+t.GOARCH,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("❌ Build failed: %v\n", err)
		} else {
			fmt.Printf("✅ Built: %s\n", t.OUT)
		}
	}
}

func syncBanners() {
	src := "banners"
	dst := "internal/asciiembed/banners"

	fmt.Println("🔄 Syncing banner files...")

	if err := os.MkdirAll(dst, 0755); err != nil {
		fmt.Printf("❌ Failed to create target folder: %v\n", err)
		return
	}

	err := filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			destPath := filepath.Join(dst, filepath.Base(path))
			return copyFile(path, destPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("❌ Sync error: %v\n", err)
	} else {
		fmt.Println("✅ Banner sync complete.")
	}
}

func copyFile(src, dst string) error {
	from, err := os.Open(src)
	if err != nil {
		return err
	}
	defer from.Close()

	to, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	return err
}
