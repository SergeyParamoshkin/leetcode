package palindromic

func FirstPalindrome(words []string) string {
	for _, word := range words {
		if len(word) == 1 {
			return word
		}

		left := 0
		rigth := len(word) - 1

		for left < rigth {
			if word[left] != word[rigth] {
				break
			}
			left++
			rigth--

			if left >= rigth {
				return word
			}
		}
	}

	return ""
}
