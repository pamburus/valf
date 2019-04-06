package valf

import (
	"time"
)

// Visitor defines the interface that allows to visit a value
// and do some action depending on its real type.
type Visitor interface {
	VisitNone()
	VisitAny(interface{})
	VisitBool(bool)
	VisitInt(int)
	VisitInt8(int8)
	VisitInt16(int16)
	VisitInt32(int32)
	VisitInt64(int64)
	VisitUint(uint)
	VisitUint8(uint8)
	VisitUint16(uint16)
	VisitUint32(uint32)
	VisitUint64(uint64)
	VisitFloat32(float32)
	VisitFloat64(float64)
	VisitDuration(time.Duration)
	VisitError(error)
	VisitTime(time.Time)
	VisitString(string)
	VisitStrings([]string)
	VisitBytes([]byte)
	VisitBools([]bool)
	VisitInts([]int)
	VisitInts8([]int8)
	VisitInts16([]int16)
	VisitInts32([]int32)
	VisitInts64([]int64)
	VisitUints([]uint)
	VisitUints8([]uint8)
	VisitUints16([]uint16)
	VisitUints32([]uint32)
	VisitUints64([]uint64)
	VisitFloats32([]float32)
	VisitFloats64([]float64)
	VisitDurations([]time.Duration)
	VisitArray(ValueArray)
	VisitObject(ValueObject)
}

// ValueArray accepts ArrayItemVisitor.
type ValueArray interface {
	ArrayItemCount() int
	AcceptArrayItemVisitor(ArrayItemVisitor)
}

// ArrayItemVisitor visits array item.
type ArrayItemVisitor interface {
	VisitArrayItem(index int, value Value)
}

// ValueObject accepts ObjectFieldVisitor.
type ValueObject interface {
	ObjectFieldCount() int
	AcceptObjectFieldVisitor(ObjectFieldVisitor)
}

// ObjectFieldVisitor visits object field.
type ObjectFieldVisitor interface {
	VisitObjectField(key string, value Value)
}

// IgnoringVisitor is an implementation of Visitor interface which does nothing.
type IgnoringVisitor struct{}

// VisitNone does nothing.
func (v IgnoringVisitor) VisitNone() {}

// VisitAny does nothing.
func (v IgnoringVisitor) VisitAny(interface{}) {}

// VisitBool does nothing.
func (v IgnoringVisitor) VisitBool(bool) {}

// VisitInt does nothing.
func (v IgnoringVisitor) VisitInt(int) {}

// VisitInt8 does nothing.
func (v IgnoringVisitor) VisitInt8(int8) {}

// VisitInt16 does nothing.
func (v IgnoringVisitor) VisitInt16(int16) {}

// VisitInt32 does nothing.
func (v IgnoringVisitor) VisitInt32(int32) {}

// VisitInt64 does nothing.
func (v IgnoringVisitor) VisitInt64(int64) {}

// VisitUint does nothing.
func (v IgnoringVisitor) VisitUint(uint) {}

// VisitUint8 does nothing.
func (v IgnoringVisitor) VisitUint8(uint8) {}

// VisitUint16 does nothing.
func (v IgnoringVisitor) VisitUint16(uint16) {}

// VisitUint32 does nothing.
func (v IgnoringVisitor) VisitUint32(uint32) {}

// VisitUint64 does nothing.
func (v IgnoringVisitor) VisitUint64(uint64) {}

// VisitFloat32 does nothing.
func (v IgnoringVisitor) VisitFloat32(float32) {}

// VisitFloat64 does nothing.
func (v IgnoringVisitor) VisitFloat64(float64) {}

// VisitDuration does nothing.
func (v IgnoringVisitor) VisitDuration(time.Duration) {}

// VisitError does nothing.
func (v IgnoringVisitor) VisitError(error) {}

// VisitTime does nothing.
func (v IgnoringVisitor) VisitTime(time.Time) {}

// VisitString does nothing.
func (v IgnoringVisitor) VisitString(string) {}

// VisitStrings does nothing.
func (v IgnoringVisitor) VisitStrings([]string) {}

// VisitBytes does nothing.
func (v IgnoringVisitor) VisitBytes([]byte) {}

// VisitBools does nothing.
func (v IgnoringVisitor) VisitBools([]bool) {}

// VisitInts does nothing.
func (v IgnoringVisitor) VisitInts([]int) {}

// VisitInts8 does nothing.
func (v IgnoringVisitor) VisitInts8([]int8) {}

// VisitInts16 does nothing.
func (v IgnoringVisitor) VisitInts16([]int16) {}

// VisitInts32 does nothing.
func (v IgnoringVisitor) VisitInts32([]int32) {}

// VisitInts64 does nothing.
func (v IgnoringVisitor) VisitInts64([]int64) {}

// VisitUints does nothing.
func (v IgnoringVisitor) VisitUints([]uint) {}

// VisitUints8 does nothing.
func (v IgnoringVisitor) VisitUints8([]uint8) {}

// VisitUints16 does nothing.
func (v IgnoringVisitor) VisitUints16([]uint16) {}

// VisitUints32 does nothing.
func (v IgnoringVisitor) VisitUints32([]uint32) {}

// VisitUints64 does nothing.
func (v IgnoringVisitor) VisitUints64([]uint64) {}

// VisitFloats32 does nothing.
func (v IgnoringVisitor) VisitFloats32([]float32) {}

// VisitFloats64 does nothing.
func (v IgnoringVisitor) VisitFloats64([]float64) {}

// VisitDurations does nothing.
func (v IgnoringVisitor) VisitDurations([]time.Duration) {}

// VisitArray does nothing.
func (v IgnoringVisitor) VisitArray(ValueArray) {}

// VisitObject does nothing.
func (v IgnoringVisitor) VisitObject(ValueObject) {}
