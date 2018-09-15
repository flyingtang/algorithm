package main

import "fmt"

type Student struct {
	Name string
}

type Node struct {
	data interface{}
	next *Node
}

// func (s *Student) set(name string) {
// 	s.Name = name
// }

func (head *Node) Append(stu Student) {

	if head.data == nil {
		head.data = stu
	} else {
		tHead := head
		for tHead.next != nil {
			tHead = tHead.next
		}
		temp := new(Node)
		temp.data = stu
		tHead.next = temp
	}
}

func (head *Node) travel() {
	tHead := head
	for ; tHead != nil; tHead = tHead.next {
		fmt.Println(tHead.data)
	}
}

func main() {

	stu := new(Node)
	// stu.set("xixaogang")
	stu.Append(Student{"xiaogang1"})
	stu.Append(Student{"xiaogang2"})
	stu.Append(Student{"xiaogang3"})
	stu.Append(Student{"xiaogang4"})
	fmt.Println("stydent", stu)
	stu.travel()
}
