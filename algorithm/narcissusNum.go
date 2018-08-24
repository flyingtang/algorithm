package algorithm

import (
	"fmt"
	"math"
)

/**
水仙花问题
例如 4位数字
9474 = 9^4 + 4^4 + 7^4 + 4^4
@Param {int} 数字位数
*/

func NarcissusNum(n int) {
	start := int(math.Pow10(n - 1))
	end := int(math.Pow10(n) - 1)
	var sum int
	for i := start; i <= end; i++ {
		// fmt.Println("i:", i)
		sum = 0
		s := i
		for j := 0; j < n; j++ {
			temp := s % 10
			sum += int(math.Pow(float64(temp), float64(n)))
			s /= 10
		}
		if i == sum {
			fmt.Printf("%d \n", i)
		}
	}
}
