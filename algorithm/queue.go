package main

import (
	"fmt"
	"strconv"
	"strings"
)

type DATA struct {
	name string
	age  int
}

const MAXQUEUELEN = 10

type Queue struct {
	data [MAXQUEUELEN]DATA
	head uint
	tail uint
}

func initQueue() (q *Queue) {
	q = new(Queue)
	return
}

func (q *Queue) insertQ(data DATA) {
	if q.tail >= MAXQUEUELEN {
		fmt.Println("队列已满")
	}
	q.data[q.tail] = data
	q.tail++
}

func (q *Queue) outQ() {
	fmt.Printf("%d %d \n", q.head, q.tail)
	if q.head >= q.tail {
		fmt.Println("队列一空！！！")
		return
	}
	fmt.Println(q.data[q.head])
	q.head++
}

func main() {
	// const length = 10
	const counter = 4
	q := initQueue()
	// fmt.Println("q= ", q)
	// q.insertQ(DATA{"aa", 12})

	var str string
	for i := 0; i < counter; i++ {
		fmt.Println("请输入 aa 12")
		if _, err := fmt.Scan(&str); err == nil {
			// fmt.Println("n = ", n)
			s := strings.Split(str, "|")
			// fmt.Println("s = ", s)
			if len(s) == 1 {
				continue
			}
			age, err := strconv.Atoi(s[1])
			if err != nil {
				fmt.Println("转换失败 ", s[1])
				continue
			}
			q.insertQ(DATA{s[0], age})
		} else {
			fmt.Println("err = ", err)
		}
	}

	// 出队列
	for i := 0; i < counter+1; i++ {
		fmt.Println("请输入0，出队列")
		fmt.Scan(&str)
		if str == "0" {
			q.outQ()
		}
	}

	// fmt.Println(q)
}
