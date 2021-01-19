package three_hill_sort

func HillSort(arr []int) []int {
	N := len(arr)
	h := 1
	for h < N/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < N; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
			}
		}
		h = h / 3
	}

	return arr
}
