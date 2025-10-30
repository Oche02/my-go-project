package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	baseLen := len(base)
	if baseLen < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}

	if nbr == 0 {
		z01.PrintRune(rune(base[0]))
		return
	}

	// Handle negative numbers
	if nbr < 0 {
		z01.PrintRune('-')
		// Convert to positive safely (handle int min)
		if nbr == -nbr {
			// if nbr == MinInt, invert manually
			printInBase(-(nbr / baseLen), base)
			z01.PrintRune(rune(base[-(nbr % baseLen)]))
			return
		}
		nbr = -nbr
	}

	printInBase(nbr, base)
}

// Helper function to check base validity
func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for i := 0; i < len(base); i++ {
		if base[i] == '+' || base[i] == '-' {
			return false
		}
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return false
			}
		}
	}
	return true
}

// Recursive function to print number in base
func printInBase(n int, base string) {
	baseLen := len(base)
	if n >= baseLen {
		printInBase(n/baseLen, base)
	}
	z01.PrintRune(rune(base[n%baseLen]))
}
