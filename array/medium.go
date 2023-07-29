package array

import (
	"math"
	"sort"
)

/*
Two Sum : Check if a pair with given sum exists in Array
Problem Statement: Given an array of integers arr[] and an integer target.

1st variant: Return YES if there exist two numbers such that their sum is equal to the target. Otherwise, return NO.

2nd variant: Return indices of the two numbers such that their sum is equal to the target. Otherwise, we will return {-1, -1}.

Note: You are not allowed to use the same element twice. Example: If the target is equal to 6 and num[1] = 3, then nums[1] + nums[1] = target is not a solution.

Example 1:
Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 14
Result: YES (for 1st variant)

	[1, 3] (for 2nd variant)

Explanation: arr[1] + arr[3] = 14. So, the answer is “YES” for the first variant and [1, 3] for 2nd variant.

Example 2:
Input Format: N = 5, arr[] = {2,6,5,8,11}, target = 15
Result: NO (for 1st variant)

	[-1, -1] (for 2nd variant)

Explanation: There exist no such two numbers whose sum is equal to the target.
*/
func twoSumProblemV1(arr []int, target int) bool {
	m := make(map[int]int)
	l := len(arr)

	for i := 0; i < l; i++ {
		m[arr[i]] = i
	}

	for i := 0; i < l; i++ {
		if v, ok := m[target-arr[i]]; ok && v != i {
			return true
		}
	}

	return false
}

func twoSumProblemV1TP(arr []int, target int) bool {

	sort.Ints(arr)
	sum, i, j := 0, 0, len(arr)-1

	for i < j {
		sum = arr[i] + arr[j]
		if sum == target {
			return true
		}
		if sum > target {
			j--
			continue
		}
		i++
	}

	return false
}

func twoSumProblemV2(arr []int, target int) []int {
	m := make(map[int]int)
	l := len(arr)

	for i := 0; i < l; i++ {
		m[arr[i]] = i
	}

	for i := 0; i < l; i++ {
		if v, ok := m[target-arr[i]]; ok && v != i {
			return []int{i, v}
		}
	}

	return []int{-1, -1}
}

/*
Problem Statement: Given an array consisting of only 0s, 1s, and 2s. Write a program to in-place sort the array without using inbuilt sort functions. ( Expected: Single pass-O(N) and constant space)

Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Input: nums = [2,0,1]
Output: [0,1,2]

Input: nums = [0]
Output: [0]
*/
func sortArrayOf012s(arr []int) {
	low, mid, high := 0, 0, len(arr)-1

	for mid <= high {
		if arr[mid] == 0 {
			arr[mid], arr[low] = arr[low], arr[mid]
			mid++
			low++
			continue
		}
		if arr[mid] == 1 {
			mid++
			continue
		}
		arr[mid], arr[high] = arr[high], arr[mid]
		high--
	}
}

/*
	Find the Majority Element that occurs more than N/2 times

Problem Statement: Given an array of N integers, write a program to return an element that occurs more than N/2 times in the given array. You may consider that such an element always exists in the array.

Example 1:
Input Format: N = 3, nums[] = {3,2,3}
Result: 3
Explanation: When we just count the occurrences of each number and compare with half of the size of the array, you will get 3 for the above solution.

Example 2:
Input Format:  N = 7, nums[] = {2,2,1,1,1,2,2}

Result: 2

Explanation: After counting the number of times each element appears and comparing it with half of array size, we get 2 as result.

Example 3:
Input Format:  N = 10, nums[] = {4,4,2,4,3,4,4,3,2,4}

Result: 4
*/
func majorityElementInArray(arr []int) int {
	count := 0
	var ele *int

	for _, v := range arr {
		if ele == nil {
			ele = &v
			count++
			continue
		}

		if *ele == v {
			count++
			continue
		}
		if count == 0 {
			*ele = v
			count++
			continue
		}
		count--
	}
	return *ele
}

/*
Kadane’s Algorithm : Maximum Subarray Sum in an Array
Problem Statement: Given an integer array arr, find the contiguous subarray (containing at least one number) which
has the largest sum and returns its sum and prints the subarray.
Example 1:

Input: arr = [-2,1,-3,4,-1,2,1,-5,4]

Output: 6

Explanation: [4,-1,2,1] has the largest sum = 6.

Examples 2:

Input: arr = [1]

Output: 1

Explanation: Array has only one element and which is giving positive sum of 1.
*/
func maxSubArraySum(arr []int) int {
	sum, maxSum := 0, math.MinInt

	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if sum > maxSum {
			maxSum = sum
		}
		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}

