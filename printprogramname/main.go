package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]

	lastSlash := 0
	for i, r := range programName {
		if r == '/' {
			lastSlash = i + 1
		}
	}

	for _, r := range programName[lastSlash:] {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}
