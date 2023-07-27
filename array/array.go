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
}
