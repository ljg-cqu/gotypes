package gotypes

type Container interface {
	Count() int
	Empty() bool
	Clear()
}

type LinkedListContainer interface {
	PushFront(val any)
	PushBack(val any)
	Back() any
	Front() any
	Reverse()
	Container
}
