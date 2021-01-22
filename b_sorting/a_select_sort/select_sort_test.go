package a_select_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func TestSelectSort(t *testing.T) {
	sortedArr := SelectSort(z_data.ArrDataList)
	t.Log(sortedArr)
}
