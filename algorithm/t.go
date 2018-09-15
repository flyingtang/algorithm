package main

import (
	"fmt"
	"strconv"
)

func main() {

	s := "10"
	fmt.Printf("%T \n", s)
	s1, _ := strconv.Atoi(s)

	var s2 uint = uint(s1)
	fmt.Printf("%d %d \n", s1, s2)

	fmt.Printf("%T \n", string(s1))
}
