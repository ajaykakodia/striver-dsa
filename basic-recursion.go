package main

import "fmt"

//  R stands for recursion

func basicRecursion() {
	printNumberR(20)
	fmt.Println()
	printNameNTimesR("Ajay", 3)
	fmt.Println()
	printReverseNumberR(9)
	fmt.Println()
	fmt.Println(sumOfFirstNaturalNumberR(9))
	fmt.Println(sumOfFirstNaturalNumberR(6))

	fmt.Println(factorial(5))
	fmt.Println(factorialR(5))

	fmt.Println(factorial(10))
	fmt.Println(factorialR(10))

	arr := []int{4, 3, 6, 35, 75, 34, 78, 22}
	fmt.Printf("arr before reverse: %v\n", arr)
	reverseArrayR(arr, 0, len(arr)-1)
	fmt.Printf("arr after reverse: %v\n", arr)
	fmt.Println(isStringPalindromeR("ABCDCBA", 0, len("ABCDCBA")-1))
	fmt.Println(isStringPalindromeR("majaaaya", 0, len("majaaaya")-1))
	fmt.Println(isStringPalindrome2R("ABCDCBA", 0))

	fmt.Println(getFibonacciR(8))
	fmt.Println(getFibonacciOR(8, make(map[int]int)))
}

// including 0 into the print
func printNumberR(n int) {
	if n == -1 {
		return
	}
	printNumberR(n - 1)
	fmt.Print(n, " ")
}

func printNameNTimesR(name string, n int) {
	if n == 0 {
		return
	}
	fmt.Print(name, " ")
	printNameNTimesR(name, n-1)
}

func printReverseNumberR(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printReverseNumberR(n - 1)
}

func sumOfFirstNaturalNumberR(n int) int {
	if n == 1 {
		return 1
	}
	return n + sumOfFirstNaturalNumberR(n-1)
}

func factorial(n int) int {

	fact := 1

	for n > 0 {
		fact = fact * n
		n = n - 1
	}
	return fact
}

func factorialR(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorialR(n-1)
}

func reverseArrayR(arr []int, s, e int) {
	if s >= e {
		return
	}
	arr[s], arr[e] = arr[e], arr[s]
	reverseArrayR(arr, s+1, e-1)
}

func isStringPalindromeR(str string, s, e int) bool {
	if s >= e {
		return true
	}
	if str[s] != str[e] {
		return false
	}
	return isStringPalindromeR(str, s+1, e-1)
}

func isStringPalindrome2R(str string, currentPosition int) bool {
	l := len(str)
	if currentPosition >= l/2 {
		return true
	}
	if str[currentPosition] != str[l-currentPosition-1] {
		return false
	}
	return isStringPalindrome2R(str, currentPosition+1)
}

func getFibonacciR(n int) int {
	if n <= 1 {
		return n
	}
	num := getFibonacciR(n-1) + getFibonacciR(n-2)
	return num
}

func getFibonacciOR(n int, m map[int]int) int {
	if n <= 1 {
		return n
	}
	f, s := 0, 0
	if v, ok := m[n-1]; ok {
		f = v
	} else {
		f = getFibonacciOR(n-1, m)
	}
	if v, ok := m[n-2]; ok {
		s = v
	} else {
		s = getFibonacciOR(n-2, m)
	}
	m[n] = f + s
	return m[n]
}
