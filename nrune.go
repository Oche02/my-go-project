package piscine

func NRune(s string, n int) rune {
	if n <= 0 {
		return 0
	}

	count := 1
	for _, r := range s {
		if count == n {
			return r
		}
		count++
	}
	return 0
}
