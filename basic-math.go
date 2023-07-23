package main

func countDigit(n int) int {
	ans := 0

	for n > 0 {
		n = n / 10
		ans++
	}

	return ans
}

func reverseANumber(n int) int {
	ans := 0

	for n > 0 {
		ans = ans*10 + n%10
		n = n / 10
	}

	return ans
}

func checkNumberIsPalindrome(n int) bool {
	rn := reverseANumber(n)
	return rn == n
}
