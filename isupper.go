package piscine

func IsUpper(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < 'A' || r > 'Z' {
			return false
		}
	}
	return true
}
