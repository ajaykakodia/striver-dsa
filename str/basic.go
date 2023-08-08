package str

import (
	"strings"
)

func reverseWordInString(s string) string {
	reverseString := ""
	words := make([]string, 0)
	words = strings.Split(s, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			reverseString = reverseString + " " + words[i]
		}
	}
	return reverseString
}
