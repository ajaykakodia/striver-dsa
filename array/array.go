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

	arr = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Leader in array: %v is: %v\n", arr, leaderInArray(arr))

	arr = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Longest Consecutive Sequence in array: %v is: %d\n", []int{7, 1, 5, 3, 6, 4}, longestConsecutiveSequence(arr))

	arr = []int{100, 200, 1, 3, 2, 4}
	fmt.Printf("Longest Consecutive Sequence in array: %v is: %d\n", []int{100, 200, 1, 3, 2, 4}, longestConsecutiveSequence(arr))

	arr = []int{7, 1, 5, 3, 6, 4}
	fmt.Printf("Longest Consecutive Sequence in array: %v is: %d\n", arr, longestConsecutiveSequenceOpt(arr))

	arr = []int{100, 200, 1, 3, 2, 4}
	fmt.Printf("Longest Consecutive Sequence in array: %v is: %d\n", arr, longestConsecutiveSequenceOpt(arr))

	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setMatrixZero(matrix)
	fmt.Printf("matrix before zero: %v and after zero replacement : %v \n", [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, matrix)

	matrix = [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setMatrixZero(matrix)
	fmt.Printf("matrix before zero: %v and after zero replacement : %v \n", [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, matrix)

	matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("matrix before rotate: ")
	printMatrix(matrix)
	rotateMatrixBy90(matrix)
	fmt.Println("matrix after rotate: ")
	printMatrix(matrix)

	matrix = [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	fmt.Println("matrix before rotate: ")
	printMatrix(matrix)
	rotateMatrixBy90Opt(matrix)
	fmt.Println("matrix after rotate: ")
	printMatrix(matrix)
	matrix = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	spiralBindingTraversalArray(matrix)

	arr = []int{1, 2, 3, -3, 1, 1, 1, 4, 2, -3}
	fmt.Printf("Count Subarray sum Equals %d of array: %v is: %d\n", 6, arr, countSubArraySumToK(arr, 6))

	arr = []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("Count Subarray sum Equals %d of array: %v is: %d\n", 6, arr, countSubArraySumToK(arr, 0))
}

func printMatrix(matrix [][]int) {
	fmt.Println()
	for _, v := range matrix {
		for _, ele := range v {
			fmt.Print(ele, "\t")
		}
		fmt.Println()
	}
	fmt.Println()
}
