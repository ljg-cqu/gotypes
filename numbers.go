package gotypes

type Number interface {
	Int | Uint | Float
}

type IntFloat interface {
	Int | Float
}

type UintFloat interface {
	Uint | Float
}

type IntUint interface {
	Int | Uint
}

type Float interface {
	float32 | float64
}

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}