/*
Stock Buy And Sell
Problem Statement: You are given an array of prices where prices[i] is the price of a given stock on an ith day.

You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

Example 1:

Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and
sell on day 5 (price = 6), profit = 6-1 = 5.

Note: That buying on day 2 and selling on day 1
is not allowed because you must buy before
you sell.

Example 2:

Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are
done and the max profit = 0.
*/
func stockBuyNSellProblem(arr []int) int {
	maxProfit, bDay := 0, arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] < bDay {
			bDay = arr[i]
			continue
		}
		if arr[i] > bDay && arr[i]-bDay > maxProfit {
			maxProfit = arr[i] - bDay
		}
	}

	return maxProfit
}

/*
Rearrange Array Elements by Sign
Variety-1

Problem Statement:

There’s an array ‘A’ of size ‘N’ with an equal number of positive and negative elements. Without altering the relative order of positive and negative elements, you must return an array of alternately positive and negative values.

Note: Start the array with positive elements.

Example 1:

Input:
arr[] = {1,2,-4,-5}, N = 4
Output:
1 -4 2 -5

Explanation:

Positive elements = 1,2
Negative elements = -4,-5
To maintain relative ordering, 1 must occur before 2, and -4 must occur before -5.

Example 2:
Input:
arr[] = {1,2,-3,-1,-2,-3}, N = 6
Output:
1 -3 2 -1 3 -2
Explanation:

Positive elements = 1,2,3
Negative elements = -3,-1,-2
To maintain relative ordering, 1 must occur before 2, and 2 must occur before 3.
Also, -3 should come before -1, and -1 should come before -2.
*/
func rearrangeArrayBySign(arr []int) {
	l, a, b := len(arr), 0, 0
	arr1, arr2 := make([]int, l/2), make([]int, l/2)

	for i := 0; i < l; i++ {
		if arr[i] < 0 {
			arr2[b] = arr[i]
			b++
			continue
		}

		arr1[a] = arr[i]
		a++
	}
	a = 0
	for i := 0; i < l/2; i++ {
		arr[a] = arr1[i]
		arr[a+1] = arr2[i]
		a = a + 2
	}
}

/*
next_permutation : find next lexicographically greater permutation
Problem Statement: Given an array Arr[] of integers, rearrange the numbers of the given array into the lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange to the lowest possible order (i.e., sorted in ascending order).

Example 1 :

Input format: Arr[] = {1,3,2}
Output: Arr[] = {2,1,3}
Explanation: All permutations of {1,2,3} are {{1,2,3} , {1,3,2}, {2,13} , {2,3,1} , {3,1,2} , {3,2,1}}. So, the next permutation just after {1,3,2} is {2,1,3}.
Example 2:

Input format: Arr[] = {3,2,1}
Output: Arr[] = {1,2,3}
Explanation: As we see all permutations of {1,2,3}, we find {3,2,1} at the last position. So, we have to return the topmost permutation.
*/
func nextPermutation(nums []int) {
	l := len(nums)
	for i := l - 2; i >= 0; i-- {
		// checking for last number taht is no in sort order
		if nums[i] >= nums[i+1] {
			continue
		}

		j := l - 1
		// find number just greater than our unsorted number and replace it
		for j > i {
			if nums[j] > nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
				break
			}
			j--
		}
		//  now every thing is sorted, reverse all element prior to unsorted one to get our result
		for k := 0; k < (l-1-i)/2; k++ {
			nums[i+k+1], nums[l-k-1] = nums[l-k-1], nums[i+k+1]
		}
		return
	}
	// handle edge case for last permutation combination in order
	sort.Ints(nums)
}

