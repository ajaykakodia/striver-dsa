package main

import "fmt"

/*
Notes to follow along:

	R - if you find R in function name in last, it's a recursive solution
	BF - its for Brute Force solution

*/

func main() {
	fmt.Println("Here we start with Striver DSA SDE sheet codes")

	// step 0:
	//pattern()

	// step 1.1:
	//basicMath()

	// step 1.2
	//basicRecursion()

	// step 1.3
	//basicHashing()

	// step 2
	arr := []int{45, 65, 34, 23, 56, 33, 4, 2, 1}
	fmt.Printf("Array before selection sorting: %v\n", arr)
	selectionSort(arr)
	fmt.Printf("Array after selection sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 33, 4, 2, 1}

	fmt.Printf("Array before bubble sorting: %v\n", arr)
	bubbleSorting(arr)
	fmt.Printf("Array after bubble sorting: %v\n", arr)

	arr = []int{45, 65, 34, 23, 56, 31, 4, 2, 1}

	fmt.Printf("Array before insertion sorting: %v\n", arr)
	insertionSort(arr)
	fmt.Printf("Array after insertion sorting: %v\n", arr)

}
