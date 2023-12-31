package main

import "fmt"

func basicHashing() {
	countFrequency([]int{10, 5, 10, 15, 10, 5})
	fmt.Println(highAndLowFrequency([]int{10, 5, 10, 15, 10, 5}))
	fmt.Println(highAndLowFrequency([]int{2, 2, 3, 4, 4, 2}))
}

func countFrequency(arr []int) {
	m := make(map[int]int)

	for _, v := range arr {
		m[v] = m[v] + 1
	}

	for k, v := range m {
		fmt.Printf("%d occurs %d times in the array\n", k, v)
	}
}

func highAndLowFrequency(arr []int) (int, int) {
	low, high, ln, hn := len(arr), 0, 0, 0
	m := make(map[int]int)
	for _, v := range arr {
		m[v] = m[v] + 1
	}

	for k, v := range m {
		if v > high {
			high = v
			hn = k
		}
		if v < low {
			low = v
			ln = k
		}
	}

	return hn, ln
}
