package valf

import (
	"fmt"
	"time"
)

// visitor defines the interface that allows to visit a value
// and do some action depending on its real type.
type visitor interface {
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
	VisitStringer(fmt.Stringer)
	VisitFormattable(string, interface{})
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
	VisitArray(ArrayReader)
	VisitObject(ObjectReader)
}

// ignoringVisitor is an implementation of Visitor interface which does nothing.
type ignoringVisitor struct{}

// VisitNone does nothing.
func (v ignoringVisitor) VisitNone() {}

// VisitAny does nothing.
func (v ignoringVisitor) VisitAny(interface{}) {}

// VisitBool does nothing.
func (v ignoringVisitor) VisitBool(bool) {}

// VisitInt does nothing.
func (v ignoringVisitor) VisitInt(int) {}

// VisitInt8 does nothing.
func (v ignoringVisitor) VisitInt8(int8) {}

// VisitInt16 does nothing.
func (v ignoringVisitor) VisitInt16(int16) {}

// VisitInt32 does nothing.
func (v ignoringVisitor) VisitInt32(int32) {}

// VisitInt64 does nothing.
func (v ignoringVisitor) VisitInt64(int64) {}

// VisitUint does nothing.
func (v ignoringVisitor) VisitUint(uint) {}

// VisitUint8 does nothing.
func (v ignoringVisitor) VisitUint8(uint8) {}

// VisitUint16 does nothing.
func (v ignoringVisitor) VisitUint16(uint16) {}

// VisitUint32 does nothing.
func (v ignoringVisitor) VisitUint32(uint32) {}

// VisitUint64 does nothing.
func (v ignoringVisitor) VisitUint64(uint64) {}

// VisitFloat32 does nothing.
func (v ignoringVisitor) VisitFloat32(float32) {}

// VisitFloat64 does nothing.
func (v ignoringVisitor) VisitFloat64(float64) {}

// VisitDuration does nothing.
func (v ignoringVisitor) VisitDuration(time.Duration) {}

// VisitError does nothing.
func (v ignoringVisitor) VisitError(error) {}

// VisitTime does nothing.
func (v ignoringVisitor) VisitTime(time.Time) {}

// VisitString does nothing.
func (v ignoringVisitor) VisitString(string) {}

// VisitStrings does nothing.
func (v ignoringVisitor) VisitStrings([]string) {}

// VisitBytes does nothing.
func (v ignoringVisitor) VisitBytes([]byte) {}

// VisitBools does nothing.
func (v ignoringVisitor) VisitBools([]bool) {}

// VisitInts does nothing.
func (v ignoringVisitor) VisitInts([]int) {}

// VisitInts8 does nothing.
func (v ignoringVisitor) VisitInts8([]int8) {}

// VisitInts16 does nothing.
func (v ignoringVisitor) VisitInts16([]int16) {}

// VisitInts32 does nothing.
func (v ignoringVisitor) VisitInts32([]int32) {}

// VisitInts64 does nothing.
func (v ignoringVisitor) VisitInts64([]int64) {}

// VisitUints does nothing.
func (v ignoringVisitor) VisitUints([]uint) {}

// VisitUints8 does nothing.
func (v ignoringVisitor) VisitUints8([]uint8) {}

// VisitUints16 does nothing.
func (v ignoringVisitor) VisitUints16([]uint16) {}

// VisitUints32 does nothing.
func (v ignoringVisitor) VisitUints32([]uint32) {}

// VisitUints64 does nothing.
func (v ignoringVisitor) VisitUints64([]uint64) {}

// VisitFloats32 does nothing.
func (v ignoringVisitor) VisitFloats32([]float32) {}

// VisitFloats64 does nothing.
func (v ignoringVisitor) VisitFloats64([]float64) {}

// VisitDurations does nothing.
func (v ignoringVisitor) VisitDurations([]time.Duration) {}

// VisitArray does nothing.
func (v ignoringVisitor) VisitArray(ArrayReader) {}

// VisitObject does nothing.
func (v ignoringVisitor) VisitObject(ObjectReader) {}

// VisitStringer does nothing.
func (v ignoringVisitor) VisitStringer(fmt.Stringer) {}

// VisitFormattable does nothing.
func (v ignoringVisitor) VisitFormattable(string, interface{}) {}
