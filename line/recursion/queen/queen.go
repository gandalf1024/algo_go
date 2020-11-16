package queen

import (
	"fmt"
	"math"
)

var max = 8

var array []int = make([]int, max)

var count = 0

var judgeCount = 0

func check(n int) {
	if n == max { //n = 8 , 其实8个皇后就既然放好
		show()
		return
	}
	//依次放入皇后，并判断是否冲突
	for i := 0; i < max; i++ {
		//先把当前这个皇后 n , 放到该行的第1列
		array[n] = i
		//判断当放置第n个皇后到i列时，是否冲突
		if judge(n) { // 不冲突
			//接着放n+1个皇后,即开始递归
			check(n + 1) //
		}
		//如果冲突，就继续执行 array[n] = i; 即将第n个皇后，放置在本行得 后移的一个位置
	}
}

func judge(n int) bool {
	judgeCount++
	for i := 0; i < n; i++ {
		if array[i] == array[n] || math.Abs(float64(n)-float64(i)) == math.Abs(float64(array[n])-float64(array[i])) {
			return false
		}
	}
	return true
}

//写一个方法，可以将皇后摆放的位置输出
func show() {
	count++
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i], " ")
	}
	fmt.Println()
}
