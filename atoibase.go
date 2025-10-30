package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLen := len(base)
	result := 0

	for _, r := range s {
		index := indexOf(r, base)
		if index == -1 {
			return 0
		}
		result = result*baseLen + index
	}
	return result
}

// Check if base is valid
func atoiBase(base string) bool {
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

// Find the index of a rune in the base string
func indexOf(r rune, base string) int {
	for i, b := range base {
		if b == r {
			return i
		}
	}
	return -1
}
