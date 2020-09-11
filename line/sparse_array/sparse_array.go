package sparse_array

//稀疏数组
import (
	"fmt"
)

func spaese_array() {
	//原始数组
	arr := [20][20]int{}
	arr[8][9] = 1
	arr[10][6] = 2
	arr[17][3] = 1

	//稀疏数组定义
	spaeseArr := make([][3]int, 0)

	//把源数组有用数据添加 --> 稀疏数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] != 0 {
				p := [3]int{i, j, arr[i][j]}
				spaeseArr = append(spaeseArr, p)
			}
		}
	}

	for _, arr := range spaeseArr {
		fmt.Println(arr)
	}
}
