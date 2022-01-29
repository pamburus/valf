package valf

// ---

// ArrayReader is an abstract array reader.
type ArrayReader interface {
	ValfReadArray() []Value
}
