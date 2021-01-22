package z_data

import (
	"math/rand"
	"time"
)

var ArrDataList []int

func init() {
	rand.Seed(time.Now().UnixNano())
	ArrDataList = make([]int, 0, 200)
	for i := 0; i < 25; i++ {
		ArrDataList = append(ArrDataList, rand.Intn(999))
	}
}
