package go_reloaded

import (
	"fmt"
	"strconv"
)

// function to convert hexadecimal to its decimal equivalent
func Hexadecimal(s []string) []string {
	for i, text := range s {
		if text == "(hex)" {
			num, err := strconv.ParseInt(s[i-1], 16, 64) // store the decimal value in; parse the literal in base 16 for hexadecimal
			if err != nil {
				fmt.Println("Error parsing the hexadecimal to integer:", err)
			}
			s[i-1] = fmt.Sprint(num)      // convert the decimal value back to string
			s = append(s[:i], s[i+1:]...) // remove the elenment to be processed
		}
	}
	return s
}
