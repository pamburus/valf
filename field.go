package valf

func NewField(key string, value Value) Field {
	return Field{
		key,
		value,
	}
}

// Field is a key and value pair.
type Field struct {
	key   string
	value Value
}

func (f Field) Key() string {
	return f.key
}

func (f Field) Value() Value {
	return f.value
}

// Snapshot returns a Field which can be safely stored for a long with guarantee that it won't be modified.
// The data of the field are copied if needed to achieve that guarantee.
func (f Field) Snapshot() Field {
	f.value = f.value.Snapshot()

	return f
}
