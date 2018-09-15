package main

import "fmt"

type Student struct {
	key  string
	Name string
	age  uint
}

type Node struct {
	data     interface{}
	nextNode *Node
}

func (node *Node) append(student Student) {

	if node.data == nil {
		node.data = student
	} else {
		tHead := node
		newNode := new(Node)
		newNode.data = student
		newNode.nextNode = nil
		for tHead.nextNode != nil {
			tHead = tHead.nextNode
		}
		tHead.nextNode = newNode
	}

}
func (head *Node) travel() {
	tHead := head
	for ; tHead != nil; tHead = tHead.nextNode {
		if data, ok := tHead.data.(Student); ok {
			fmt.Printf("%s, %s, %d \n", data.key, data.Name, data.age)
		}

	}
}

func main() {

	student := Student{"1", "xiaogang1", 19}
	head := new(Node)
	head.append(student)
	head.append(student)
	head.travel()
	// fmt.Println("link", head)
}
