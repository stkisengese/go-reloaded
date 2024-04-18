package go_reloaded

import (
	"fmt"
	"strconv"
)

func Binary(s []string) []string {
	for i, text := range s {
		if text == "(bin)" {
			num, err := strconv.ParseInt(s[i-1], 2, 64) // parse the literal in base 2 for binary and store its value in 'num'
			if err != nil {
				fmt.Println("Error parsing integer:", err)
			}
			s[i-1] = fmt.Sprint(num)      // convert the decimal back to string
			s = append(s[:i], s[i+1:]...) // remove the "(bin)" element
		}
	}
	return s
}
