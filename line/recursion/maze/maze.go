package maze

func setWay(ma *[8][7]int, i, j int) bool {
	if ma[6][5] == 2 { // 通路已经找到ok
		return true
	} else {
		if ma[i][j] == 0 { //如果当前这个点还没有走过
			//按照策略 下->右->上->左  走
			ma[i][j] = 2            // 假定该点是可以走通.
			if setWay(ma, i+1, j) { //向下走
				return true
			} else if setWay(ma, i, j+1) { //向右走
				return true
			} else if setWay(ma, i-1, j) { //向上
				return true
			} else if setWay(ma, i, j-1) { //向左走
				return true
			} else {
				//说明该点是走不通，是死路
				ma[i][j] = 3
				return false
			}
		} else { // 如果ma[i][j] != 0 , 可能是 1， 2， 3
			return false
		}
	}
}

//修改找路的策略，改成 上->右->下->左
func setWay2(ma *[8][7]int, i, j int) bool {
	if ma[6][5] == 2 { // 通路已经找到ok
		return true
	} else {
		if ma[i][j] == 0 { //如果当前这个点还没有走过
			//按照策略 上->右->下->左
			ma[i][j] = 2             // 假定该点是可以走通.
			if setWay2(ma, i-1, j) { //向上走
				return true
			} else if setWay2(ma, i, j+1) { //向右走
				return true
			} else if setWay2(ma, i+1, j) { //向下
				return true
			} else if setWay2(ma, i, j-1) { // 向左走
				return true
			} else {
				//说明该点是走不通，是死路
				ma[i][j] = 3
				return false
			}
		} else { // 如果ma[i][j] != 0 , 可能是 1， 2， 3
			return false
		}
	}
}
