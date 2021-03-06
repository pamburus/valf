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
	if !v.bits.Const() {
		switch v.bits.Type() {
		case TypeNone:
		case TypeAny:
			snapshotAny(v)
		case TypeBytes:
			snapshotBytes(v)
		case TypeBools:
			snapshotBools(v)
		case TypeInts:
			snapshotInts(v)
		case TypeInts8:
			snapshotInts8(v)
		case TypeInts16:
			snapshotInts16(v)
		case TypeInts32:
			snapshotInts32(v)
		case TypeInts64:
			snapshotInts64(v)
		case TypeUints:
			snapshotUints(v)
		case TypeUints8:
			snapshotUints8(v)
		case TypeUints16:
			snapshotUints16(v)
		case TypeUints32:
			snapshotUints32(v)
		case TypeUints64:
			snapshotUints64(v)
		case TypeFloats32:
			snapshotFloats32(v)
		case TypeFloats64:
			snapshotFloats64(v)
		case TypeDurations:
			snapshotDurations(v)
		case TypeStrings:
			snapshotStrings(v)
		case TypeArray:
			snapshotArray(v)
		case TypeObject:
			snapshotObject(v)
		case TypeStringer:
			snapshotStringer(v)
		case TypeFormatter:
			snapshotFormatter(v)

		default:
			panic(fmt.Errorf("snapf: internal error: unhandled value type: %v", v.bits.Type()))
		}
	}
}

func snapshotBytes(v *Value) {
	cc := make([]byte, len(v.vBytes))
	copy(cc, v.vBytes)
	v.vBytes = cc
	v.bits |= bitsConst
}

func snapshotBools(v *Value) {
	s := *(*[]bool)(unsafe.Pointer(&v.vBytes))
	cc := make([]bool, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotInts(v *Value) {
	s := *(*[]int)(unsafe.Pointer(&v.vBytes))
	cc := make([]int, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotInts8(v *Value) {
	s := *(*[]int8)(unsafe.Pointer(&v.vBytes))
	cc := make([]int8, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotInts16(v *Value) {
	s := *(*[]int16)(unsafe.Pointer(&v.vBytes))
	cc := make([]int16, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotInts32(v *Value) {
	s := *(*[]int32)(unsafe.Pointer(&v.vBytes))
	cc := make([]int32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotInts64(v *Value) {
	s := *(*[]int64)(unsafe.Pointer(&v.vBytes))
	cc := make([]int64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotUints(v *Value) {
	s := *(*[]uint)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotUints8(v *Value) {
	s := *(*[]uint8)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint8, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotUints16(v *Value) {
	s := *(*[]uint16)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint16, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotUints32(v *Value) {
	s := *(*[]uint32)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotUints64(v *Value) {
	s := *(*[]uint64)(unsafe.Pointer(&v.vBytes))
	cc := make([]uint64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotFloats32(v *Value) {
	s := *(*[]float32)(unsafe.Pointer(&v.vBytes))
	cc := make([]float32, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotFloats64(v *Value) {
	s := *(*[]float64)(unsafe.Pointer(&v.vBytes))
	cc := make([]float64, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotDurations(v *Value) {
	s := *(*[]time.Duration)(unsafe.Pointer(&v.vBytes))
	cc := make([]time.Duration, len(s))
	copy(cc, s)
	v.vBytes = *(*[]byte)(unsafe.Pointer(&cc))
	v.bits |= bitsConst
}

func snapshotStringer(v *Value) {
	v.vString = v.vAny.(fmt.Stringer).String()
	v.vAny = nil
	v.bits = bits(TypeString) | bitsConst
}

func snapshotFormatter(v *Value) {
	v.vString = fmt.Sprintf(v.vString, v.vAny)
	v.vAny = nil
	v.bits = bits(TypeString) | bitsConst
}

func snapshotStrings(v *Value) {
	s := v.vAny.([]string)
	cc := make([]string, len(s))
	copy(cc, s)
	v.vAny = cc
	v.bits |= bitsConst
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
	v.bits |= bitsConst
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
	v.bits |= bitsConst
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
