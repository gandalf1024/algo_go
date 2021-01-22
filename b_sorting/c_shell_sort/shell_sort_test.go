package c_shell_sort

import (
	"algo_go/b_sorting/z_data"
	"testing"
)

func Test_ShellSort(t *testing.T) {
	sortedArr := ShellSort(z_data.ArrDataList)
	t.Log(sortedArr)
}
