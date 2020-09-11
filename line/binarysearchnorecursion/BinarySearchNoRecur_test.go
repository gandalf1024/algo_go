package binarysearchnorecursion

import "testing"

func Test_s(t *testing.T) {
	//测试
	arr := []int{1, 3, 8, 10, 11, 67, 100}
	index := binarySearch(arr, 100)
	println("index=", index)

}
