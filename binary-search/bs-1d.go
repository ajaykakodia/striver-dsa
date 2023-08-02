package binarysearch

/*
Coding problem example:
Problem statement: You are given a sorted array of integers and a target, your task is to search for the target in the given array. Assume the given array does not contain any duplicate numbers.

Let’s say the given array is = {3, 4, 6, 7, 9, 12, 16, 17} and target = 6.
*/
func binarySearchArray(arr []int, k int) bool {
	l := len(arr)
	start, end, mid := 0, l-1, 0
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == k {
			return true
		}
		if arr[mid] > k {
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	return false
}

func binarySearchArrayR(arr []int, k int) bool {
	if len(arr) == 0 {
		return false
	}

	start, end := 0, len(arr)-1
	mid := (start + end) / 2

	if arr[mid] == k {
		return true
	}
	if arr[mid] > k {
		return binarySearchArrayR(arr[:mid], k)
	}
	return binarySearchArrayR(arr[mid+1:], k)
}

/*
Implement Lower Bound
Problem Statement: Given a sorted array of N integers and an integer x, write a program to find the lower bound of x.

Example 1:
Input Format: N = 4, arr[] = {1,2,2,3}, x = 2
Result: 1
Explanation: Index 1 is the smallest index such that arr[1] >= x.

Example 2:
Input Format: N = 5, arr[] = {3,5,8,15,19}, x = 9
Result: 3
Explanation: Index 3 is the smallest index such that arr[3] >= x.
*/
func lowerBoundOfElement(arr []int, k int) int {
	l := len(arr)
	ans := k
	start, end, mid := 0, l-1, 0
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] >= k {
			ans = mid
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	return ans
}

/*
Implement Upper Bound
Problem Statement: Given a sorted array of N integers and an integer x, write a program to find the upper bound of x.

Example 1:
Input Format: N = 4, arr[] = {1,2,2,3}, x = 2
Result: 3
Explanation: Index 3 is the smallest index such that arr[3] > x.

Example 2:
Input Format: N = 6, arr[] = {3,5,8,9,15,19}, x = 9
Result: 4
Explanation: Index 4 is the smallest index such that arr[4] > x.
*/
func upperBoundOfElement(arr []int, k int) int {
	l := len(arr)
	ans := k
	start, end, mid := 0, l-1, 0
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] > k {
			ans = mid
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	return ans
}
