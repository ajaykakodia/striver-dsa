package binarysearch

import "fmt"

func BS() {
	arr := []int{3, 4, 6, 7, 9, 12, 16, 17}
	fmt.Printf("Element %d is exists in array: %v - %t\n", 6, arr, binarySearchArray(arr, 6))
	fmt.Printf("Element %d is exists in array: %v - %t\n", 13, arr, binarySearchArray(arr, 13))
	arr = []int{-1, 0, 3, 5, 9, 12}
	fmt.Printf("Element %d is exists in array: %v - %t\n", 9, arr, binarySearchArray(arr, 9))
	arr = []int{3, 4, 6, 7, 9, 12, 16, 17}
	fmt.Printf("By recursion, Element %d is exists in array: %v - %t \n", 6, arr, binarySearchArray(arr, 6))
	arr = []int{-1, 0, 3, 5, 9, 12}
	fmt.Printf("By recursion, Element %d is exists in array: %v - %t \n", 9, arr, binarySearchArray(arr, 9))

	arr = []int{3, 4, 6, 7, 9, 12, 16, 17}
	fmt.Printf("Lower bound of Element %d in array: %v is %d \n", 10, arr, lowerBoundOfElement(arr, 10))

	arr = []int{1, 2, 2, 3}
	fmt.Printf("Upper bound of Element %d in array: %v is %d \n", 2, arr, upperBoundOfElement(arr, 2))
	arr = []int{5, 8, 9, 12, 14, 17, 19, 21}
	fmt.Printf("Insert Position of Element %d in array: %v is %d \n", 15, arr, searchInsertPosition(arr, 15))

	arr = []int{1, 2, 4, 7}
	fmt.Printf("Insert Position of Element %d in array: %v is %d \n", 10, arr, searchInsertPosition(arr, 10))

	arr = []int{3, 4, 4, 7, 8, 10}
	f, c := floorNCeil(arr, 5)
	fmt.Printf("Floor and Ceil of Element %d in array: %v is %d, %d \n", 5, arr, f, c)

	f, c = floorNCeil(arr, 8)
	fmt.Printf("Floor and Ceil of Element %d in array: %v is %d, %d \n", 8, arr, f, c)

	arr = []int{3, 4, 13, 13, 13, 20, 40}
	fmt.Printf("Last Occurrence of Element %d in array: %v is %d \n", 13, arr, lastOccurrenceOfKey(arr, 13))

	fmt.Printf("Total count of Element %d in array: %v is %d \n", 13, arr, countOccurrenceOfKey(arr, 13))

	arr = []int{2, 2, 3, 3, 3, 3, 4}
	fmt.Printf("Total count of Element %d in array: %v is %d \n", 3, arr, countOccurrenceOfKey(arr, 3))

	arr = []int{4, 5, 6, 7, 0, 1, 2, 3}
	fmt.Printf("Element %d is existed in rotated array: %v at %d index \n", 0, arr, searchElementInRotatedArray(arr, 0))

	arr = []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Printf("Element %d is existed in rotated array: %v at %d index \n", 3, arr, searchElementInRotatedArray(arr, 3))

	arr = []int{4, 5, 6, 7, 8, 0, 1, 2, 3}
	fmt.Printf("Element %d is existed in rotated array: %v at %d index \n", 6, arr, searchElementInRotatedArray(arr, 6))

	arr = []int{7, 8, 1, 2, 3, 3, 3, 4, 5, 6}
	fmt.Printf("Element %d is existed in rotated array: %v - %t \n", 3, arr, searchElementInArrayWithDuplicateValue(arr, 3))
}
