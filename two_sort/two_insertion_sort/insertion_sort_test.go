package two_insertion_sort

import (
	"algo_go/two_sort/test_data"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, arr := range test_data.ArrDataList {
		sortedArr := InsertionSort(arr)
		t.Log(sortedArr)
	}
}
