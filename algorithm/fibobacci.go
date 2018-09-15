package main

import "fmt"

func fibobacci(n int64) int64 {
	if 2 == n || 1 == n {
		return 1
	} else {
		return fibobacci(n-1) + fibobacci(n-2)
	}

}

func main() {
	var n int64
	fmt.Println("请输入数字n")
	fmt.Scan(&n)
	res := fibobacci(n)
	fmt.Println("res: ", res)
}
