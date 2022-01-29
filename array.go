package valf

// ---

// ValueArray accepts ArrayItemVisitor.
type ValueArray interface {
	ArrayItemCount() int
	AcceptArrayItemVisitor(ArrayItemVisitor)
}

// ArrayItemVisitor visits array item.
type ArrayItemVisitor interface {
	VisitArrayItem(index int, value Value)
}
