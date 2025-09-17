package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

type MatchPattern struct {
	Nullen int
	Regexp *regexp.Regexp
}

func rename(path string, newName string) {
	err := os.Rename(path, filepath.Join(filepath.Dir(path), newName))
	if err != nil {
		panic(err)
	}
}

func main() {
	doTheWork := false

	if len(os.Args) > 1 {
		if arg := os.Args[1]; arg == "-w" {
			doTheWork = true
		} else {
			fmt.Printf("unknown argument: %s (either give NO args or give '-w' as first and only arg)\n", arg)
			os.Exit(1)
		}
	}

	matchPatterns := []MatchPattern{
		{2, regexp.MustCompile(`\b([A-Z]{1,3})(\d{1})\b`)},
		{1, regexp.MustCompile(`\b([A-Z]{1,3})(\d{2})\b`)},
	}

	green := color.New(color.FgGreen).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()
	yellow := color.New(color.FgYellow).PrintfFunc()

	skipPatterns := []*regexp.Regexp{
		regexp.MustCompile("^gps"),
		regexp.MustCompile("^track"),
	}

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == ".git" {
			yellow("skipping git dir\n")
			return filepath.SkipDir
		}

		if !info.IsDir() {

			for _, skipPattern := range skipPatterns {
				if skipPattern.MatchString(strings.ToLower(info.Name())) {
					yellow("skipping: %s\n", info.Name())
					return nil
				}
			}

			for _, matchPattern := range matchPatterns {
				if loc := matchPattern.Regexp.FindStringIndex(info.Name()); loc != nil {
					fmt.Print(filepath.Dir(path) + "/")
					red("%s", info.Name())
					fmt.Print(" -> ")
					firstMatch := info.Name()[loc[0]:loc[1]]
					new := info.Name()[:loc[0]] + matchPattern.Regexp.ReplaceAllString(firstMatch, "${1}"+strings.Repeat("0", matchPattern.Nullen)+"${2}") + info.Name()[loc[1]:]
					green("%s\n", new)
					if doTheWork {
						rename(path, new)
					}
					return nil
				}
			}
		}

		return nil
	})
	if err != nil {
		fmt.Printf("error walking: %v\n", err)
		return
	}
}
