package algorithm

import "fmt"

/**
经典的约瑟夫问题, 41 个人，数到3的自杀，问约瑟夫和他的朋友怎样才能存活下来
@Prama {int} nums 共有多少人
@Prama {int} alive 要活多少人
*/
func F(nums, alive int) {
	arr := make([]int, nums) // 初始化一个人数为nums的切片
	const killNum = 3        // 自杀间隔
	                 // 循环计数
	counter, pos ,i := 1, -1, 0    // 已经自杀人数, 当前的位置
	for counter <= nums {
		for {
			pos = (pos + 1) % nums
			if arr[pos] == 0 {
				i++
			}
			if i == killNum {
				i = 0
				break
			}
		}
		arr[pos] = counter
		counter++
	}

	// 要活人数编号
	fmt.Printf("这 %d 需要存活的人初始位置应排在以下序号:\n", alive)
	alives := nums - alive
	for i := 0; i < nums; i++ {
		if arr[i] > alives {
			fmt.Printf("初始编号：%d, 约瑟夫环编号:%d \n", i+1, arr[i])
		}
	}

}
