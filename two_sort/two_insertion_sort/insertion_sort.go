package two_insertion_sort

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && less(arr[j], arr[j-1]); j-- {
			exch(arr, j, j-1)
		}
	}
	return arr
}

func less(v, w int) bool {
	return v-w < -1
}

func exch(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
