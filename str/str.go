package str

import "fmt"

func StringProblems() {
	fmt.Println("We are in string methods")
	fmt.Println(reverseWordInString("the sky is blue"))
	fmt.Println(reverseWordInString("  hello world  "))
	fmt.Println(reverseWordInString("a good   example"))

	fmt.Printf("%s is anagram: %t\n", "ragar", isStringAnagram("ragar"))

	fmt.Printf("%s and %s are anagram: %t\n", "CAT", "ACT", isTwoStringsAreAnagram("CAT", "ACT"))

	fmt.Printf("%s and %s are anagram: %t\n", "INTEGER", "TEGERNI", isTwoStringsAreAnagram("INTEGER", "TEGERNI"))
}
