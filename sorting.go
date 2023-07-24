package main

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
