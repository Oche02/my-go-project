package piscine

func IsLower(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < 'a' || r > 'z' {
			return false
		}
	}
	return true
}
