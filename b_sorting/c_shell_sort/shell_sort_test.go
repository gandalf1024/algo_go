package c_shell_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func Test_HillSort(t *testing.T) {
	for _, arr := range z_data.ArrDataList {
		sortedArr := ShellSort(arr)
		t.Log(sortedArr)
	}
}
