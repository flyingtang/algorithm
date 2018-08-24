package algorithm

import "fmt"

/**
百钱百鸡
公 5 母 3 小鸡 1/3
*/
func Checken(m, n int) {
	var x, y, z int
	for ; x <= m; x++ {
		for y = 0; y <= m; y++ {
			z = m - x - y
			if z >= 0 && z%3 == 0 && 5*x+3*y+z/3 == n {
				fmt.Printf("x = %d y = %d z = %d \n", x, y, z)
			}
		}
	}
}
