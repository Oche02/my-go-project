package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	// If no arguments, print newline
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	}

	args := os.Args[1:]
	var allRunes []rune
	var vowels []rune

	// Combine all arguments into one slice of runes, including spaces
	for i, arg := range args {
		for _, r := range arg {
			allRunes = append(allRunes, r)
			if isVowel(r) {
				vowels = append(vowels, r)
			}
		}
		if i != len(args)-1 {
			allRunes = append(allRunes, ' ')
		}
	}

	// Reverse the vowels
	for i, j := 0, len(vowels)-1; i < j; i, j = i+1, j-1 {
		vowels[i], vowels[j] = vowels[j], vowels[i]
	}

	// Replace vowels in order with reversed ones
	vIndex := 0
	for _, r := range allRunes {
		if isVowel(r) {
			z01.PrintRune(vowels[vIndex])
			vIndex++
		} else {
			z01.PrintRune(r)
		}
	}

	z01.PrintRune('\n')
}
