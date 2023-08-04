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

/*
Search Insert Position
Problem Statement: You are given a sorted array arr of distinct values and a target value x. You need to search for the index of the target value in the array.

If the value is present in the array, then return its index. Otherwise, determine the index where it would be inserted in the array while maintaining the sorted order.

Example 1:
Input Format: arr[] = {1,2,4,7}, x = 6
Result: 3
Explanation: 6 is not present in the array. So, if we will insert 6 in the 3rd index(0-based indexing), the array will still be sorted. {1,2,4,6,7}.

Example 2:
Input Format: arr[] = {1,2,4,7}, x = 2
Result: 1
Explanation: 2 is present in the array and so we will return its index i.e. 1.

*/
func searchInsertPosition(arr []int, ele int) int {
	ans := len(arr)
	start, end := 0, len(arr)-1
	mid := 0

	for start <= end {
		mid = (start + end) / 2

		if arr[mid] >= ele {
			ans = mid
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	return ans
}

/*
Floor and Ceil in Sorted Array
Problem Statement: You’re given an sorted array arr of n integers and an integer x. Find the floor and ceiling of x in arr[0..n-1].
The floor of x is the largest element in the array which is smaller than or equal to x.
The ceiling of x is the smallest element in the array greater than or equal to x.

Example 1:
Input Format: n = 6, arr[] ={3, 4, 4, 7, 8, 10}, x= 5
Result: 4 7
Explanation: The floor of 5 in the array is 4, and the ceiling of 5 in the array is 7.

Example 2:
Input Format: n = 6, arr[] ={3, 4, 4, 7, 8, 10}, x= 8
Result: 8 8
Explanation: The floor of 8 in the array is 8, and the ceiling of 8 in the array is also 8.
*/
func floorNCeil(nums []int, target int) (int, int) {
	l := len(nums)
	floor, ceil, start, end, mid := nums[l-1], nums[l-1], 0, l-1, 0

	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return nums[mid], nums[mid]
		}
		if nums[mid] > target {
			ceil = nums[mid]
			end = mid - 1
			continue
		}
		floor = nums[mid]
		start = mid + 1
	}

	return floor, ceil
}

/*
Last occurrence in a sorted array
Given a sorted array of N integers, write a program to find the index of the last occurrence of the target key. If the target is not found then return -1.

Note: Consider 0 based indexing

Example 1:
Input: N = 7, target=13, array[] = {3,4,13,13,13,20,40}
Output: 4
Explanation: As the target value is 13 , it appears for the first time at index number 2.

Example 2:
Input: N = 7, target=60, array[] = {3,4,13,13,13,20,40}
Output: -1
Explanation: Target value 60 is not present in the array
*/
func lastOccurrenceOfKey(nums []int, target int) int {
	l := len(nums)
	ans := -1
	start, end, mid := 0, l-1, 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			ans = mid
			start = mid + 1
			continue
		}
		if nums[mid] < target {
			start = mid + 1
			continue
		}
		end = mid - 1
	}

	return ans
}

/*
Count Occurrences in Sorted Array
Problem Statement: You are given a sorted array containing N integers and a number X, you have to find the occurrences of X in the given array.

Example 1:
Input: N = 7,  X = 3 , array[] = {2, 2 , 3 , 3 , 3 , 3 , 4}
Output: 4
Explanation: 3 is occurring 4 times in
the given array so it is our answer.

Example 2:
Input: N = 8,  X = 2 , array[] = {1, 1, 2, 2, 2, 2, 2, 3}
Output: 5
Explanation: 2 is occurring 5 times in the given array so it is our answer.

*/
func countOccurrenceOfKey(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	mid := (start + end) / 2
	if nums[mid] == target {
		return 1 + countOccurrenceOfKey(nums[:mid], target) + countOccurrenceOfKey(nums[mid+1:], target)
	}
	if nums[mid] > target {
		return countOccurrenceOfKey(nums[:mid], target)
	}
	return countOccurrenceOfKey(nums[mid+1:], target)
}

/*
Search Element in a Rotated Sorted Array
Problem Statement: Given an integer array arr of size N, sorted in ascending order (with distinct values) and a target value k. Now the array is rotated at some pivot point unknown to you. Find the index at which k is present and if k is not present return -1.

Example 1:
Input Format: arr = [4,5,6,7,0,1,2,3], k = 0
Result: 4
Explanation: Here, the target is 0. We can see that 0 is present in the given rotated sorted array, nums. Thus, we get output as 4, which is the index at which 0 is present in the array.

Example 2:
Input Format: arr = [4,5,6,7,0,1,2], k = 3
Result: -1
Explanation: Here, the target is 3. Since 3 is not present in the given rotated sorted array. Thus, we get the output as -1.

*/

func searchElementInRotatedArray(arr []int, target int) int {

	lowIndex, highIndex := 0, len(arr)-1
	mid := 0
	for lowIndex <= highIndex {
		mid = (lowIndex + highIndex) / 2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > arr[lowIndex] {
			if target <= arr[mid] && target >= arr[lowIndex] {
				highIndex = mid - 1
			} else {
				lowIndex = mid + 1
			}
			continue
		}

		if target >= arr[mid] && target <= arr[highIndex] {
			lowIndex = mid + 1
			continue
		}
		highIndex = mid - 1
	}

	return -1
}
