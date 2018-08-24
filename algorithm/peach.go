package algorithm

import "fmt"

/**
猴子吃桃问题
每天吃一半+多一个 十天后剩下一个
@Prama {int} day 天数
*/

func Peach(days int) (res int) {
	if days == 1 {
		res = 1
		return
	}
	fmt.Println("days = ", days)
	res = (Peach(days-1)+1)*2
	fmt.Println("res= ", res)
	return
}
