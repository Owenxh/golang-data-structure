package tree

type Ordered interface {
	string | int | int8 | int16 | int32 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | uintptr
}

// UF Union Find
type UF interface {
	// GetSize returns elements count of UF
	GetSize() int
	// IsConnected indicates whether the element p & q is connected
	IsConnected(p, q int) bool
	// UnionElements connects element p & q
	UnionElements(p, q int)
}
