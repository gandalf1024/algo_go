package binarysearchnorecursion

//二分查找
func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mind := (low + high) / 2
		if arr[mind] == target {
			return mind
		} else if arr[mind] > target {
			high = mind - 1
		} else {
			low = mind + 1
		}
	}
	return -1
}
