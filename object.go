package valf

// ---

// ValueObject accepts ObjectFieldVisitor.
type ValueObject interface {
	ObjectFieldCount() int
	AcceptObjectFieldVisitor(ObjectFieldVisitor)
}

// ObjectFieldVisitor visits object field.
type ObjectFieldVisitor interface {
	VisitObjectField(key string, value Value)
}
