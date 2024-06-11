package go_reloaded

import (
	"fmt"
	"strconv"
	"strings"
)

// UpperCase function is modifying text by Uppercasing them.
func UpperCase(s []string) []string {
	for i := 0; i < len(s); i++ {
		if strings.Contains(s[i], "up,") {
			number, err := strconv.Atoi(s[i+1][:len(s[i+1])-1])
			if err != nil || number > i {
				fmt.Printf("Error at conversion '%v' or 'up, %v' out of range.\n", err, number)
			} else {
				for j := i - number; j < i; j++ { // indices to modify
					s[j] = strings.ToUpper(s[j])
				}
			} // removing the (up, n) slices
			s = append(s[:i], s[i+2:]...)

		} else if strings.Contains(s[i], "up") {
			s[i-1] = strings.ToUpper(s[i-1])
			s = append(s[:i], s[i+1:]...) // removing the "(up)" slice
		}
	}
	return s
}
