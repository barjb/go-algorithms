package main

import "fmt"

type Node struct{
	Value int
	Previous *Node
	Next *Node
}
type DoublyLinkedList struct{
	Root *Node
	Middle *Node
	Length int
}

func (d *DoublyLinkedList) Insert(t *Node, v int) int {
	if d.Root == nil{ 
		d.Root = &Node{v,nil,nil}
		d.Length = 1
		return 0
	}
	if t.Next == nil{
		t.Next = &Node{v,t,nil}
		d.Length++
		return 1
	}
	return d.Insert(t.Next,v)
}

func (d DoublyLinkedList) Traverse(n *Node){
	if n == nil {
		fmt.Println("List is empty")
		return
	}
	for n != nil {
		fmt.Printf("%d -> ", n.Value)
		n = n.Next
	}
	fmt.Println()
}

func main(){
	list := new(DoublyLinkedList)
	list.Traverse(list.Root)

	list.Insert(list.Root,10)
	list.Insert(list.Root,20)
	list.Insert(list.Root,30)
	list.Insert(list.Root,10)
	list.Insert(list.Root,20)
	list.Traverse(list.Root)
}