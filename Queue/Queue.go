package main

import "fmt"

type Node struct{
	Value int
	Next *Node
}

type Queue struct{
	queue *Node
	Size int
}

func Create() Queue{
	queue := Queue{nil,0}
	return queue
}

func (q *Queue) Push(v int) bool{
	if q.queue == nil{
		q.queue = &Node{v,nil}
		q.Size = 1
		return true
	}
	n := &Node{v,nil}
	n.Next = q.queue
	q.queue = n
	q.Size++
	return true
}

func (q *Queue) Pop() (int,bool){
	if q.Size == 0{
		return 0, false
	}
	if q.Size == 1 {
		temp := q.queue
		q.queue = nil
		q.Size--
		return temp.Value,true
	}
	temp := q.queue
	for temp.Next.Next != nil{
		temp = temp.Next
	}
	v := temp.Next.Value
	temp.Next = nil
	q.Size--
	return v,true
}

func (q Queue) Traverse(){
	if q.Size == 0{
		fmt.Println("Empty queue")
		return
	}
	temp := q.queue
	for temp != nil{
		fmt.Printf("%d -> ",temp.Value)
		temp = temp.Next
	}
	fmt.Println()
}

func main(){
	queue := Create()
	queue.Push(10)
	fmt.Println(queue.Size)
	queue.Push(11)
	queue.Push(12)
	queue.Push(15)
	fmt.Println(queue.Size)
	queue.Push(13)
	fmt.Println(queue.Size)
	queue.Traverse()
	queue.Pop()
	queue.Traverse()
	queue.Pop()
	queue.Traverse()
	queue.Pop()
	queue.Traverse()
	queue.Pop()
	queue.Traverse()
}