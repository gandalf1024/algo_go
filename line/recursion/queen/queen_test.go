package queen

import (
	"fmt"
	"testing"
)

//八皇后问题(回溯算法)
func Test_queen(t *testing.T) {
	check(0)
	fmt.Printf("一共有%d解法\n", count)
	fmt.Printf("一共判断冲突的次数%d次\n", judgeCount) // 1.5w
}
