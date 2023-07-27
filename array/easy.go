package array

import (
	"fmt"
	"math"
	"sort"
)

func easy() {

	arr := []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	fmt.Printf("largest element in arr: %v is %d.\n", arr, largestElementInArray(arr))

	arr = []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	sl, ss := secondLargestAnd2ndSmallestElement(arr)
	fmt.Printf("2nd largest element is %d and 2nd smallest element is %d in arr: %v.\n", sl, ss, arr)
	arr = []int{1}
	sl, ss = secondLargestAnd2ndSmallestElement(arr)
	fmt.Printf("2nd largest element is %d and 2nd smallest element is %d in arr: %v.\n", sl, ss, arr)

	arr = []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	fmt.Printf("Array: %v is sorted = %v\n", arr, isArraySorted(arr))
	arr = []int{2, 4, 6, 7, 8, 10}
	fmt.Printf("Array: %v is sorted = %v\n", arr, isArraySorted(arr))

	arr = []int{5, 7, 7, 9, 9, 9, 9, 11, 15, 17, 17, 18, 19, 25, 25}

	fmt.Printf("Array: %v after duplicate removal = %v\n", []int{5, 7, 7, 9, 9, 9, 9, 11, 15, 17, 17, 18, 19, 25, 25}, removeDuplicateForSortedArray(arr))

	arr1 := []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	fmt.Printf("Array: %v after duplicate removal = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, removeDuplicateForSortedArray(arr1))

	arr = []int{1, 1}

	fmt.Printf("Array: %v after duplicate removal = %v\n", []int{1, 1}, removeDuplicateForSortedArray(arr))

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	leftRotateArrayByOneBF(arr)
	fmt.Printf("Array: %v after left rotate by one = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, arr)

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	fmt.Printf("Array: %v after left rotate by one = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, leftRotateArrayByOneGolang(arr))

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	rotateArrayByKElement(arr, 3, "left")
	fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "left", arr)

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	rotateArrayByKElement(arr, 3, "right")
	fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "right", arr)

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "left", rotateArrayByKElementGoLang(arr, 3, "left"))

	fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "right", rotateArrayByKElementGoLang(arr, 3, "right"))

	arr = []int{1, 0, 2, 3, 0, 4, 0, 1}
	moveZeroToEndInArray(arr)
	fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{1, 0, 2, 3, 0, 4, 0, 1}, arr)

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	moveZeroToEndInArray(arr)
	fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, arr)
	arr = []int{0, 0, 0, 0, 0}
	moveZeroToEndInArray(arr)
	fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{0, 0, 0, 0, 0}, arr)

	arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	fmt.Printf("In Array: %v, %d present at position: %d\n", arr, 19, linearSearch(arr, 19))
	fmt.Printf("In Array: %v, %d present at position: %d\n", arr, 27, linearSearch(arr, 27))

	arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 3, 4, 4, 5, 11, 12}
	fmt.Printf("Union of Array, %v and %v is: %v\n", arr1, arr2, unionOfTwoArrayBF(arr1, arr2))
	fmt.Printf("Union of Array with Space Solution, %v and %v is: %v\n", arr1, arr2, unionOfTwoArrayWithSpace(arr1, arr2))
	fmt.Printf("Union of Two Sorted Array, %v and %v is: %v\n", arr1, arr2, unionOfTwoSortedArray(arr1, arr2))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	fmt.Printf("Missing Number in Array: %v is : %d\n", arr, findMissingNumber1(arr))

	arr = []int{1, 2, 3, 4, 5}
	fmt.Printf("Missing Number in Array: %v is : %d\n", arr, findMissingNumber(arr))

	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 10}
	fmt.Printf("Missing Number in Array: %v is : %d\n", arr, findMissingNumberXOR(arr))

	arr = []int{1, 1, 0, 1, 1, 1, 0, 1}
	fmt.Printf("Max Ones in Array: %v is : %d\n", arr, countMaxConsecutiveOns(arr))

	arr = []int{1, 1, 0, 1, 1, 1, 0, 1, 1, 1, 1}
	fmt.Printf("Max Ones in Array: %v is : %d\n", arr, countMaxConsecutiveOns(arr))

	arr = []int{0, 0, 0, 0}
	fmt.Printf("Max Ones in Array: %v is : %d\n", arr, countMaxConsecutiveOns(arr))

	arr = []int{4, 1, 2, 1, 2}
	fmt.Printf("Unique Number in Array: %v is : %d\n", arr, numberThatAppearsOnceWithSpace(arr))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Unique Number in Array: %v is : %d\n", arr, numberThatAppearsOnceWithXOR(arr))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Longest SubArray in Array: %v is : %d\n", arr, longestSubArrayWithSum(arr, 11))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Longest SubArray in Array: %v is : %d\n", arr, longestSubArrayWithSum(arr, 12))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Longest SubArray in Array - Two Pointer: %v is : %d\n", arr, longestSubArrayWithSumTP(arr, 11))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Longest SubArray in Array - Two Pointer: %v is : %d\n", arr, longestSubArrayWithSumTP(arr, 12))

	arr = []int{-1, 1, 1}
	fmt.Printf("Longest SubArray in Array: %v is : %d\n", arr, longestSubArrayWithSum2(arr, 1))

	arr = []int{2, 3, 5}
	fmt.Printf("Longest SubArray in Array: %v is : %d\n", arr, longestSubArrayWithSum2(arr, 5))

	arr = []int{2, 3, 5, -2, -3, 4}
	fmt.Printf("Longest SubArray in Array: %v is : %d\n", arr, longestSubArrayWithSum2(arr, 5))

	arr = []int{4, 1, 2, 1, 2, 4, 5, 6, 7, 7, 6}
	fmt.Printf("Longest SubArray in Array - Hash Map: %v is : %d\n", arr, longestSubArrayWithSumHMap(arr, 12))

	arr = []int{-1, 1, 1}
	fmt.Printf("Longest SubArray in Array - Hash Map: %v is : %d\n", arr, longestSubArrayWithSumHMap(arr, 1))

	arr = []int{2, 3, 5}
	fmt.Printf("Longest SubArray in Array - Hash Map: %v is : %d\n", arr, longestSubArrayWithSumHMap(arr, 5))

	arr = []int{2, 3, 5, -2, -3, 4}
	fmt.Printf("Longest SubArray in Array - Hash Map: %v is : %d\n", arr, longestSubArrayWithSumHMap(arr, 5))

}

