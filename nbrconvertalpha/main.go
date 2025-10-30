package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	isUpper := false

	if len(args) > 0 && args[0] == "--upper" {
		isUpper = true
		args = args[1:]
	}

	printed := false

	for _, arg := range args {
		n := 0
		valid := true

		for _, r := range arg {
			if r < '0' || r > '9' {
				valid = false
				break
			}
			n = n*10 + int(r-'0')
		}

		if !valid || n < 1 || n > 26 {
			z01.PrintRune(' ')
			printed = true
			continue
		}

		if isUpper {
			z01.PrintRune(rune('A' + n - 1))
		} else {
			z01.PrintRune(rune('a' + n - 1))
		}
		printed = true
	}

	if printed {
		z01.PrintRune('\n')
	}
}
