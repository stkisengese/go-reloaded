package go_reloaded

// VowelArticle has logic for adding article for words starting with vowels
func VowelArticle(s []string) []string {
	vowels := []byte{'a', 'e', 'i', 'o', 'u', 'h', 'A', 'E', 'I', 'O', 'U', 'H'}

	for i, word := range s {
		if i < len(s)-1 {
			for _, letter := range vowels {
				if word == "a" && s[i+1][0] == letter {
					s[i] = "an"
					s = append(s[:i], s[i:]...)

				} else if word == "A" && s[i+1][0] == letter {
					s[i] = "An"
					s = append(s[:i], s[i:]...)
				}
			}
		}
	}
	return s
}
