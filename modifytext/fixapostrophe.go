package go_reloaded

import (
	"regexp"
	"strings"
)

func FixApostrophe(s []string) []string {
	pattern := regexp.MustCompile(`'\s*(.*?)\s*'`)
	newString := pattern.ReplaceAllString(strings.Join(s, " "), "'$1' ")
	// text := newString + "\n"
	return strings.Fields(newString)
}
