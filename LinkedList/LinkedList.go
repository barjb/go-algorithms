package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func InsertNode(t *Node, v int) int {
	if t == nil {
		fmt.Println("in nil")
		t = &Node{v, nil}
		fmt.Printf("%d -> %d\n", t.Value,t.Next)
		return 0
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return 1
	}
	return InsertNode(t.Next, v)
}

func traverse(t *Node) {
	if t == nil {
		fmt.Println("List is empty")
		return
	}
	for ;t != nil; {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func Exists(t *Node, v int) bool {
	if t.Value == v{
		return true
	}
	if t.Next == nil {
		return false
	}
	return Exists(t.Next, v)
}

func Size(t *Node) int {
	if t == nil {
		fmt.Println("List is empty")
		return 0
	}
	var i = 0
	for ;t != nil; {
		t = t.Next
		i++
	}
	return i
}
func main() {
	root := &Node{Value:1, Next: nil}
	InsertNode(root, -11)
	InsertNode(root, 2)
	fmt.Println(Exists(root,2))
	fmt.Println(Exists(root,10))
	fmt.Println(Size(root))
	traverse(root)
}
