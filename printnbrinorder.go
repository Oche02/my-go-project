package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}

	var digits []rune

	for n > 0 {
		d := n % 10
		digits = append(digits, rune('0'+d))
		n /= 10
	}

	for i := 0; i < len(digits)-1; i++ {
		for j := i + 1; j < len(digits); j++ {
			if digits[i] > digits[j] {
				digits[i], digits[j] = digits[j], digits[i]
			}
		}
	}

	for _, r := range digits {
		z01.PrintRune(r)
	}
}
