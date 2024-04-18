package go_reloaded

import (
	"strconv"
	"strings"
)

func UpperCase(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "up") {
			if strings.Contains(s[i], "up,") {
				number, _ := strconv.Atoi(s[i+1][:len(s[i+1])-1])
				for j := i - number; j < i; j++ {
					s[j] = strings.ToUpper(s[j])
				}
				s = append(s[:i], s[i+2:]...)
			} else {
				s[i-1] = strings.ToUpper(s[i-1])
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	return s
}
