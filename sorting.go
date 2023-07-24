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
