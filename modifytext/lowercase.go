package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// function to modify text to lower case
func LowerCase(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "low,") { // search for text to modify
			number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
			if err != nil || number > i {
				fmt.Printf("Error at conversion; '%v' or 'low, %v' out of range.\n", err, number)
				
			} else {
				for j := i - number; j < i; j++ { // selecting text to modify
					s[j] = strings.ToLower(s[j])
				}
			}
			s = append(s[:i], s[i+2:]...)

		} else if strings.Contains(s[i], "low") { // search for text to modify
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
