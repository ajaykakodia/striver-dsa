package array

import "fmt"

func ArrayProblems() {
	fmt.Println("Here we go with array problems ..... ")

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

}
