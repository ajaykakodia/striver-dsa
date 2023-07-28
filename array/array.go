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

	arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Printf("Max SubArray Sum  in array: %v is: %d\n", arr, maxSubArraySum(arr))

	arr = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Max Profit after stock selling in rates: %v is: %d\n", arr, stockBuyNSellProblem(arr))

	arr = []int{3, 1, -2, -5, 2, -4}
	rearrangeArrayBySign(arr)
	fmt.Printf("Array before arranging: %v and after arranging: %v:\n", []int{3, 1, -2, -5, 2, -4}, arr)

	arr = []int{2, 1, 5, 4, 3, 0, 0}
	nextPermutation(arr)
	fmt.Printf("Next Permutation of array: %v is: %v:\n", []int{2, 1, 5, 4, 3, 0, 0}, arr)
}
