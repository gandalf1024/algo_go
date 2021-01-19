package two_insertion_sort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && arr[j] < arr[j-1]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func InsertionBinarySort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		if j <= 0 {
			continue
		}

		index := binarySearch(arr, arr[j], j)
		arr[j], arr[index] = arr[index], arr[j]
	}
	return arr
}

func binarySearch(arr []int, target int, high int) int {
	low := 0
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
