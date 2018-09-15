package main

import "fmt"

func getRes(head, foot uint) (r, c uint) {
	fmt.Println("head", head, c)
	for ; c <= head; c++ {
		r = head - c
		if r*4+c*2 == foot {
			return r, c
		}
	}
	return 0, 0
}

func main() {
	var head, foot uint

	fmt.Println("请输入头数:")
	fmt.Scan(&head)

	fmt.Println("请输入脚数:")
	fmt.Scan(&foot)
	fmt.Printf("%T", head)
	r, c := getRes(head, foot)
	if r == 0 && c == 0 {
		fmt.Println("未找到结果")
	} else {
		fmt.Printf("兔子 %d , 🐔 %d", r, c)
	}
}
