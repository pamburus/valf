package valf

type Kind byte

const (
	KindNone Kind = iota
	KindAny
	KindBool
	KindInt
	KindFloat
	KindString
	KindError
	KindTime
	KindDuration
	KindBytes
	KindArray
	KindObject
)
