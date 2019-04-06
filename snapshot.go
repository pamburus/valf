package valf

import (
	"errors"
	"fmt"
	"time"
	"unsafe"
)

// Snapshotter is the interface that allows to do a custom snapshotting strategy of a value.
type Snapshotter interface {
	TakeSnapshot() interface{}
}

// Snapshot changes the v so that it can be safely stored for a long with guarantee that it won't be modified.
// The data of the value are copied if it should be copied to achieve that guarantee.
func Snapshot(v *Value) {
	if v.t&valueTypeConst == 0 {
		switch v.t & valueTypeMask {
		case valueTypeNone:
		case valueTypeAny:
			snapshotAny(v)
		case valueTypeBytes:
			snapshotBytes(v)
		case valueTypeBools:
			snapshotBools(v)
		case valueTypeInts:
			snapshotInts(v)
		case valueTypeInts8:
			snapshotInts8(v)
		case valueTypeInts16:
			snapshotInts16(v)
		case valueTypeInts32:
			snapshotInts32(v)
		case valueTypeInts64:
			snapshotInts64(v)
		case valueTypeUints:
			snapshotUints(v)
		case valueTypeUints8:
			snapshotUints8(v)
		case valueTypeUints16:
			snapshotUints16(v)
		case valueTypeUints32:
			snapshotUints32(v)
		case valueTypeUints64:
			snapshotUints64(v)
		case valueTypeFloats32:
			snapshotFloats32(v)
		case valueTypeFloats64:
			snapshotFloats64(v)
		case valueTypeDurations:
			snapshotDurations(v)
		case valueTypeStrings:
			snapshotStrings(v)
		case valueTypeArray:
			snapshotArray(v)
		case valueTypeObject:
			snapshotObject(v)
		case valueTypeStringer:
			snapshotStringer(v)
		case valueTypeFormatter:
			snapshotFormatter(v)

		default:
			panic(fmt.Errorf("snapf: internal error: unhandled value type: %v", v.t))
		}
	}
}

func snapshotBytes(v *Value) {
	cc := make([]byte, len(v.vBytes))
	copy(cc, v.vBytes)
	v.vBytes = cc
	v.t |= valueTypeConst
}

func snapshotBools(v *Value) {
	s := *(*[]bool)(unsafe.Pointer(&v.vBytes))
	cc := make([]bool, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotInts(v *Value) {
	s := *(*[]int)(unsafe.Pointer(&v.vBytes))
	cc := make([]int, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotInts8(v *Value) {
	s := *(*[]int8)(unsafe.Pointer(&v.vBytes))
	cc := make([]int8, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotInts16(v *Value) {
	s := *(*[]int16)(unsafe.Pointer(&v.vBytes))
	cc := make([]int16, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotInts32(v *Value) {
	s := *(*[]int32)(unsafe.Pointer(&v.vBytes))
	cc := make([]int32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotInts64(v *Value) {
	s := *(*[]int64)(unsafe.Pointer(&v.vBytes))
	cc := make([]int64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotUints(v *Value) {
	s := *(*[]uint)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotUints8(v *Value) {
	s := *(*[]uint8)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint8, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotUints16(v *Value) {
	s := *(*[]uint16)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint16, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotUints32(v *Value) {
	s := *(*[]uint32)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotUints64(v *Value) {
	s := *(*[]uint64)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotFloats32(v *Value) {
	s := *(*[]float32)(unsafe.Pointer(&v.vBytes))
	cc := make([]float32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotFloats64(v *Value) {
	s := *(*[]float64)(unsafe.Pointer(&v.vBytes))
	cc := make([]float64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotDurations(v *Value) {
	s := *(*[]time.Duration)(unsafe.Pointer(&v.vBytes))
	cc := make([]time.Duration, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.t |= valueTypeConst
}

func snapshotStringer(v *Value) {
	v.vString = v.vAny.(fmt.Stringer).String()
	v.vAny = nil
	v.t = valueTypeString | valueTypeConst
}

func snapshotFormatter(v *Value) {
	v.vString = fmt.Sprintf(v.vString, v.vAny)
	v.vAny = nil
	v.t = valueTypeString | valueTypeConst
}

func snapshotStrings(v *Value) {
	s := v.vAny.([]string)
	cc := make([]string, len(s))
	copy(cc, s)
	v.vAny = cc
	v.t |= valueTypeConst
}

func snapshotAny(v *Value) {
	snapshotter, ok := v.vAny.(Snapshotter)
	if !ok {
		panic(errors.New("snapf: cannot snapshot value with type Any since it does not implement Snapshotter interface"))
	}

	*v = ConstAny(snapshotter.TakeSnapshot())
}

func snapshotArray(v *Value) {
	a := v.vAny.(ValueArray)
	s := arraySnapshotter{arraySnapshot{make([]Value, a.ArrayItemCount())}}
	a.AcceptArrayItemVisitor(&s)
	v.vAny = s.snapshot
	v.t |= valueTypeConst
}

type arraySnapshotter struct {
	snapshot arraySnapshot
}

func (s *arraySnapshotter) VisitArrayItem(index int, value Value) {
	Snapshot(&value)
	s.snapshot.items[index] = value
}

type arraySnapshot struct {
	items []Value
}

func (s arraySnapshot) ArrayItemCount() int {
	return len(s.items)
}

func (s arraySnapshot) AcceptArrayItemVisitor(visitor ArrayItemVisitor) {
	for i, v := range s.items {
		visitor.VisitArrayItem(i, v)
	}
}

type objectField struct {
	Name  string
	Value Value
}

func snapshotObject(v *Value) {
	o := v.vAny.(ValueObject)
	s := objectSnapshotter{objectSnapshot{make([]objectField, 0, o.ObjectFieldCount())}}
	o.AcceptObjectFieldVisitor(&s)
	v.vAny = s.snapshot
	v.t |= valueTypeConst
}

type objectSnapshotter struct {
	snapshot objectSnapshot
}

func (s *objectSnapshotter) VisitObjectField(name string, value Value) {
	Snapshot(&value)
	s.snapshot.fields = append(s.snapshot.fields, objectField{name, value})
}

type objectSnapshot struct {
	fields []objectField
}

func (s objectSnapshot) ObjectFieldCount() int {
	return len(s.fields)
}

func (s objectSnapshot) AcceptObjectFieldVisitor(visitor ObjectFieldVisitor) {
	for _, field := range s.fields {
		visitor.VisitObjectField(field.Name, field.Value)
	}
}
