package go_reloaded

import "strings"

func FixPunctuation(s []string) []string {
	punctuations := []string{",", ".", "!", "?", ":", ";"}

	for i, word := range s {
		for _, mark := range punctuations {
			if strings.HasPrefix(word, string(mark)) && !strings.HasSuffix(word, string(mark)) && s[i] != s[len(s)-1] {
				s[i-1] += string(s[i][0])
				s[i] = word[1:]
			}
		}
	}

	for i, word := range s {
		for _, mark := range punctuations {
			if strings.HasPrefix(word, string(mark)) && strings.HasSuffix(word, string(mark)) && s[i] != s[len(s)-1] {
				s[i-1] += s[i]
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	for i, word := range s {
		for _, mark := range punctuations {
			if strings.HasPrefix(word, string(mark)) && s[i] == s[len(s)-1] {
				s[i-1] += s[i]
				s = s[:len(s)-1]
			}
		}
	}

	return s
}
