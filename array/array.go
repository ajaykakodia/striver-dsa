package array

import "fmt"

func ArrayProblems() {
	fmt.Println("Here we go with array problems ..... ")

	// easy problem section
	//easy()

	// from here medium problem starts

	arr := []int{2, 6, 5, 8, 11}
	fmt.Printf("Target %d exists in array %v: %t\n", 14, arr, twoSumProblemV1(arr, 14))

	arr = []int{2, 6, 5, 8, 11}
	fmt.Printf("Target %d exists in array %v: at indices %v\n", 14, arr, twoSumProblemV2(arr, 14))

	arr = []int{2, 3, 6, 5, 8, 11, 7, 15, 16}
	fmt.Printf("Target %d exists in array %v: %t\n", 15, arr, twoSumProblemV1TP(arr, 15))

	arr = []int{0, 2, 1, 2, 0, 1, 0, 2, 2, 1, 1, 0, 2, 1}
	sortArrayOf012s(arr)
	fmt.Printf("Array before sorting: %v and after sorting: %v:\n", []int{0, 2, 1, 2, 0, 1, 0, 2, 2, 1, 1, 0, 2, 1}, arr)

	arr = []int{0, 2, 1, 2, 0, 1, 0, 2, 2, 1, 1, 0, 2, 1, 2, 2, 3, 0, 2, 2}
	fmt.Printf("Majority Element in array: %v is: %d\n", arr, majorityElementInArray(arr))

}
