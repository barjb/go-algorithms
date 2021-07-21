package main

import (
	"fmt"
	"math/rand"
	"time"
)

type BinaryTree struct {
	Value int
	Number int
	Left  *BinaryTree
	Right *BinaryTree
}

func Traverse(t *BinaryTree) {
	if t == nil {
		return
	}
	fmt.Printf("%d:%d ",t.Number, t.Value)
	Traverse(t.Left)
	Traverse(t.Right)
}

func InsertNode(t *BinaryTree, v int,n int) *BinaryTree {
	if t == nil {
		return &BinaryTree{Left: nil, Right: nil, Value: v, Number: n}
	}
	if v < t.Value {
		t.Left = InsertNode(t.Left, v,2*n)
		return t
	}
	t.Right = InsertNode(t.Right, v,2*n+1)
	return t
}

func CreateBinaryTree(n int) *BinaryTree {
	var t *BinaryTree
	rand.Seed(time.Now().Unix())
	for i:=0 ; i < n; i++{
		v := rand.Intn(n)
		t = InsertNode(t, v, 1)
	}
	return t
}



func main() {
	BinaryTree := CreateBinaryTree(8)
	Traverse(BinaryTree)
}
