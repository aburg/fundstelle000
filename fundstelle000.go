package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"

	"github.com/fatih/color"
)

func main() {
	green := color.New(color.FgGreen).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()

	pattern := regexp.MustCompile(`\b([A-Z]{1,3})(\d{1,2})\b`)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == ".git" {
			yellow("skipping git dir\n")
			return filepath.SkipDir
		}

		if pattern.MatchString(info.Name()) {
			red("%s", info.Name())
			fmt.Print(" -> ")
			green("%s\n", info.Name())
		}

		// fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking: %v\n", err)
		return
	}
}
