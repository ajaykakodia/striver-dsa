package array

import "fmt"

func ArrayProblems() {
	fmt.Println("Here we go with array problems ..... ")

	// arr := []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	// fmt.Printf("largest element in arr: %v is %d.\n", arr, largestElementInArray(arr))

	// arr = []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	// sl, ss := secondLargestAnd2ndSmallestElement(arr)
	// fmt.Printf("2nd largest element is %d and 2nd smallest element is %d in arr: %v.\n", sl, ss, arr)
	// arr = []int{1}
	// sl, ss = secondLargestAnd2ndSmallestElement(arr)
	// fmt.Printf("2nd largest element is %d and 2nd smallest element is %d in arr: %v.\n", sl, ss, arr)

	// arr = []int{8, 10, 5, 7, 9, 2, 20, 17, 25, 12, 13}
	// fmt.Printf("Array: %v is sorted = %v\n", arr, isArraySorted(arr))
	// arr = []int{2, 4, 6, 7, 8, 10}
	// fmt.Printf("Array: %v is sorted = %v\n", arr, isArraySorted(arr))

	// arr = []int{5, 7, 7, 9, 9, 9, 9, 11, 15, 17, 17, 18, 19, 25, 25}

	// fmt.Printf("Array: %v after duplicate removal = %v\n", []int{5, 7, 7, 9, 9, 9, 9, 11, 15, 17, 17, 18, 19, 25, 25}, removeDuplicateForSortedArray(arr))

	// arr1 := []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// fmt.Printf("Array: %v after duplicate removal = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, removeDuplicateForSortedArray(arr1))

	// arr = []int{1, 1}

	// fmt.Printf("Array: %v after duplicate removal = %v\n", []int{1, 1}, removeDuplicateForSortedArray(arr))

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// leftRotateArrayByOneBF(arr)
	// fmt.Printf("Array: %v after left rotate by one = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, arr)

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// fmt.Printf("Array: %v after left rotate by one = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, leftRotateArrayByOneGolang(arr))

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// rotateArrayByKElement(arr, 3, "left")
	// fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "left", arr)

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// rotateArrayByKElement(arr, 3, "right")
	// fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "right", arr)

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "left", rotateArrayByKElementGoLang(arr, 3, "left"))

	// fmt.Printf("Array: %v after rotate by K element %d to %s position = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, 3, "right", rotateArrayByKElementGoLang(arr, 3, "right"))

	// arr = []int{1, 0, 2, 3, 0, 4, 0, 1}
	// moveZeroToEndInArray(arr)
	// fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{1, 0, 2, 3, 0, 4, 0, 1}, arr)

	// arr = []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	// moveZeroToEndInArray(arr)
	// fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{5, 7, 9, 11, 15, 17, 18, 19, 25}, arr)
	// arr = []int{0, 0, 0, 0, 0}
	// moveZeroToEndInArray(arr)
	// fmt.Printf("Array: %v after moving all zeros in last = %v\n", []int{0, 0, 0, 0, 0}, arr)

	arr := []int{5, 7, 9, 11, 15, 17, 18, 19, 25}
	fmt.Printf("In Array: %v, %d present at position: %d\n", arr, 19, linearSearch(arr, 19))
	fmt.Printf("In Array: %v, %d present at position: %d\n", arr, 27, linearSearch(arr, 27))

	arr1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr2 := []int{2, 3, 4, 4, 5, 11, 12}
	fmt.Printf("Union of Array, %v and %v is: %v\n", arr1, arr2, unionOfTwoArrayBF(arr1, arr2))
	fmt.Printf("Union of Array with Space Solution, %v and %v is: %v\n", arr1, arr2, unionOfTwoArrayWithSpace(arr1, arr2))
	fmt.Printf("Union of Two Sorted Array, %v and %v is: %v\n", arr1, arr2, unionOfTwoSortedArray(arr1, arr2))
}
