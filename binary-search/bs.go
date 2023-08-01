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
}
