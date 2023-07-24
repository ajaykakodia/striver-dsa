package main

import "fmt"

func sorting() {
	arr := []int{45, 65, 34, 23, 56, 33, 4, 2, 1}
	fmt.Printf("Array before selection sorting: %v\n", arr)
	selectionSort(arr)
	fmt.Printf("Array after selection sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 33, 4, 2, 1}

	fmt.Printf("Array before bubble sorting: %v\n", arr)
	bubbleSorting(arr)
	fmt.Printf("Array after bubble sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 33, 4, 2, 1}

	fmt.Printf("Array before recursive bubble sorting: %v\n", arr)
	bubbleSortR(arr)
	fmt.Printf("Array after recursive bubble sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 31, 4, 2, 1}

	fmt.Printf("Array before insertion sorting: %v\n", arr)
	insertionSort(arr)
	fmt.Printf("Array after insertion sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 31, 4, 2, 1}

	fmt.Printf("Array before recursive insertion sorting: %v\n", arr)
	insertionSortR(arr)
	fmt.Printf("Array after recursive insertion sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 31, 4, 2, 1}

	fmt.Printf("Array before merge sorting: %v\n", arr)
	fmt.Printf("Array after merge sorting: %v\n", mergeSorting(arr))

	arr = []int{45, 65, 34, 23, 56, 31, 4, 2, 1}

	fmt.Printf("Array before quick sorting: %v\n", arr)
	quickSorting(arr)
	fmt.Printf("Array after quick sorting: %v\n", arr)
}

func selectionSort(arr []int) {
	minIndex := 0
	l := len(arr)
	for i := 0; i < l-1; i++ {
		minIndex = i
		for j := i + 1; j < l; j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}
}

func bubbleSorting(arr []int) {
	l := len(arr)
	isSwapHappened := false
	for i := 0; i < l; i++ {
		isSwapHappened = false
		for j := 0; j < l-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapHappened = true
			}
		}
		if !isSwapHappened {
			break
		}
	}
}

// recursive bubble sort
func bubbleSortR(arr []int) {
	if len(arr) <= 1 {
		return
	}
	// doing sorting from last array
	bubbleSortR(arr[1:])
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}
}

func insertionSort(arr []int) {
	l := len(arr)
	var elem int

	for i := 1; i < l; i++ {
		elem = arr[i]
		insertPosition := -1
		for j := i - 1; j >= 0; j-- {
			if arr[j] > elem {
				arr[j+1] = arr[j]
				insertPosition = j
				continue
			}
			break
		}
		if insertPosition != -1 {
			arr[insertPosition] = elem
		}
	}
}

// recursive insertion sort
func insertionSortR(arr []int) {
	l := len(arr)
	if l <= 1 {
		return
	}
	// doing sorting from last array
	insertionSortR(arr[1:])
	ele, insertPosition := arr[l-1], -1

	for i := l - 2; i >= 0; i-- {
		if arr[i] > ele {
			arr[i+1] = arr[i]
			insertPosition = i
			continue
		}
		break
	}
	if insertPosition != -1 {
		arr[insertPosition] = ele
	}
}

func quickSorting(arr []int) {
	l := len(arr)

	if l <= 1 {
		return
	}

	pivotIndex := getPivotIndex(arr)
	quickSorting(arr[:pivotIndex])
	quickSorting(arr[pivotIndex+1:])
}

func mergeSorting(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	mid := len(arr) / 2
	arr1 := mergeSorting(arr[0:mid])
	arr2 := mergeSorting(arr[mid:])
	return mergeTwoSortedArray(arr1, arr2)
}

func mergeTwoSortedArray(arr1, arr2 []int) []int {
	arr := []int{}

	l1, l2 := len(arr1), len(arr2)
	i, j := 0, 0

	for i < l1 && j < l2 {
		if arr1[i] < arr2[j] {
			arr = append(arr, arr1[i])
			i++
		} else {
			arr = append(arr, arr2[j])
			j++
		}
	}

	if i < l1 {
		arr = append(arr, arr1[i:]...)
	}

	if j < l2 {
		arr = append(arr, arr2[j:]...)
	}

	return arr
}

func getPivotIndex(arr []int) int {
	l := len(arr)
	if l <= 1 {
		return 0
	}

	ele := arr[0]
	index := 0

	for i := 1; i < l; i++ {
		if arr[i] < ele {
			index = index + 1
		}
	}

	arr[index], arr[0] = arr[0], arr[index]

	i, j := 0, l-1

	for i < j {
		for arr[i] < ele {
			i++
		}
		for arr[j] > ele {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	return index
}
