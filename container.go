package gotypes

type Container interface {
	Count() int
	Empty() bool
	Clear()
}

type LinkedListContainer interface {
	PushFront(val any)
	PushBack(val any)
	Reverse()
	Container
}
