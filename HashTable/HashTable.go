package main

import "fmt"

type HashTable struct {
	Table map[int]*Node
	Size  int
}

type Node struct {
	Value int
	Next  *Node
}

func hashFunction(i, size int) int {
	return i % size
}

func (h *HashTable) Insert(v int) int {
	index := hashFunction(v, h.Size)
	h.Table[index] = &Node{Value: v, Next: h.Table[index]}
	return index
}

func (h HashTable) Traverse() {
	for i := range h.Table {
		if h.Table[i] != nil {
			l := h.Table[i]
			for l != nil {
				fmt.Printf("%d -> ", l.Value)
				l = l.Next
			}
			fmt.Println()
		}
	}
}

func Create(size int) HashTable{
	table := make(map[int]*Node, size)
	hash := HashTable{Table: table, Size: size}
	return hash
}

func main() {
	hashTable := Create(10)
	for i := 0; i < 120; i++ {
		hashTable.Insert(i)
	}
	hashTable.Traverse()
}
