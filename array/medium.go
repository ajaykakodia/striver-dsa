package array

import "sort"

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
