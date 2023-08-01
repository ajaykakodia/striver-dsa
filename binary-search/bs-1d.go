package binarysearch

func binarySearchArray(arr []int, k int) bool {
	l := len(arr)
	start, end, mid := 0, l-1, 0
	for start <= end {
		mid = (start + end) / 2
		if arr[mid] == k {
			return true
		}
		if arr[mid] > k {
			end = mid - 1
			continue
		}
		start = mid + 1
	}

	return false
}

func binarySearchArrayR(arr []int, k int) bool {
	if len(arr) == 0 {
		return false
	}

	start, end := 0, len(arr)-1
	mid := (start + end) / 2

	if arr[mid] == k {
		return true
	}
	if arr[mid] > k {
		return binarySearchArrayR(arr[:mid], k)
	}
	return binarySearchArrayR(arr[mid+1:], k)
}
