package array

import "math"

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