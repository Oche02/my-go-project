package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord {
				// Capitalize first letter if it's lowercase
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				newWord = false
			} else {
				// Lowercase other letters
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			newWord = true // next character starts a new word
		}
	}

	return string(runes)
}
