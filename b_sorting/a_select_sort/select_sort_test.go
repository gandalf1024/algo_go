package a_select_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func TestSelectSort(t *testing.T) {
	for _, arr := range z_data.ArrDataList {
		sortedArr := SelectSort(arr)
		t.Log(sortedArr)
	}
}
