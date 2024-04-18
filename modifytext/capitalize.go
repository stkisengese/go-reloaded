package go_reloaded

import (
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func Capitalize(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "cap,") {
			number, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
			for j := i - number; j < i; j++ {
				s[j] = cases.Title(language.English).String(s[j])
			}
			s = append(s[:i], s[i+2:]...)
		} else if strings.Contains(s[i], "cap") {
			s[i-1] = cases.Title(language.English).String(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}