func largestElementInArray(arr []int) int {
	lElem := arr[0]

	for _, v := range arr {
		if v > lElem {
			lElem = v
		}
	}
	return lElem
}

func secondLargestAnd2ndSmallestElement(arr []int) (int, int) {
	slEle, ssEle, lEle, sEle := math.MinInt, math.MaxInt, math.MinInt, math.MaxInt

	for _, v := range arr {
		if v > lEle {
			slEle = lEle
			lEle = v
		}
		if v < sEle {
			ssEle = sEle
			sEle = v
		}
		if v < lEle && v > slEle {
			slEle = v
		}
		if v > sEle && v < ssEle {
			ssEle = v
		}
	}
	if slEle == math.MinInt {
		slEle = -1
	}

	if ssEle == math.MaxInt {
		ssEle = -1
	}

	return slEle, ssEle
}

func isArraySorted(arr []int) bool {
	ans := true

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			ans = false
			break
		}
	}

	return ans
}

/*
Given an integer array sorted in non-decreasing order, remove the duplicates in place such that each unique element appears only once. The relative order of the elements should be kept the same.

If there are k elements after removing the duplicates, then the first k elements of the array should hold the final result. It does not matter what you leave beyond the first k elements.

Note: Return k after placing the final result in the first k slots of the array.
*/
func removeDuplicateForSortedArray(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	replaceWith := 1
	for i := 1; i < len(arr); i++ {
		if arr[i] != arr[i-1] {
			arr[replaceWith] = arr[i]
			replaceWith++
		}
	}

	return arr[:replaceWith]
}

func leftRotateArrayByOneBF(arr []int) {
	temp := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
}

func leftRotateArrayByOneGolang(arr []int) []int {
	arr = append(arr[1:], arr[0])
	return arr
}

