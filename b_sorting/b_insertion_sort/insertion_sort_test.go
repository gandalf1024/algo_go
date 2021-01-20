package b_insertion_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, arr := range z_data.ArrDataList {
		sortedArr := InsertionSort(arr)
		t.Log(sortedArr)
	}
}
