package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	// Step 1: Convert nbr (string) from baseFrom → base 10 integer
	num := 0
	for _, ch := range nbr {
		index := indexInBase(baseFrom, ch)
		num = num*len(baseFrom) + index
	}

	// Step 2: If the number is 0, return the first symbol of baseTo
	if num == 0 {
		return string(baseTo[0])
	}

	// Step 3: Convert the base 10 integer → baseTo string
	result := ""
	for num > 0 {
		remainder := num % len(baseTo)
		result = string(baseTo[remainder]) + result
		num /= len(baseTo)
	}

	return result
}

// Helper function: find index of a character in the base string
func indexInBase(base string, ch rune) int {
	for i, c := range base {
		if c == ch {
			return i
		}
	}
	return -1
}
