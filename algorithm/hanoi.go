package algorithm

import "fmt"

var count int = 0

func Hanoi(n int, a, b, c byte) {
	if n == 1 {
		count++
		fmt.Printf("第%d次移动 : \t 圆盘从%c 移动到 %c \n", count, a, c)
	} else {
		Hanoi(n-1, a, c, b)
		count++
		fmt.Printf("第%d次移动 : \t 圆盘从%c 移动到 %c \n ", count, a, c)
		Hanoi(n-1, b, a, c)
	}
}
