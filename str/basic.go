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
	return strings.TrimLeft(reverseString, " ")
}

func isStringAnagram(str string) bool {
	l := len(str)
	for i := 0; i < l/2; i++ {
		if str[i] != str[l-1-i] {
			return false
		}
	}

	return true
}

func isTwoStringsAreAnagram(str1, str2 string) bool {
	m := make(map[rune]int)
	if len(str1) != len(str2) {
		return false
	}

	for _, v := range str1 {
		m[v]++
	}

	for _, v := range str2 {
		if m[v] == 0 {
			return false
		}
		m[v]--
		if m[v] == 0 {
			delete(m, v)
		}
	}

	return !(len(m) > 0)
}
