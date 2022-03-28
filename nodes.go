package gotypes

type Node struct {
	Val  any
	Prev *Node
	Next *Node
}

type NodeTree struct {
	Val   any
	Left  *Node
	Right *Node
}
