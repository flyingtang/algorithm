package algorithm

import "fmt"

/**
八皇后问题， 为了避免皇后吃了皇后，在8 * 8 为棋盘上，两个皇后不能再同一横轴、纵轴、对称线上
@Parama {[]int} queens 棋盘，采用一维切片记录结果
@Parama {int} cur 当前位置
*/

func Queen(queens []int, cur int) {
	length := len(queens)
	if cur == length {
		fmt.Println(queens)
		return
	}
	for i := 0; i < length; i++ {

		queens[cur] = i
		flag := true
		for j := 0; j < cur; j++ {

			var t = i - queens[j]
			if t < 0 {
				t = -t
			}

			if queens[j] == i || t == cur-j {
				flag = false
				break
			}
		}
		if flag {
			Queen(queens, cur+1)
		}
	}
}
