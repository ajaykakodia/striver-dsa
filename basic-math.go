package main

import (
	"fmt"
	"math"
)

func basicMath() {

	fmt.Println(countDigit(1234))
	fmt.Println(countDigit(43430000))

	fmt.Println(reverseANumber(123))
	fmt.Println(reverseANumber(43430000))

	fmt.Println(checkNumberIsPalindrome(123))
	fmt.Println(checkNumberIsPalindrome(121))

	fmt.Println("GCD of numbers - 256, 56: ", gcdBF(256, 56))
	fmt.Println("GCD of numbers - 34, 17: ", gcdBF(34, 17))
	fmt.Println("GCD of numbers - 11, 13: ", gcdBF(11, 13))

	fmt.Println("GCD of numbers - 256, 56: ", gcd(256, 56))
	fmt.Println("GCD of numbers - 34, 17: ", gcd(34, 17))
	fmt.Println("GCD of numbers - 11, 13: ", gcd(11, 13))

	fmt.Println(checkNumberIsArmstrong(153))
	fmt.Println(checkNumberIsArmstrong(170))

	printAllDivisorBF(26)
	printAllDivisor(36)
	printAllDivisor(97)

	fmt.Println(isPrimeNumberBF(17))
	fmt.Println(isPrimeNumberBF(47))
	fmt.Println(isPrimeNumberBF(129))

	fmt.Println(isPrimeNumber(17))
	fmt.Println(isPrimeNumber(47))
	fmt.Println(isPrimeNumber(129))
}

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

// BruteForce
func gcdBF(a, b int) int {
	ans := 1

	minNum := int(math.Min(float64(a), float64(b)))

	for i := minNum; i > 0; i-- {
		if a%i == 0 && b%i == 0 {
			ans = i
			break
		}
	}

	return ans
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func checkNumberIsArmstrong(n int) bool {
	m, numOfDigits, sum := n, 0, 0
	for m != 0 {
		m = m / 10
		numOfDigits++
	}

	m = n

	for m != 0 {
		sum = sum + int(math.Pow(float64(m%10), float64(numOfDigits)))
		m = m / 10
	}

	return n == sum
}

// BruteForce
func printAllDivisorBF(n int) {
	arr := []int{}
	arr = append(arr, 1)
	for i := 2; i <= n/2+1; i++ {
		if n%i == 0 {
			arr = append(arr, i)
		}
	}

	arr = append(arr, n)

	fmt.Printf("All divisor of %d - %v\n", n, arr)
}

func printAllDivisor(n int) {
	arr := []int{}
	sqr := int(math.Sqrt(float64(n)))
	for i := 1; i <= sqr; i++ {
		if n%i == 0 {
			arr = append(arr, i)
			div := n / i
			if i != div {
				arr = append(arr, div)
			}
		}
	}

	fmt.Printf("All divisor of %d - %v\n", n, arr)
}

// BruteForce
func isPrimeNumberBF(n int) bool {
	for i := 2; i < n/2+1; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isPrimeNumber(n int) bool {
	sqr := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqr; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
