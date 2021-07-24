package main

import "fmt"

type Node struct{
	Value int
	Next *Node
}

type Stack struct{
	stack *Node
	Size int
}

func Create() Stack{
	stack := Stack{nil,0}
	return stack
}

func (s *Stack) Push(v int) bool{
	if s.stack == nil{
		s.stack = &Node{v,nil}
		s.Size = 1
		return true
	}
	n := &Node{v,nil}

	n.Next = s.stack
	s.stack = n
	s.Size++
	return true
}

func (s *Stack) Pop() (int,bool){
	if s.Size == 0{
		return 0, false
	}
	if s.Size == 1 {
		temp := s.stack
		s.stack = nil
		s.Size--
		return temp.Value,true
	}
	temp := s.stack
	s.stack = s.stack.Next
	s.Size--
	return temp.Value,true
}

func (s Stack) Traverse(){
	if s.Size == 0{
		fmt.Println("Empty stack")
		return
	}
	for s.stack != nil{
		fmt.Printf("%d -> ",s.stack.Value)
		s.stack = s.stack.Next
	}
	fmt.Println()
}

func main(){
	stack := Create()
	stack.Push(10)
	fmt.Println(stack.Size)
	stack.Push(11)
	stack.Push(12)
	stack.Push(15)
	fmt.Println(stack.Size)
	stack.Push(13)
	fmt.Println(stack.Size)
	stack.Traverse()
	stack.Pop()
	stack.Traverse()
	stack.Pop()
	stack.Traverse()
	stack.Pop()
	stack.Traverse()
	stack.Pop()
	stack.Traverse()
}