func rotateArrayByKElement(arr []int, position int, place string) {
	if place == "left" {
		temp := make([]int, position)
		for i := 0; i < position; i++ {
			temp[i] = arr[i]
		}
		for i := 0; i < len(arr)-position; i++ {
			arr[i] = arr[i+position]
		}
		for i := 0; i < len(temp); i++ {
			arr[len(arr)-position+i] = temp[i]
		}
	} else {
		temp := make([]int, position)
		for i := 0; i < position; i++ {
			temp[i] = arr[len(arr)-position+i]
		}
		for i := len(arr) - position - 1; i >= 0; i-- {
			arr[i+position] = arr[i]
		}
		copy(arr, temp)
	}
}

func rotateArrayByKElementGoLang(arr []int, position int, place string) []int {
	if place == "left" {
		arr = append(arr[position:], arr[:position]...)
		return arr
	}
	l := len(arr)
	arr = append(arr[l-position:], arr[:l-position]...)
	return arr
}

func moveZeroToEndInArray(arr []int) {
	l := len(arr)
	endZeroPos := l - 1
	startPosition := 0

	for startPosition < endZeroPos {

		for startPosition < l && arr[startPosition] != 0 {
			startPosition++
		}
		for endZeroPos > -1 && arr[endZeroPos] == 0 {
			endZeroPos--
		}
		if startPosition < endZeroPos {
			arr[startPosition], arr[endZeroPos] = arr[endZeroPos], arr[startPosition]
			startPosition++
			endZeroPos--
		}
	}
}

func linearSearch(arr []int, item int) int {
	for i, v := range arr {
		if v == item {
			return i
		}
	}
	return -1
}

func unionOfTwoArrayBF(arr1, arr2 []int) []int {
	arr := []int{}

	for _, v := range arr1 {
		if linearSearch(arr, v) == -1 {
			arr = append(arr, v)
		}
	}
	for _, v := range arr2 {
		if linearSearch(arr, v) == -1 {
			arr = append(arr, v)
		}
	}

	return arr
}

func unionOfTwoArrayWithSpace(arr1, arr2 []int) []int {

	m := make(map[int]bool)

	for _, v := range arr1 {
		m[v] = true
	}
	for _, v := range arr2 {
		m[v] = true
	}

	arr, i := make([]int, len(m)), 0
	for k := range m {
		arr[i] = k
		i++
	}
	sort.Ints(arr)
	return arr
}

func unionOfTwoSortedArray(arr1, arr2 []int) []int {

	arr := []int{}
	i, j, k := 0, 0, -1
	la, lb := len(arr1), len(arr2)

	for i < la && j < lb {
		for i < la && arr1[i] < arr2[j] {
			if k == -1 || arr[k] != arr1[i] {
				arr = append(arr, arr1[i])
				k++
			}
			i++
		}

		for i < la && j < lb && arr2[j] <= arr1[i] {
			if k == -1 || arr[k] != arr2[j] {
				arr = append(arr, arr2[j])
				k++
			}
			j++
		}
	}

	for i < la {
		if k == -1 || arr[k] != arr1[i] {
			arr = append(arr, arr1[i])
			k++
		}
		i++
	}

	for j < lb {
		if k == -1 || arr[k] != arr2[j] {
			arr = append(arr, arr2[j])
			k++
		}
		j++
	}

	return arr
}

/*
Problem Statement: Given an integer N and an array of size N-1 containing N-1 numbers between 1 to N.
Find the number(between 1 to N), that is not present in the given array.
*/
func findMissingNumber1(arr []int) int {
	missingNumber := 1

	for _, v := range arr {
		if missingNumber != v {
			return missingNumber
		}
		missingNumber++
	}

	return missingNumber
}

func findMissingNumber(arr []int) int {
	l := len(arr)
	sum := (l + 1) * (l + 2) / 2
	arrSum := 0
	for i := 0; i < l; i++ {
		arrSum += arr[i]
	}
	return sum - arrSum
}

func findMissingNumberXOR(arr []int) int {
	missingNumber := 0
	l := len(arr)
	for i := 1; i <= l+1; i++ {
		missingNumber = missingNumber ^ i
	}
	for i := 0; i < l; i++ {
		missingNumber = missingNumber ^ arr[i]
	}
	return missingNumber
}

