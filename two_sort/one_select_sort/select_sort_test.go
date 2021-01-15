package one_select_sort

import (
	"algo_go/two_sort/test_data"
	"testing"
)

func TestSelectSort(t *testing.T) {
	for _, arr := range test_data.ArrDataList {
		sortedArr := SelectSort(arr)
		t.Log(sortedArr)
	}
}
