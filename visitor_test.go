package valf

import (
	"testing"
	"time"
)

func TestIgnoringVisitor(t *testing.T) {
	var v IgnoringVisitor
	v.VisitNone()
	v.VisitAny(nil)
	v.VisitBool(false)
	v.VisitInt(1)
	v.VisitInt8(2)
	v.VisitInt16(3)
	v.VisitInt32(4)
	v.VisitInt64(5)
	v.VisitUint(6)
	v.VisitUint8(7)
	v.VisitUint16(8)
	v.VisitUint32(9)
	v.VisitUint64(10)
	v.VisitFloat32(11.0)
	v.VisitFloat64(11.1)
	v.VisitDuration(time.Second)
	v.VisitError(nil)
	v.VisitTime(time.Now())
	v.VisitString("s")
	v.VisitStrings([]string{"s1", "s2"})
	v.VisitBytes([]byte{1, 2, 3})
	v.VisitBools([]bool{false, true})
	v.VisitInts([]int{0, 1, 2})
	v.VisitInts8([]int8{3, 4, 5})
	v.VisitInts16([]int16{6, 7, 8})
	v.VisitInts32([]int32{9, 10, 11})
	v.VisitInts64([]int64{12, 13, 14})
	v.VisitUints([]uint{15, 16, 17})
	v.VisitUints8([]uint8{19, 20, 21})
	v.VisitUints16([]uint16{22, 23, 24})
	v.VisitUints32([]uint32{25, 26, 27})
	v.VisitUints64([]uint64{28, 29, 30})
	v.VisitFloats32([]float32{12.1, 12.2})
	v.VisitFloats64([]float64{13.1, 13.2})
	v.VisitDurations([]time.Duration{time.Second, time.Hour})
	v.VisitArray(nil)
	v.VisitObject(nil)
}
