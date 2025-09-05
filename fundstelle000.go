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

	// one digit -> three digits
	pattern1 := regexp.MustCompile(`\b([A-Z]{1,3})(\d{1})\b`)
	// two digits -> three digits
	pattern2 := regexp.MustCompile(`\b([A-Z]{1,3})(\d{2})\b`)

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == ".git" {
			yellow("skipping git dir\n")
			return filepath.SkipDir
		}

		if pattern1.MatchString(info.Name()) {
			red("%s", info.Name())
			fmt.Print(" -> ")
			new := pattern1.ReplaceAllString(info.Name(), "${1}00$2")
			green("%s\n", new)
		}

		if pattern2.MatchString(info.Name()) {
			red("%s", info.Name())
			fmt.Print(" -> ")
			new := pattern2.ReplaceAllString(info.Name(), "${1}0$2")
			green("%s\n", new)
		}

		// fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking: %v\n", err)
		return
	}
}
