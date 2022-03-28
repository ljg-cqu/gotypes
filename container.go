package gotypes

type Container interface {
	Count() int
	Empty() bool
	Clear()
}
