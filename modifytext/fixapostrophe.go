package go_reloaded

import (
	"regexp"
	"strings"
)

// FixApostrophe function to punctuate single quote marks appropriately.
func FixApostrophe(s []string) []string {
	pattern := regexp.MustCompile(`'\s*(.*?)\s*'`)
	newString := pattern.ReplaceAllString(strings.Join(s, " "), "'$1' ")

	return strings.Fields(newString)
}
