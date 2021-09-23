package main

type Node struct {
	Key Value
	Value Value
	Pre *Node
	Next *Node
}

type LRUCache struct {
	Limit int
	HashMap map[Value]*Node
}

type Value interface {}


