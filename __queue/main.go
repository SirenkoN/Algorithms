package main

import (
	"container/list"
	"fmt"
)

func Push(elem interface{}, queue *list.List) {
	for temp := queue.Front(); ; temp = temp.Next() {
		if temp == nil {
			queue.PushBack(elem)
			break
		}
	}
}

func Pop(queue *list.List) interface{} {
	elem := queue.Front()
	queue.Remove(queue.Front())
	return elem
}

func printQueue(queue *list.List) {
	for temp := queue.Front(); temp != nil; temp = temp.Next() {
		fmt.Print(temp.Value)
	}
}

func ReverseList(l *list.List) *list.List {
	reversedList := list.New()
	for temp := l.Front(); temp != nil; temp = temp.Next() {
		reversedList.PushFront(temp.Value)
	}

	for temp1 := reversedList.Front(); temp1 != nil; temp1 = temp1.Next() {
		fmt.Print(temp1.Value)
	}

	return reversedList
}

func main() {

	queue := list.New()

	Push(1, queue)

	Push(2, queue)

	Push(3, queue)

	printQueue(queue) // 123
	//fmt.Print(len(queue))
	ReverseList(queue)

	//Pop(queue)
	fmt.Println("")
	//printQueue(queue) // в ту же строку: 23

}
