package three_hill_sort

import (
	"algo_go/two_sort/test_data"
	"testing"
)

func Test_HillSort(t *testing.T) {
	for _, arr := range test_data.ArrDataList {
		sortedArr := HillSort(arr)
		t.Log(sortedArr)
	}
}
