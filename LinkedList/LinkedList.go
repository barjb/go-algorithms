package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct{
	Root *Node
	Middle *Node
	Length int
}

func (l *LinkedList) Insert(n *Node, v int) int{
	if l.Root == nil{
		l.Root = &Node{v,nil}
		l.Length = 1
		return 0
	}
	if n.Value == v{
		fmt.Printf("Node with value %d exists\n",v)
		return -1
	}
	if n.Next == nil{
		n.Next = &Node{v,nil}
		l.Length++
		return 1
	}
	return l.Insert(n.Next,v)
}

func (l LinkedList) Traverse(n *Node){
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

func (l *LinkedList) Exists(n *Node,v int) bool{
	if n.Value == v{
		return true
	}
	if n.Next == nil{
		return false
	}
	return l.Exists(n.Next,v)
}

func (l LinkedList) Size() int {
	if l.Root == nil {
		fmt.Println("List is empty")
		return 0
	}
	var i = 0
	for l.Root != nil {
		l.Root = l.Root.Next
		i++
	}
	return i
}
func main() {
	list := new(LinkedList)
	list.Insert(list.Root,10)
	list.Insert(list.Root,20)
	list.Insert(list.Root,30)
	list.Insert(list.Root,40)
	list.Insert(list.Root,15)
	list.Insert(list.Root,11)

	list.Traverse(list.Root)
	
	fmt.Println(list.Exists(list.Root,10))
	fmt.Println(list.Exists(list.Root,11))
	fmt.Println(list.Exists(list.Root,12))

	fmt.Println(list.Size())
}
