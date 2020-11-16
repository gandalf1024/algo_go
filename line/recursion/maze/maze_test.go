package maze

import "testing"

//迷宫算法(寻找最优路径)
func Test_maze(t *testing.T) {
	// 先创建一个二维数组，模拟迷宫
	// 地图
	var ma = [8][7]int{}

	// 使用1 表示墙
	// 上下全部置为1
	for i := 0; i < 7; i++ {
		ma[0][i] = 1
		ma[7][i] = 1
	}

	// 左右全部置为1
	for i := 0; i < 8; i++ {
		ma[i][0] = 1
		ma[i][6] = 1
	}
	//设置挡板, 1 表示
	ma[3][1] = 1
	ma[3][2] = 1
	//		ma[1][2] = 1;
	//		ma[2][2] = 1;

	// 输出地图
	println("地图的情况")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			print(ma[i][j], "  ")
		}
		println()
	}

	//使用递归回溯给小球找路
	//setWay(&ma, 1, 1);
	setWay2(&ma, 1, 1)

	//输出新的地图, 小球走过，并标识过的递归
	println("小球走过，并标识过的 地图的情况")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			print(ma[i][j], " ")
		}
		println()
	}

}
