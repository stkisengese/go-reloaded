package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// function to captalize every first letter in a word
func Capitalize(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "cap,") { // search for number of words to modify
			number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
			
			if err != nil || number > i {
				fmt.Printf("Error at conversion; '%v' or 'cap, %v' is out of range.\n", err, number)
			} else {
				for j := i - number; j < i; j++ { // the indices to capitalize
					s[j] = cases.Title(language.English).String(s[j])
				}
			}
			s = append(s[:i], s[i+2:]...)

		} else if strings.Contains(s[i], "cap") {
			s[i-1] = cases.Title(language.English).String(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
