package algorithm

import (
	"math"
)

/**
回文问题：例如121 12321 是回文
@Prama {int} num 传入的数字
@Return {bool} res 是回文 true, 不是 false
*/

func Huiwen(num int) (res bool) {
	// 判断多少位
	k := 1     // 判断标准
	digit := 0 //  位数
	for k > 0 {
		k = num - int(math.Pow10(digit))
		digit++
	}
	digit-- //  正确的位数
	sum := 0
	n := num
	// 交换高低位， 判断是否是相等
	for i := 0; i < digit; i++ {
		temp := n % 10                           // 个位
		sum += temp * int(math.Pow10(digit-1-i)) // 放进高位
		n /= 10                                  // 去除个位
	}
	if sum == num {
		return true
	}
	return false
}