/*
Leaders in an Array
Problem Statement: Given an array, print all the elements which are leaders. A Leader is an element that is greater than all of the elements on its right side in the array.

Example 1:
Input: [4, 7, 1, 0]
Output: 7 1 0

Explanation: Rightmost element is always a leader. 7 and 1 are greater than the elements in their right side.

Example 2:
Input: [10, 22, 12, 3, 0, 6]
Output: 22 12 6

Explanation: 6 is a leader. In addition to that, 12 is greater than all the elements in its right side (3, 0, 6), also 22 is greater than 12, 3, 0, 6.
*/
func leaderInArray(arr []int) []int {
	leader := []int{}
	l := len(arr)
	if l == 0 {
		return leader
	}

	lastLeader := arr[l-1]
	leader = append(leader, lastLeader)

	for i := l - 2; i >= 0; i-- {
		if arr[i] > lastLeader {
			leader = append(leader, arr[i])
			lastLeader = arr[i]
		}
	}

	return leader
}

/*
Longest Consecutive Sequence in an Array
Problem Statement: You are given an array of ‘N’ integers. You need to find the length of the longest sequence which contains the consecutive elements.

Example 1:

Input: [100, 200, 1, 3, 2, 4]
Output: 4

Explanation: The longest consecutive subsequence is 1, 2, 3, and 4.

Input: [3, 8, 5, 7, 6]
Output: 4

Explanation: The longest consecutive subsequence is 5, 6, 7, and 8.
*/
func longestConsecutiveSequence(arr []int) int {
	ans, count := 0, 0
	sort.Ints(arr)

	for i := 1; i < len(arr); i++ {
		if arr[i-1]+1 == arr[i] {
			count++
			continue
		}
		if count > ans {
			ans = count
		}
		count = 0
	}

	if count > ans {
		ans = count
	}

	return ans + 1
}

func longestConsecutiveSequenceOpt(arr []int) int {
	ans, count := 0, 0
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = true
	}

	for k := range m {
		count = 1
		i := k - 1
		for m[i] {
			count++
			delete(m, i)
			i--
		}
		i = k + 1
		for m[i] {
			count++
			delete(m, i)
			i++
		}
		if count > ans {
			ans = count
		}
	}

	return ans
}

/*
Set Matrix Zeroes

Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.

You must do it in place.

Example 1:

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]
Output: [[1,0,1],[0,0,0],[1,0,1]]
Example 2:

Input: matrix = [[0,1,2,0],[3,4,5,2],[1,3,1,5]]
Output: [[0,0,0,0],[0,4,5,0],[0,3,1,0]]
*/
func setMatrixZero(matrix [][]int) {
	zeroMatrix := [][]int{}

	for i, row := range matrix {
		for j, col := range row {
			if col == 0 {
				zeroMatrix = append(zeroMatrix, []int{i, j})
			}
		}
	}
	row, col := len(matrix), len(matrix[0])
	for _, v := range zeroMatrix {
		for i := 0; i < col; i++ {
			matrix[v[0]][i] = 0
		}
		for j := 0; j < row; j++ {
			matrix[j][v[1]] = 0
		}
	}
}

/*
Rotate Image by 90 degree
Problem Statement: Given a matrix, your task is to rotate the matrix 90 degrees clockwise.

Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]

Output: [[7,4,1],[8,5,2],[9,6,3]]

Explanation: Rotate the matrix simply by 90 degree clockwise and return the matrix.

Example 2:

Input: [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]

Output:[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

Explanation: Rotate the matrix simply by 90 degree clockwise and return the matrix
*/
func rotateMatrixBy90(matrix [][]int) {
	currentEle, ele, nextRow, nextCol, l := 0, 0, 0, 0, len(matrix)

	for i := 0; i < l-1; i++ {
		for j := i; j < l-1-i; j++ {
			currentEle = matrix[i][j]
			nextRow = j
			nextCol = l - 1 - i
			ele = matrix[nextRow][nextCol]
			matrix[nextRow][nextCol] = currentEle

			currentEle = ele
			nextRow, nextCol = nextCol, l-1-nextRow
			ele = matrix[nextRow][nextCol]
			matrix[nextRow][nextCol] = currentEle

			currentEle = ele
			nextRow, nextCol = nextCol, l-1-nextRow
			ele = matrix[nextRow][nextCol]
			matrix[nextRow][nextCol] = currentEle

			currentEle = ele
			nextRow, nextCol = nextCol, l-1-nextRow
			matrix[nextRow][nextCol] = currentEle
		}
	}
}

func rotateMatrixBy90Opt(matrix [][]int) {
	l := len(matrix)
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	printMatrix(matrix)
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			matrix[i][j], matrix[i][l-1-j] = matrix[i][l-1-j], matrix[i][j]
		}
	}
}
