package b_insertion_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	sortedArr := InsertionSort(z_data.ArrDataList)
	t.Log(sortedArr)
}