/*
Problem Statement: Given an array that contains only 1 and 0 return the count of maximum consecutive ones in the array.
*/
func countMaxConsecutiveOns(arr []int) int {
	maxOnes := 0
	count := 0

	for _, v := range arr {
		if v == 1 {
			count++
			continue
		}
		if count > maxOnes {
			maxOnes = count
		}
		count = 0
	}

	if count > maxOnes {
		maxOnes = count
	}
	return maxOnes
}

/*
Find the number that appears once, and the other numbers twice
Problem Statement: Given a non-empty array of integers arr, every element appears twice except for one. Find that single one.
*/
func numberThatAppearsOnceWithSpace(arr []int) int {
	m := make(map[int]bool)

	for _, v := range arr {
		m[v] = !m[v]
	}

	for k, v := range m {
		if v {
			return k
		}
	}
	return -1
}

func numberThatAppearsOnceWithXOR(arr []int) int {
	num := 0

	for _, v := range arr {
		num = num ^ v
	}

	return num
}

/*
Longest SubArray with given Sum K(Positives)
Problem Statement: Given an array and a sum k, we need to print the length of the longest sub-array that sums to k.
*/
func longestSubArrayWithSum(arr []int, sum int) int {
	ans, count, latestSum, initialIndex := 0, 0, 0, 0

	for i := 0; i < len(arr); i++ {
		latestSum = latestSum + arr[i]
		if latestSum < sum {
			count++
			continue
		}
		if latestSum > sum {
			for initialIndex <= i {
				latestSum = latestSum - arr[initialIndex]
				if latestSum > sum {
					initialIndex++
					count--
					continue
				}

				if latestSum == sum {
					if count > ans {
						ans = count
					}
					initialIndex = i + 1
					latestSum = 0
					count = 0
					break
				}
				initialIndex++
				break
			}
			continue
		}

		if count > ans {
			ans = count
		}
		initialIndex = i + 1
		latestSum = 0
		count = 0
	}

	return ans
}

func longestSubArrayWithSumTP(arr []int, sum int) int {
	ans, latestSum, initialIndex := 0, 0, 0

	for i := 0; i < len(arr); i++ {
		latestSum = latestSum + arr[i]
		if latestSum < sum {
			continue
		}
		if latestSum == sum {
			if i-initialIndex > ans {
				ans = i - initialIndex + 1
			}
			continue
		}
		for initialIndex <= i {
			latestSum = latestSum - arr[initialIndex]
			if latestSum == sum {
				if i-initialIndex > ans {
					ans = i - initialIndex
				}
				initialIndex++
				break
			}
			initialIndex++
			if latestSum < sum {
				break
			}
		}

	}

	return ans
}

/*
Longest SubArray with given Sum K(Positives and Negatives)
Problem Statement: Given an array and a sum k, we need to print the length of the longest sub-array that sums to k.

Example 1:
Input Format: N = 3, k = 5, array[] = {2,3,5}
Result: 2
Explanation: The longest subarray with sum 5 is {2, 3}. And its length is 2.

Example 2:
Input Format: N = 3, k = 1, array[] = {-1, 1, 1}
Result: 3
Explanation: The longest subarray with sum 1 is {-1, 1, 1}. And its length is 3.
*/
func longestSubArrayWithSum2(arr []int, sum int) int {
	ans, latestSum, j := 0, 0, 0

	for i := 0; i < len(arr); i++ {
		latestSum = latestSum + arr[i]

		if latestSum == sum {
			if i+1 > ans {
				ans = i + 1
			}
			continue
		}
		j = 0
		asum := latestSum
		for j <= i {
			asum = asum - arr[j]
			if asum == sum {
				if i-j+1 > ans {
					ans = i - j + 1
				}
				break
			}
			j++
		}
	}

	return ans
}

func longestSubArrayWithSumHMap(arr []int, sum int) int {
	ans, latestSum, m := 0, 0, make(map[int]int)

	for i := 0; i < len(arr); i++ {

		latestSum = latestSum + arr[i]
		if latestSum == sum {
			ans = i + 1
		}

		if v, ok := m[latestSum-sum]; ok {
			if i-v > ans {
				ans = i - v
			}
		}

		if _, ok := m[latestSum]; !ok {
			m[latestSum] = i
		}
	}

	return ans
}
