package gotypes

type Node struct {
	Val  any
	Next *Node
}

type NodePN struct {
	Val  any
	Prev *Node
	Next *Node
}

type NodeLR struct {
	Val   any
	Left  *Node
	Right *Node
}
