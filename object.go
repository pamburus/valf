package valf

// ---

// ObjectReader is an abstract object reader.
type ObjectReader interface {
	ValfReadObject() []Field
}
