package valf

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type testStringer string

func (s testStringer) String() string {
	return string(s)
}

type mockFormatter struct {
	t     *testing.T
	value string
}

func (f *mockFormatter) Format(s fmt.State, c rune) {
	require.Equal(f.t, 'v', c)
	_, ok := s.Width()
	require.Equal(f.t, false, ok)
	_, ok = s.Precision()
	require.Equal(f.t, false, ok)
	require.Equal(f.t, true, s.Flag(int('#')))

	n, err := s.Write([]byte(f.value))
	require.Equal(f.t, len(f.value), n)
	require.Equal(f.t, error(nil), err)
}

type mockArray []Value

func (a mockArray) ArrayItemCount() int {
	return len(a)
}

func (a mockArray) AcceptArrayItemVisitor(visitor ArrayItemVisitor) {
	for i, value := range a {
		visitor.VisitArrayItem(i, value)
	}
}

type efficientMockArray struct {
	v []int
}

func (a *efficientMockArray) ArrayItemCount() int {
	return len(a.v)
}

func (a *efficientMockArray) AcceptArrayItemVisitor(visitor ArrayItemVisitor) {
	for i, value := range a.v {
		visitor.VisitArrayItem(i, Int(value))
	}
}

type efficientMockArrayWithSnapshot struct {
	efficientMockArray
}

func (a *efficientMockArrayWithSnapshot) TakeSnapshot() interface{} {
	return &efficientMockArrayWithSnapshot{efficientMockArray{append([]int{}, a.v...)}}
}

type mockArraySnapshotter []Value

func (a mockArraySnapshotter) ArrayItemCount() int {
	return len(a)
}

func (a mockArraySnapshotter) AcceptArrayItemVisitor(visitor ArrayItemVisitor) {
	for i, value := range a {
		visitor.VisitArrayItem(i, value)
	}
}

func (a mockArraySnapshotter) TakeSnapshot() interface{} {
	return mockArraySnapshotter(append([]Value{}, a...))
}

type mockObject map[string]Value

func (o mockObject) ObjectFieldCount() int {
	return len(o)
}

func (o mockObject) AcceptObjectFieldVisitor(visitor ObjectFieldVisitor) {
	for key, value := range o {
		visitor.VisitObjectField(key, value)
	}
}

type mockObjectSnapshotter map[string]Value

func (o mockObjectSnapshotter) ObjectFieldCount() int {
	return len(o)
}

func (o mockObjectSnapshotter) AcceptObjectFieldVisitor(visitor ObjectFieldVisitor) {
	for key, value := range o {
		visitor.VisitObjectField(key, value)
	}
}

func (o mockObjectSnapshotter) TakeSnapshot() interface{} {
	cc := make(mockObjectSnapshotter, len(o))
	for k, v := range o {
		cc[k] = v
	}

	return cc
}

type stdLogger interface {
	Fatal(args ...interface{})
}

type mockVisitor struct {
	t stdLogger
}

func (v mockVisitor) VisitNone() {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitAny(interface{}) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitBool(bool) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInt(int) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInt8(int8) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInt16(int16) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInt32(int32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInt64(int64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUint(uint) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUint8(uint8) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUint16(uint16) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUint32(uint32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUint64(uint64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitFloat32(float32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitFloat64(float64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitDuration(time.Duration) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitError(error) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitTime(time.Time) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitString(string) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitStrings([]string) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitBytes([]byte) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitBools([]bool) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInts([]int) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInts8([]int8) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInts16([]int16) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInts32([]int32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitInts64([]int64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUints([]uint) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUints8([]uint8) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUints16([]uint16) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUints32([]uint32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitUints64([]uint64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitFloats32([]float32) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitFloats64([]float64) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitDurations([]time.Duration) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitArray(ValueArray) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitObject(ValueObject) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitStringer(fmt.Stringer) {
	v.t.Fatal("unexpected function call")
}

func (v mockVisitor) VisitFormattable(string, interface{}) {
	v.t.Fatal("unexpected function call")
}

func newMockNoneVisitor(t *testing.T) *mockNoneVisitor {
	return &mockNoneVisitor{mockVisitor: mockVisitor{t}}
}

type mockNoneVisitor struct {
	mockVisitor
	visited bool
}

func (v *mockNoneVisitor) VisitNone() {
	v.visited = true
}

func newMockAnyVisitor(t *testing.T) *mockAnyVisitor {
	return &mockAnyVisitor{mockVisitor: mockVisitor{t}}
}

type mockAnyVisitor struct {
	mockVisitor
	value   interface{}
	visited bool
}

func (v *mockAnyVisitor) VisitAny(value interface{}) {
	v.value = value
	v.visited = true
}

func newMockBoolVisitor(t *testing.T) *mockBoolVisitor {
	return &mockBoolVisitor{mockVisitor: mockVisitor{t}}
}

type mockBoolVisitor struct {
	mockVisitor
	value   bool
	visited bool
}

func (v *mockBoolVisitor) VisitBool(value bool) {
	v.value = value
	v.visited = true
}

func newMockIntVisitor(t *testing.T) *mockIntVisitor {
	return &mockIntVisitor{mockVisitor: mockVisitor{t}}
}

type mockIntVisitor struct {
	mockVisitor
	value   int
	visited bool
}

func (v *mockIntVisitor) VisitInt(value int) {
	v.value = value
	v.visited = true
}

func newMockInt8Visitor(t *testing.T) *mockInt8Visitor {
	return &mockInt8Visitor{mockVisitor: mockVisitor{t}}
}

type mockInt8Visitor struct {
	mockVisitor
	value   int8
	visited bool
}

func (v *mockInt8Visitor) VisitInt8(value int8) {
	v.value = value
	v.visited = true
}

func newMockInt16Visitor(t *testing.T) *mockInt16Visitor {
	return &mockInt16Visitor{mockVisitor: mockVisitor{t}}
}

type mockInt16Visitor struct {
	mockVisitor
	value   int16
	visited bool
}

func (v *mockInt16Visitor) VisitInt16(value int16) {
	v.value = value
	v.visited = true
}

func newMockInt32Visitor(t *testing.T) *mockInt32Visitor {
	return &mockInt32Visitor{mockVisitor: mockVisitor{t}}
}

type mockInt32Visitor struct {
	mockVisitor
	value   int32
	visited bool
}

func (v *mockInt32Visitor) VisitInt32(value int32) {
	v.value = value
	v.visited = true
}

func newMockInt64Visitor(t *testing.T) *mockInt64Visitor {
	return &mockInt64Visitor{mockVisitor: mockVisitor{t}}
}

type mockInt64Visitor struct {
	mockVisitor
	value   int64
	visited bool
}

func (v *mockInt64Visitor) VisitInt64(value int64) {
	v.value = value
	v.visited = true
}

func newMockUintVisitor(t *testing.T) *mockUintVisitor {
	return &mockUintVisitor{mockVisitor: mockVisitor{t}}
}

type mockUintVisitor struct {
	mockVisitor
	value   uint
	visited bool
}

func (v *mockUintVisitor) VisitUint(value uint) {
	v.value = value
	v.visited = true
}

func newMockUint8Visitor(t *testing.T) *mockUint8Visitor {
	return &mockUint8Visitor{mockVisitor: mockVisitor{t}}
}

type mockUint8Visitor struct {
	mockVisitor
	value   uint8
	visited bool
}

func (v *mockUint8Visitor) VisitUint8(value uint8) {
	v.value = value
	v.visited = true
}

func newMockUint16Visitor(t *testing.T) *mockUint16Visitor {
	return &mockUint16Visitor{mockVisitor: mockVisitor{t}}
}

type mockUint16Visitor struct {
	mockVisitor
	value   uint16
	visited bool
}

func (v *mockUint16Visitor) VisitUint16(value uint16) {
	v.value = value
	v.visited = true
}

func newMockUint32Visitor(t *testing.T) *mockUint32Visitor {
	return &mockUint32Visitor{mockVisitor: mockVisitor{t}}
}

type mockUint32Visitor struct {
	mockVisitor
	value   uint32
	visited bool
}

func (v *mockUint32Visitor) VisitUint32(value uint32) {
	v.value = value
	v.visited = true
}

func newMockUint64Visitor(t *testing.T) *mockUint64Visitor {
	return &mockUint64Visitor{mockVisitor: mockVisitor{t}}
}

type mockUint64Visitor struct {
	mockVisitor
	value   uint64
	visited bool
}

func (v *mockUint64Visitor) VisitUint64(value uint64) {
	v.value = value
	v.visited = true
}

func newMockFloat32Visitor(t *testing.T) *mockFloat32Visitor {
	return &mockFloat32Visitor{mockVisitor: mockVisitor{t}}
}

type mockFloat32Visitor struct {
	mockVisitor
	value   float32
	visited bool
}

func (v *mockFloat32Visitor) VisitFloat32(value float32) {
	v.value = value
	v.visited = true
}

func newMockFloat64Visitor(t *testing.T) *mockFloat64Visitor {
	return &mockFloat64Visitor{mockVisitor: mockVisitor{t}}
}

type mockFloat64Visitor struct {
	mockVisitor
	value   float64
	visited bool
}

func (v *mockFloat64Visitor) VisitFloat64(value float64) {
	v.value = value
	v.visited = true
}

func newMockStringVisitor(t *testing.T) *mockStringVisitor {
	return &mockStringVisitor{mockVisitor: mockVisitor{t}}
}

type mockStringVisitor struct {
	mockVisitor
	value   string
	visited bool
}

func (v *mockStringVisitor) VisitString(value string) {
	v.value = value
	v.visited = true
}

func newMockStringerVisitor(t *testing.T) *mockStringerVisitor {
	return &mockStringerVisitor{mockVisitor: mockVisitor{t}}
}

type mockStringerVisitor struct {
	mockVisitor
	value   fmt.Stringer
	visited bool
}

func (v *mockStringerVisitor) VisitStringer(value fmt.Stringer) {
	v.value = value
	v.visited = true
}

func newMockFormattableVisitor(t *testing.T) *mockFormattableVisitor {
	return &mockFormattableVisitor{mockVisitor: mockVisitor{t}}
}

type mockFormattableVisitor struct {
	mockVisitor
	format  string
	value   interface{}
	visited bool
}

func (v *mockFormattableVisitor) VisitFormattable(format string, value interface{}) {
	v.format = format
	v.value = value
	v.visited = true
}

func newMockDurationVisitor(t *testing.T) *mockDurationVisitor {
	return &mockDurationVisitor{mockVisitor: mockVisitor{t}}
}

type mockDurationVisitor struct {
	mockVisitor
	value   time.Duration
	visited bool
}

func (v *mockDurationVisitor) VisitDuration(value time.Duration) {
	v.value = value
	v.visited = true
}

func newMockTimeVisitor(t *testing.T) *mockTimeVisitor {
	return &mockTimeVisitor{mockVisitor: mockVisitor{t}}
}

type mockTimeVisitor struct {
	mockVisitor
	value   time.Time
	visited bool
}

func (v *mockTimeVisitor) VisitTime(value time.Time) {
	v.value = value
	v.visited = true
}

func newMockErrorVisitor(t *testing.T) *mockErrorVisitor {
	return &mockErrorVisitor{mockVisitor: mockVisitor{t}}
}

type mockErrorVisitor struct {
	mockVisitor
	value   error
	visited bool
}

func (v *mockErrorVisitor) VisitError(value error) {
	v.value = value
	v.visited = true
}

func newMockBytesVisitor(t *testing.T) *mockBytesVisitor {
	return &mockBytesVisitor{mockVisitor: mockVisitor{t}}
}

type mockBytesVisitor struct {
	mockVisitor
	value   []byte
	visited bool
}

func (v *mockBytesVisitor) VisitBytes(value []byte) {
	v.value = value
	v.visited = true
}

func newMockIntsVisitor(t *testing.T) *mockIntsVisitor {
	return &mockIntsVisitor{mockVisitor: mockVisitor{t}}
}

type mockIntsVisitor struct {
	mockVisitor
	value   []int
	visited bool
}

func (v *mockIntsVisitor) VisitInts(value []int) {
	v.value = value
	v.visited = true
}

func newMockInts8Visitor(t *testing.T) *mockInts8Visitor {
	return &mockInts8Visitor{mockVisitor: mockVisitor{t}}
}

type mockInts8Visitor struct {
	mockVisitor
	value   []int8
	visited bool
}

func (v *mockInts8Visitor) VisitInts8(value []int8) {
	v.value = value
	v.visited = true
}

func newMockInts16Visitor(t *testing.T) *mockInts16Visitor {
	return &mockInts16Visitor{mockVisitor: mockVisitor{t}}
}

type mockInts16Visitor struct {
	mockVisitor
	value   []int16
	visited bool
}

func (v *mockInts16Visitor) VisitInts16(value []int16) {
	v.value = value
	v.visited = true
}

func newMockInts32Visitor(t *testing.T) *mockInts32Visitor {
	return &mockInts32Visitor{mockVisitor: mockVisitor{t}}
}

type mockInts32Visitor struct {
	mockVisitor
	value   []int32
	visited bool
}

func (v *mockInts32Visitor) VisitInts32(value []int32) {
	v.value = value
	v.visited = true
}

func newMockInts64Visitor(t *testing.T) *mockInts64Visitor {
	return &mockInts64Visitor{mockVisitor: mockVisitor{t}}
}

type mockInts64Visitor struct {
	mockVisitor
	value   []int64
	visited bool
}

func (v *mockInts64Visitor) VisitInts64(value []int64) {
	v.value = value
	v.visited = true
}

func newMockUintsVisitor(t *testing.T) *mockUintsVisitor {
	return &mockUintsVisitor{mockVisitor: mockVisitor{t}}
}

type mockUintsVisitor struct {
	mockVisitor
	value   []uint
	visited bool
}

func (v *mockUintsVisitor) VisitUints(value []uint) {
	v.value = value
	v.visited = true
}

func newMockUints8Visitor(t *testing.T) *mockUints8Visitor {
	return &mockUints8Visitor{mockVisitor: mockVisitor{t}}
}

type mockUints8Visitor struct {
	mockVisitor
	value   []uint8
	visited bool
}

func (v *mockUints8Visitor) VisitUints8(value []uint8) {
	v.value = value
	v.visited = true
}

func newMockUints16Visitor(t *testing.T) *mockUints16Visitor {
	return &mockUints16Visitor{mockVisitor: mockVisitor{t}}
}

type mockUints16Visitor struct {
	mockVisitor
	value   []uint16
	visited bool
}

func (v *mockUints16Visitor) VisitUints16(value []uint16) {
	v.value = value
	v.visited = true
}

func newMockUints32Visitor(t *testing.T) *mockUints32Visitor {
	return &mockUints32Visitor{mockVisitor: mockVisitor{t}}
}

type mockUints32Visitor struct {
	mockVisitor
	value   []uint32
	visited bool
}

func (v *mockUints32Visitor) VisitUints32(value []uint32) {
	v.value = value
	v.visited = true
}

func newMockUints64Visitor(t *testing.T) *mockUints64Visitor {
	return &mockUints64Visitor{mockVisitor: mockVisitor{t}}
}

type mockUints64Visitor struct {
	mockVisitor
	value   []uint64
	visited bool
}

func (v *mockUints64Visitor) VisitUints64(value []uint64) {
	v.value = value
	v.visited = true
}

func newMockBoolsVisitor(t *testing.T) *mockBoolsVisitor {
	return &mockBoolsVisitor{mockVisitor: mockVisitor{t}}
}

type mockBoolsVisitor struct {
	mockVisitor
	value   []bool
	visited bool
}

func (v *mockBoolsVisitor) VisitBools(value []bool) {
	v.value = value
	v.visited = true
}

func newMockFloats32Visitor(t *testing.T) *mockFloats32Visitor {
	return &mockFloats32Visitor{mockVisitor: mockVisitor{t}}
}

type mockFloats32Visitor struct {
	mockVisitor
	value   []float32
	visited bool
}

func (v *mockFloats32Visitor) VisitFloats32(value []float32) {
	v.value = value
	v.visited = true
}

func newMockFloats64Visitor(t *testing.T) *mockFloats64Visitor {
	return &mockFloats64Visitor{mockVisitor: mockVisitor{t}}
}

type mockFloats64Visitor struct {
	mockVisitor
	value   []float64
	visited bool
}

func (v *mockFloats64Visitor) VisitFloats64(value []float64) {
	v.value = value
	v.visited = true
}

func newMockStringsVisitor(t *testing.T) *mockStringsVisitor {
	return &mockStringsVisitor{mockVisitor: mockVisitor{t}}
}

type mockStringsVisitor struct {
	mockVisitor
	value   []string
	visited bool
}

func (v *mockStringsVisitor) VisitStrings(value []string) {
	v.value = value
	v.visited = true
}

func newMockDurationsVisitor(t *testing.T) *mockDurationsVisitor {
	return &mockDurationsVisitor{mockVisitor: mockVisitor{t}}
}

type mockDurationsVisitor struct {
	mockVisitor
	value   []time.Duration
	visited bool
}

func (v *mockDurationsVisitor) VisitDurations(value []time.Duration) {
	v.value = value
	v.visited = true
}

func newMockArrayVisitor(t stdLogger) *mockArrayVisitor {
	return &mockArrayVisitor{mockVisitor: mockVisitor{t}}
}

type mockArrayVisitor struct {
	mockVisitor
	value   []Value
	visited bool
}

func (v *mockArrayVisitor) VisitArray(array ValueArray) {
	if array != nil {
		v.value = make([]Value, array.ArrayItemCount())
		array.AcceptArrayItemVisitor(v)
	}
	v.visited = true
}

func (v *mockArrayVisitor) VisitArrayItem(index int, value Value) {
	v.value[index] = value
}

func newMockObjectVisitor(t *testing.T) *mockObjectVisitor {
	return &mockObjectVisitor{mockVisitor: mockVisitor{t}}
}

type mockObjectVisitor struct {
	mockVisitor
	count   int
	value   map[string]Value
	visited bool
}

func (v *mockObjectVisitor) VisitObject(object ValueObject) {
	if object != nil {
		v.value = map[string]Value{}
		v.count = object.ObjectFieldCount()
		object.AcceptObjectFieldVisitor(v)
	}
	v.visited = true
}

func (v *mockObjectVisitor) VisitObjectField(key string, value Value) {
	v.value[key] = value
}

func TestValueNone(t *testing.T) {
	visitor := newMockNoneVisitor(t)
	value := Value{}
	require.Equal(t, TypeNone, value.Type())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
}

func TestValueNoneSnapshot(t *testing.T) {
	visitor := newMockNoneVisitor(t)
	value := Value{}.Snapshot()
	require.Equal(t, TypeNone, value.Type())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
}

func TestValueBool(t *testing.T) {
	visitor := newMockBoolVisitor(t)
	value := Bool(true)
	require.Equal(t, TypeBool, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, true, visitor.value)
}

func TestValueAnyBool(t *testing.T) {
	visitor := newMockBoolVisitor(t)
	value := Any(true)
	require.Equal(t, TypeBool, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, true, visitor.value)
}

func TestValueConstAnyBool(t *testing.T) {
	visitor := newMockBoolVisitor(t)
	value := ConstAny(true)
	require.Equal(t, TypeBool, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, true, visitor.value)
}

func BenchmarkValueBoolConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Bool(true)
	}
	avoidOptimization(r)
}

func TestValueInt(t *testing.T) {
	visitor := newMockIntVisitor(t)
	value := Int(42)
	require.Equal(t, TypeInt, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, 42, visitor.value)
}

func TestValueAnyInt(t *testing.T) {
	visitor := newMockIntVisitor(t)
	value := Any(42)
	require.Equal(t, TypeInt, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, 42, visitor.value)
}

func TestValueConstAnyInt(t *testing.T) {
	visitor := newMockIntVisitor(t)
	value := ConstAny(42)
	require.Equal(t, TypeInt, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, 42, visitor.value)
}

func BenchmarkValueIntConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Int(42)
	}
	avoidOptimization(r)
}

func TestValueInt8(t *testing.T) {
	visitor := newMockInt8Visitor(t)
	value := Int8(42)
	require.Equal(t, TypeInt8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int8(42), visitor.value)
}

func TestValueAnyInt8(t *testing.T) {
	visitor := newMockInt8Visitor(t)
	value := Any(int8(42))
	require.Equal(t, TypeInt8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int8(42), visitor.value)
}

func TestValueConstAnyInt8(t *testing.T) {
	visitor := newMockInt8Visitor(t)
	value := ConstAny(int8(42))
	require.Equal(t, TypeInt8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int8(42), visitor.value)
}

func BenchmarkValueInt8Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Int8(42)
	}
	avoidOptimization(r)
}
func TestValueInt16(t *testing.T) {
	visitor := newMockInt16Visitor(t)
	value := Int16(42)
	require.Equal(t, TypeInt16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int16(42), visitor.value)
}

func TestValueAnyInt16(t *testing.T) {
	visitor := newMockInt16Visitor(t)
	value := Any(int16(42))
	require.Equal(t, TypeInt16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int16(42), visitor.value)
}

func TestValueConstAnyInt16(t *testing.T) {
	visitor := newMockInt16Visitor(t)
	value := ConstAny(int16(42))
	require.Equal(t, TypeInt16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int16(42), visitor.value)
}

func BenchmarkValueInt16Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Int16(42)
	}
	avoidOptimization(r)
}

func BenchmarkValueAnyInt16Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Any(int16(42))
	}
	avoidOptimization(r)
}

func TestValueInt32(t *testing.T) {
	visitor := newMockInt32Visitor(t)
	value := Int32(42)
	require.Equal(t, TypeInt32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int32(42), visitor.value)
}

func TestValueAnyInt32(t *testing.T) {
	visitor := newMockInt32Visitor(t)
	value := Any(int32(42))
	require.Equal(t, TypeInt32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int32(42), visitor.value)
}

func TestValueConstAnyInt32(t *testing.T) {
	visitor := newMockInt32Visitor(t)
	value := ConstAny(int32(42))
	require.Equal(t, TypeInt32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int32(42), visitor.value)
}

func BenchmarkValueInt32Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Int32(42)
	}
	avoidOptimization(r)
}

func TestValueInt64(t *testing.T) {
	visitor := newMockInt64Visitor(t)
	value := Int64(42)
	require.Equal(t, TypeInt64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int64(42), visitor.value)
}

func TestValueAnyInt64(t *testing.T) {
	visitor := newMockInt64Visitor(t)
	value := Any(int64(42))
	require.Equal(t, TypeInt64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int64(42), visitor.value)
}

func TestValueConstAnyInt64(t *testing.T) {
	visitor := newMockInt64Visitor(t)
	value := ConstAny(int64(42))
	require.Equal(t, TypeInt64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, int64(42), visitor.value)
}

func BenchmarkValueInt64Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Int64(42)
	}
	avoidOptimization(r)
}

func TestValueUint(t *testing.T) {
	visitor := newMockUintVisitor(t)
	value := Uint(42)
	require.Equal(t, TypeUint, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint(42), visitor.value)
}

func TestValueAnyUint(t *testing.T) {
	visitor := newMockUintVisitor(t)
	value := Any(uint(42))
	require.Equal(t, TypeUint, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint(42), visitor.value)
}

func TestValueConstAnyUint(t *testing.T) {
	visitor := newMockUintVisitor(t)
	value := ConstAny(uint(42))
	require.Equal(t, TypeUint, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint(42), visitor.value)
}

func BenchmarkValueUintConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Uint(42)
	}
	avoidOptimization(r)
}

func TestValueUint8(t *testing.T) {
	visitor := newMockUint8Visitor(t)
	value := Uint8(42)
	require.Equal(t, TypeUint8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint8(42), visitor.value)
}

func TestValueAnyUint8(t *testing.T) {
	visitor := newMockUint8Visitor(t)
	value := Any(uint8(42))
	require.Equal(t, TypeUint8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint8(42), visitor.value)
}

func TestValueConstAnyUint8(t *testing.T) {
	visitor := newMockUint8Visitor(t)
	value := ConstAny(uint8(42))
	require.Equal(t, TypeUint8, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint8(42), visitor.value)
}

func BenchmarkValueUint8Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Uint8(42)
	}
	avoidOptimization(r)
}

func TestValueUint16(t *testing.T) {
	visitor := newMockUint16Visitor(t)
	value := Uint16(42)
	require.Equal(t, TypeUint16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint16(42), visitor.value)
}

func TestValueAnyUint16(t *testing.T) {
	visitor := newMockUint16Visitor(t)
	value := Any(uint16(42))
	require.Equal(t, TypeUint16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint16(42), visitor.value)
}

func TestValueConstAnyUint16(t *testing.T) {
	visitor := newMockUint16Visitor(t)
	value := ConstAny(uint16(42))
	require.Equal(t, TypeUint16, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint16(42), visitor.value)
}

func BenchmarkValueUint16Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Uint16(42)
	}
	avoidOptimization(r)
}

func TestValueUint32(t *testing.T) {
	visitor := newMockUint32Visitor(t)
	value := Uint32(42)
	require.Equal(t, TypeUint32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint32(42), visitor.value)
}

func TestValueAnyUint32(t *testing.T) {
	visitor := newMockUint32Visitor(t)
	value := Any(uint32(42))
	require.Equal(t, TypeUint32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint32(42), visitor.value)
}

func TestValueConstAnyUint32(t *testing.T) {
	visitor := newMockUint32Visitor(t)
	value := ConstAny(uint32(42))
	require.Equal(t, TypeUint32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint32(42), visitor.value)
}

func BenchmarkValueUint32Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Uint32(42)
	}
	avoidOptimization(r)
}

func TestValueUint64(t *testing.T) {
	visitor := newMockUint64Visitor(t)
	value := Uint64(42)
	require.Equal(t, TypeUint64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint64(42), visitor.value)
}

func TestValueAnyUint64(t *testing.T) {
	visitor := newMockUint64Visitor(t)
	value := Any(uint64(42))
	require.Equal(t, TypeUint64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint64(42), visitor.value)
}

func TestValueConstAnyUint64(t *testing.T) {
	visitor := newMockUint64Visitor(t)
	value := ConstAny(uint64(42))
	require.Equal(t, TypeUint64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, uint64(42), visitor.value)
}

func BenchmarkValueUint64Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Uint64(42)
	}
	avoidOptimization(r)
}

func TestValueFloat32(t *testing.T) {
	visitor := newMockFloat32Visitor(t)
	value := Float32(0.42)
	require.Equal(t, TypeFloat32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float32(0.42), visitor.value)
}

func TestValueAnyFloat32(t *testing.T) {
	visitor := newMockFloat32Visitor(t)
	value := Any(float32(0.42))
	require.Equal(t, TypeFloat32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float32(0.42), visitor.value)
}

func TestValueConstAnyFloat32(t *testing.T) {
	visitor := newMockFloat32Visitor(t)
	value := ConstAny(float32(0.42))
	require.Equal(t, TypeFloat32, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float32(0.42), visitor.value)
}

func BenchmarkValueFloat32Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Float32(0.42)
	}
	avoidOptimization(r)
}

func TestValueFloat64(t *testing.T) {
	visitor := newMockFloat64Visitor(t)
	value := Float64(0.42)
	require.Equal(t, TypeFloat64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float64(0.42), visitor.value)
}

func TestValueAnyFloat64(t *testing.T) {
	visitor := newMockFloat64Visitor(t)
	value := Any(float64(0.42))
	require.Equal(t, TypeFloat64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float64(0.42), visitor.value)
}

func TestValueConstAnyFloat64(t *testing.T) {
	visitor := newMockFloat64Visitor(t)
	value := ConstAny(float64(0.42))
	require.Equal(t, TypeFloat64, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, float64(0.42), visitor.value)
}

func BenchmarkValueFloat64Construction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Float64(0.42)
	}
	avoidOptimization(r)
}

func TestValueString(t *testing.T) {
	visitor := newMockStringVisitor(t)
	v := "test"
	value := String(v)
	require.Equal(t, TypeString, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueAnyString(t *testing.T) {
	visitor := newMockStringVisitor(t)
	v := "test"
	value := Any(v)
	require.Equal(t, TypeString, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyString(t *testing.T) {
	visitor := newMockStringVisitor(t)
	v := "test"
	value := ConstAny(v)
	require.Equal(t, TypeString, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func BenchmarkValueStringConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = String("test")
	}
	avoidOptimization(r)
}

func TestValueStringer(t *testing.T) {
	visitor := newMockStringerVisitor(t)
	v := testStringer("test")
	value := Stringer(v)
	require.Equal(t, TypeStringer, value.Type())
	require.Equal(t, false, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstStringer(t *testing.T) {
	visitor := newMockStringerVisitor(t)
	v := testStringer("test")
	value := ConstStringer(v)
	require.Equal(t, TypeStringer, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilStringer(t *testing.T) {
	visitor := newMockStringerVisitor(t)
	value := Stringer(nil)
	require.Equal(t, TypeStringer, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, nil, visitor.value)
}

func TestValueAnyStringer(t *testing.T) {
	visitor := newMockStringerVisitor(t)
	v := testStringer("test")
	value := Any(v)
	require.Equal(t, TypeStringer, value.Type())
	require.Equal(t, false, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyStringer(t *testing.T) {
	visitor := newMockStringerVisitor(t)
	v := testStringer("test")
	value := ConstAny(v)
	require.Equal(t, TypeStringer, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func BenchmarkValueStringerConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Stringer(testStringer("test"))
	}
	avoidOptimization(r)
}

func TestValueFormattable(t *testing.T) {
	visitor := newMockFormattableVisitor(t)
	v := "test"
	f := &mockFormatter{t: t, value: v}
	value := Formattable("%#v", f)
	require.Equal(t, TypeFormattable, value.Type())
	require.Equal(t, false, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, "%#v", visitor.format)
	require.Equal(t, f, visitor.value)
}

func TestValueConstFormattable(t *testing.T) {
	visitor := newMockFormattableVisitor(t)
	v := "test"
	f := &mockFormatter{t: t, value: v}
	value := ConstFormattable("%#v", f)
	require.Equal(t, TypeFormattable, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, "%#v", visitor.format)
	require.Equal(t, f, visitor.value)
}

func TestValueDuration(t *testing.T) {
	visitor := newMockDurationVisitor(t)
	value := Duration(time.Second)
	require.Equal(t, TypeDuration, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, time.Second, visitor.value)
}

func TestValueAnyDuration(t *testing.T) {
	visitor := newMockDurationVisitor(t)
	value := Any(time.Second)
	require.Equal(t, TypeDuration, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, time.Second, visitor.value)
}

func TestValueConstAnyDuration(t *testing.T) {
	visitor := newMockDurationVisitor(t)
	value := ConstAny(time.Second)
	require.Equal(t, TypeDuration, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, time.Second, visitor.value)
}

func BenchmarkValueDurationConstruction(b *testing.B) {
	var r Value
	for i := 0; i != b.N; i++ {
		r = Duration(time.Second)
	}
	avoidOptimization(r)
}

func TestValueTime(t *testing.T) {
	visitor := newMockTimeVisitor(t)
	now := time.Now()
	value := Time(now)
	require.Equal(t, TypeTime, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, now.UnixNano(), visitor.value.UnixNano())
	require.Equal(t, now.Location(), visitor.value.Location())
}

func TestValueEmptyTime(t *testing.T) {
	visitor := newMockTimeVisitor(t)
	v := time.Time{}
	value := Time(v)
	require.Equal(t, TypeTime, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v.UnixNano(), visitor.value.UnixNano())
	require.Equal(t, v.Location(), visitor.value.Location())
}

func TestValueAnyTime(t *testing.T) {
	visitor := newMockTimeVisitor(t)
	now := time.Now()
	value := Any(now)
	require.Equal(t, TypeTime, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, now.UnixNano(), visitor.value.UnixNano())
	require.Equal(t, now.Location(), visitor.value.Location())
}

func TestValueConstAnyTime(t *testing.T) {
	visitor := newMockTimeVisitor(t)
	now := time.Now()
	value := ConstAny(now)
	require.Equal(t, TypeTime, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, now.UnixNano(), visitor.value.UnixNano())
	require.Equal(t, now.Location(), visitor.value.Location())
}

func BenchmarkValueTimeConstruction(b *testing.B) {
	var r Value
	now := time.Now()
	for i := 0; i != b.N; i++ {
		r = Time(now)
	}
	avoidOptimization(r)
}

func TestValueError(t *testing.T) {
	visitor := newMockErrorVisitor(t)
	err := errors.New("some error")
	value := Error(err)
	require.Equal(t, TypeError, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, err, visitor.value)
}

func TestValueNilError(t *testing.T) {
	visitor := newMockErrorVisitor(t)
	var err error
	value := Error(err)
	require.Equal(t, TypeError, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, err, visitor.value)
}

func TestValueAnyError(t *testing.T) {
	visitor := newMockErrorVisitor(t)
	err := errors.New("some error")
	value := Any(err)
	require.Equal(t, TypeError, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, err, visitor.value)
}

func TestValueConstAnyError(t *testing.T) {
	visitor := newMockErrorVisitor(t)
	err := errors.New("some error")
	value := ConstAny(err)
	require.Equal(t, TypeError, value.Type())
	require.Equal(t, true, value.Const())
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, err, visitor.value)
}

func TestValueBytes(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []byte("some value")
	value := Bytes(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstBytes(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []byte("some value")
	value := Bytes(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilBytes(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	value := Bytes(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []byte(nil), visitor.value)
}

func TestValueAnyBytes(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []byte("some value")
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyBytes(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []byte("some value")
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueAnyNil(t *testing.T) {
	visitor := newMockAnyVisitor(t)
	value := Any(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, nil, visitor.value)
}

func TestValueConstAnyNil(t *testing.T) {
	visitor := newMockAnyVisitor(t)
	value := ConstAny(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, nil, visitor.value)
}

func BenchmarkValueErrorConstruction(b *testing.B) {
	var r Value
	err := errors.New("some error")
	for i := 0; i != b.N; i++ {
		r = Error(err)
	}
	avoidOptimization(r)
}

func TestValueInts(t *testing.T) {
	visitor := newMockIntsVisitor(t)
	v := []int{1, 2, 3}
	value := Ints(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstInts(t *testing.T) {
	visitor := newMockIntsVisitor(t)
	v := []int{1, 2, 3}
	value := ConstInts(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilInts(t *testing.T) {
	visitor := newMockIntsVisitor(t)
	value := Ints(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []int(nil), visitor.value)
}

func TestValueAnyInts(t *testing.T) {
	visitor := newMockIntsVisitor(t)
	v := []int{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyInts(t *testing.T) {
	visitor := newMockIntsVisitor(t)
	v := []int{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func BenchmarkValueIntsConstruction(b *testing.B) {
	var r Value
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i != b.N; i++ {
		r = Ints(v)
	}
	avoidOptimization(r)
}

func BenchmarkValueIntsSnapshot(b *testing.B) {
	var r Value
	v := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i != b.N; i++ {
		r = Ints(v)
		Snapshot(&r)
	}
	avoidOptimization(r)
}

func TestValueInts8(t *testing.T) {
	visitor := newMockInts8Visitor(t)
	v := []int8{1, 2, 3}
	value := Ints8(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstInts8(t *testing.T) {
	visitor := newMockInts8Visitor(t)
	v := []int8{1, 2, 3}
	value := ConstInts8(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilInts8(t *testing.T) {
	visitor := newMockInts8Visitor(t)
	value := Ints8(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []int8(nil), visitor.value)
}

func TestValueAnyInts8(t *testing.T) {
	visitor := newMockInts8Visitor(t)
	v := []int8{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyInts8(t *testing.T) {
	visitor := newMockInts8Visitor(t)
	v := []int8{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueInts16(t *testing.T) {
	visitor := newMockInts16Visitor(t)
	v := []int16{1, 2, 3}
	value := Ints16(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstInts16(t *testing.T) {
	visitor := newMockInts16Visitor(t)
	v := []int16{1, 2, 3}
	value := ConstInts16(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilInts16(t *testing.T) {
	visitor := newMockInts16Visitor(t)
	value := Ints16(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []int16(nil), visitor.value)
}

func TestValueAnyInts16(t *testing.T) {
	visitor := newMockInts16Visitor(t)
	v := []int16{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyInts16(t *testing.T) {
	visitor := newMockInts16Visitor(t)
	v := []int16{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueInts32(t *testing.T) {
	visitor := newMockInts32Visitor(t)
	v := []int32{1, 2, 3}
	value := Ints32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstInts32(t *testing.T) {
	visitor := newMockInts32Visitor(t)
	v := []int32{1, 2, 3}
	value := ConstInts32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilInts32(t *testing.T) {
	visitor := newMockInts32Visitor(t)
	value := Ints32(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []int32(nil), visitor.value)
}

func TestValueAnyInts32(t *testing.T) {
	visitor := newMockInts32Visitor(t)
	v := []int32{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyInts32(t *testing.T) {
	visitor := newMockInts32Visitor(t)
	v := []int32{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueInts64(t *testing.T) {
	visitor := newMockInts64Visitor(t)
	v := []int64{1, 2, 3}
	value := Ints64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstInts64(t *testing.T) {
	visitor := newMockInts64Visitor(t)
	v := []int64{1, 2, 3}
	value := ConstInts64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilInts64(t *testing.T) {
	visitor := newMockInts64Visitor(t)
	value := Ints64(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []int64(nil), visitor.value)
}

func TestValueAnyInts64(t *testing.T) {
	visitor := newMockInts64Visitor(t)
	v := []int64{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyInts64(t *testing.T) {
	visitor := newMockInts64Visitor(t)
	v := []int64{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueUints(t *testing.T) {
	visitor := newMockUintsVisitor(t)
	v := []uint{1, 2, 3}
	value := Uints(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstUints(t *testing.T) {
	visitor := newMockUintsVisitor(t)
	v := []uint{1, 2, 3}
	value := Uints(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilUints(t *testing.T) {
	visitor := newMockUintsVisitor(t)
	value := Uints(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []uint(nil), visitor.value)
}

func TestValueAnyUints(t *testing.T) {
	visitor := newMockUintsVisitor(t)
	v := []uint{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyUints(t *testing.T) {
	visitor := newMockUintsVisitor(t)
	v := []uint{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueUints8(t *testing.T) {
	visitor := newMockUints8Visitor(t)
	v := []uint8{1, 2, 3}
	value := Uints8(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstUints8(t *testing.T) {
	visitor := newMockUints8Visitor(t)
	v := []uint8{1, 2, 3}
	value := ConstUints8(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilUints8(t *testing.T) {
	visitor := newMockUints8Visitor(t)
	value := Uints8(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []uint8(nil), visitor.value)
}

func TestValueAnyUints8(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []uint8{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyUints8(t *testing.T) {
	visitor := newMockBytesVisitor(t)
	v := []uint8{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueUints16(t *testing.T) {
	visitor := newMockUints16Visitor(t)
	v := []uint16{1, 2, 3}
	value := Uints16(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstUints16(t *testing.T) {
	visitor := newMockUints16Visitor(t)
	v := []uint16{1, 2, 3}
	value := ConstUints16(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilUints16(t *testing.T) {
	visitor := newMockUints16Visitor(t)
	value := Uints16(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []uint16(nil), visitor.value)
}

func TestValueAnyUints16(t *testing.T) {
	visitor := newMockUints16Visitor(t)
	v := []uint16{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyUints16(t *testing.T) {
	visitor := newMockUints16Visitor(t)
	v := []uint16{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueUints32(t *testing.T) {
	visitor := newMockUints32Visitor(t)
	v := []uint32{1, 2, 3}
	value := Uints32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstUints32(t *testing.T) {
	visitor := newMockUints32Visitor(t)
	v := []uint32{1, 2, 3}
	value := ConstUints32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilUints32(t *testing.T) {
	visitor := newMockUints32Visitor(t)
	value := Uints32(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []uint32(nil), visitor.value)
}

func TestValueAnyUints32(t *testing.T) {
	visitor := newMockUints32Visitor(t)
	v := []uint32{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyUints32(t *testing.T) {
	visitor := newMockUints32Visitor(t)
	v := []uint32{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueUints64(t *testing.T) {
	visitor := newMockUints64Visitor(t)
	v := []uint64{1, 2, 3}
	value := Uints64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstUints64(t *testing.T) {
	visitor := newMockUints64Visitor(t)
	v := []uint64{1, 2, 3}
	value := ConstUints64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilUints64(t *testing.T) {
	visitor := newMockUints64Visitor(t)
	value := Uints64(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []uint64(nil), visitor.value)
}

func TestValueAnyUints64(t *testing.T) {
	visitor := newMockUints64Visitor(t)
	v := []uint64{1, 2, 3}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyUints64(t *testing.T) {
	visitor := newMockUints64Visitor(t)
	v := []uint64{1, 2, 3}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueBools(t *testing.T) {
	visitor := newMockBoolsVisitor(t)
	v := []bool{true, false}
	value := Bools(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstBools(t *testing.T) {
	visitor := newMockBoolsVisitor(t)
	v := []bool{true, false}
	value := ConstBools(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilBools(t *testing.T) {
	visitor := newMockBoolsVisitor(t)
	value := Bools(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []bool(nil), visitor.value)
}

func TestValueAnyBools(t *testing.T) {
	visitor := newMockBoolsVisitor(t)
	v := []bool{true, false}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyBools(t *testing.T) {
	visitor := newMockBoolsVisitor(t)
	v := []bool{true, false}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueFloats32(t *testing.T) {
	visitor := newMockFloats32Visitor(t)
	v := []float32{0.42, -0.42}
	value := Floats32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstFloats32(t *testing.T) {
	visitor := newMockFloats32Visitor(t)
	v := []float32{0.42, -0.42}
	value := ConstFloats32(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilFloats32(t *testing.T) {
	visitor := newMockFloats32Visitor(t)
	value := Floats32(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []float32(nil), visitor.value)
}

func TestValueAnyFloats32(t *testing.T) {
	visitor := newMockFloats32Visitor(t)
	v := []float32{0.42, -0.42}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyFloats32(t *testing.T) {
	visitor := newMockFloats32Visitor(t)
	v := []float32{0.42, -0.42}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueFloats64(t *testing.T) {
	visitor := newMockFloats64Visitor(t)
	v := []float64{0.42, -0.42}
	value := Floats64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstFloats64(t *testing.T) {
	visitor := newMockFloats64Visitor(t)
	v := []float64{0.42, -0.42}
	value := ConstFloats64(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilFloats64(t *testing.T) {
	visitor := newMockFloats64Visitor(t)
	value := Floats64(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []float64(nil), visitor.value)
}

func TestValueAnyFloats64(t *testing.T) {
	visitor := newMockFloats64Visitor(t)
	v := []float64{0.42, -0.42}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyFloats64(t *testing.T) {
	visitor := newMockFloats64Visitor(t)
	v := []float64{0.42, -0.42}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueStrings(t *testing.T) {
	visitor := newMockStringsVisitor(t)
	v := []string{"a", "b"}
	value := Strings(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstStrings(t *testing.T) {
	visitor := newMockStringsVisitor(t)
	v := []string{"a", "b"}
	value := ConstStrings(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilStrings(t *testing.T) {
	visitor := newMockStringsVisitor(t)
	value := Strings(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []string(nil), visitor.value)
}

func TestValueAnyStrings(t *testing.T) {
	visitor := newMockStringsVisitor(t)
	v := []string{"a", "b"}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyStrings(t *testing.T) {
	visitor := newMockStringsVisitor(t)
	v := []string{"a", "b"}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueDurations(t *testing.T) {
	visitor := newMockDurationsVisitor(t)
	v := []time.Duration{time.Second, time.Hour}
	value := Durations(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstDurations(t *testing.T) {
	visitor := newMockDurationsVisitor(t)
	v := []time.Duration{time.Second, time.Hour}
	value := Durations(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilDurations(t *testing.T) {
	visitor := newMockDurationsVisitor(t)
	value := Durations(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []time.Duration(nil), visitor.value)
}

func TestValueAnyDurations(t *testing.T) {
	visitor := newMockDurationsVisitor(t)
	v := []time.Duration{time.Second, time.Hour}
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyDurations(t *testing.T) {
	visitor := newMockDurationsVisitor(t)
	v := []time.Duration{time.Second, time.Hour}
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueArray(t *testing.T) {
	visitor := newMockArrayVisitor(t)
	v := []Value{Int(42), String("10")}
	value := Array(mockArray(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueNilArray(t *testing.T) {
	visitor := newMockArrayVisitor(t)
	value := Array(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, []Value(nil), visitor.value)
}

func TestValueAnyArray(t *testing.T) {
	visitor := newMockArrayVisitor(t)
	v := []Value{Int(42), String("10")}
	value := Any(mockArray(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyArray(t *testing.T) {
	visitor := newMockArrayVisitor(t)
	v := []Value{Int(42), String("10")}
	value := ConstAny(mockArray(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func BenchmarkValueArray(b *testing.B) {
	b.Run("Construction", func(b *testing.B) {
		v := []int{42, 43, 44}
		a := &efficientMockArray{v}
		var r Value
		for i := 0; i != b.N; i++ {
			r = Array(a)
		}
		avoidOptimization(r)
	})
	b.Run("Snapshot", func(b *testing.B) {
		v := []int{42, 43, 44}
		a := Array(&efficientMockArray{v})
		var r Value
		for i := 0; i != b.N; i++ {
			r = a
			Snapshot(&r)
		}
		avoidOptimization(r)
	})
	b.Run("ConstSnapshot", func(b *testing.B) {
		v := []int{42, 43, 44}
		a := ConstArray(&efficientMockArray{v})
		var r Value
		for i := 0; i != b.N; i++ {
			r = a
			Snapshot(&r)
		}
		avoidOptimization(r)
	})
	b.Run("CustomSnapshot", func(b *testing.B) {
		v := []int{42, 43, 44}
		a := Array(&efficientMockArrayWithSnapshot{efficientMockArray{v}})
		var r Value
		for i := 0; i != b.N; i++ {
			r = a
			Snapshot(&r)
		}
		avoidOptimization(r)
	})
}

func TestValueObject(t *testing.T) {
	visitor := newMockObjectVisitor(t)
	v := map[string]Value{"int": Int(42), "string": String("10")}
	value := Object(mockObject(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, len(v), visitor.count)
	require.Equal(t, v, visitor.value)
}

func TestValueNilObject(t *testing.T) {
	visitor := newMockObjectVisitor(t)
	value := Object(nil)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, 0, visitor.count)
	require.Equal(t, map[string]Value(nil), visitor.value)
}

func TestValueAnyObject(t *testing.T) {
	visitor := newMockObjectVisitor(t)
	v := map[string]Value{"int": Int(42), "string": String("10")}
	value := Any(mockObject(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, len(v), visitor.count)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyObject(t *testing.T) {
	visitor := newMockObjectVisitor(t)
	v := map[string]Value{"int": Int(42), "string": String("10")}
	value := ConstAny(mockObject(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, len(v), visitor.count)
	require.Equal(t, v, visitor.value)
}

func TestCorruptedValue(t *testing.T) {
	value := Value{bits: 255}
	visitor := ignoringVisitor{}
	require.Panics(t, func() {
		value.acceptVisitor(visitor)
	})
}

type customString string

func TestValueAnyCustomString(t *testing.T) {
	v := "test"
	visitor := newMockStringVisitor(t)
	value := Any(customString(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomString(t *testing.T) {
	v := "test"
	visitor := newMockStringVisitor(t)
	value := ConstAny(customString(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customBool bool

func TestValueAnyCustomBool(t *testing.T) {
	v := true
	visitor := newMockBoolVisitor(t)
	value := Any(customBool(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomBool(t *testing.T) {
	v := true
	visitor := newMockBoolVisitor(t)
	value := ConstAny(customBool(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customInt int

func TestValueAnyCustomInt(t *testing.T) {
	v := 42
	visitor := newMockIntVisitor(t)
	value := Any(customInt(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomInt(t *testing.T) {
	v := 42
	visitor := newMockIntVisitor(t)
	value := ConstAny(customInt(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customInt8 int8

func TestValueAnyCustomInt8(t *testing.T) {
	v := int8(42)
	visitor := newMockInt8Visitor(t)
	value := Any(customInt8(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomInt8(t *testing.T) {
	v := int8(42)
	visitor := newMockInt8Visitor(t)
	value := ConstAny(customInt8(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customInt16 int16

func TestValueAnyCustomInt16(t *testing.T) {
	v := int16(42)
	visitor := newMockInt16Visitor(t)
	value := Any(customInt16(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomInt16(t *testing.T) {
	v := int16(42)
	visitor := newMockInt16Visitor(t)
	value := ConstAny(customInt16(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customInt32 int32

func TestValueAnyCustomInt32(t *testing.T) {
	v := int32(42)
	visitor := newMockInt32Visitor(t)
	value := Any(customInt32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomInt32(t *testing.T) {
	v := int32(42)
	visitor := newMockInt32Visitor(t)
	value := ConstAny(customInt32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customInt64 int64

func TestValueAnyCustomInt64(t *testing.T) {
	v := int64(42)
	visitor := newMockInt64Visitor(t)
	value := Any(customInt64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomInt64(t *testing.T) {
	v := int64(42)
	visitor := newMockInt64Visitor(t)
	value := ConstAny(customInt64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customUint uint

func TestValueAnyCustomUint(t *testing.T) {
	v := uint(42)
	visitor := newMockUintVisitor(t)
	value := Any(customUint(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomUint(t *testing.T) {
	v := uint(42)
	visitor := newMockUintVisitor(t)
	value := ConstAny(customUint(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customUint8 uint8

func TestValueAnyCustomUint8(t *testing.T) {
	v := uint8(42)
	visitor := newMockUint8Visitor(t)
	value := Any(customUint8(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomUint8(t *testing.T) {
	v := uint8(42)
	visitor := newMockUint8Visitor(t)
	value := ConstAny(customUint8(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customUint16 uint16

func TestValueAnyCustomUint16(t *testing.T) {
	v := uint16(42)
	visitor := newMockUint16Visitor(t)
	value := Any(customUint16(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomUint16(t *testing.T) {
	v := uint16(42)
	visitor := newMockUint16Visitor(t)
	value := ConstAny(customUint16(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customUint32 uint32

func TestValueAnyCustomUint32(t *testing.T) {
	v := uint32(42)
	visitor := newMockUint32Visitor(t)
	value := Any(customUint32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomUint32(t *testing.T) {
	v := uint32(42)
	visitor := newMockUint32Visitor(t)
	value := ConstAny(customUint32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customUint64 uint64

func TestValueAnyCustomUint64(t *testing.T) {
	v := uint64(42)
	visitor := newMockUint64Visitor(t)
	value := Any(customUint64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomUint64(t *testing.T) {
	v := uint64(42)
	visitor := newMockUint64Visitor(t)
	value := ConstAny(customUint64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customFloat32 float32

func TestValueAnyCustomFloat32(t *testing.T) {
	v := float32(0.42)
	visitor := newMockFloat32Visitor(t)
	value := Any(customFloat32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomFloat32(t *testing.T) {
	v := float32(0.42)
	visitor := newMockFloat32Visitor(t)
	value := ConstAny(customFloat32(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type customFloat64 float64

func TestValueAnyCustomFloat64(t *testing.T) {
	v := float64(0.42)
	visitor := newMockFloat64Visitor(t)
	value := Any(customFloat64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyCustomFloat64(t *testing.T) {
	v := float64(0.42)
	visitor := newMockFloat64Visitor(t)
	value := ConstAny(customFloat64(v))
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

type emptyStruct struct{}

func TestValueAnyEmptyStruct(t *testing.T) {
	v := emptyStruct{}
	visitor := newMockAnyVisitor(t)
	value := Any(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func TestValueConstAnyEmptyStruct(t *testing.T) {
	v := emptyStruct{}
	visitor := newMockAnyVisitor(t)
	value := ConstAny(v)
	value.acceptVisitor(visitor)
	require.Equal(t, true, visitor.visited)
	require.Equal(t, v, visitor.value)
}

func avoidOptimization(v Value) {
	if v.Type() != TypeNone {
		testValuePlaceholder.Consume(v)
	}
}

type testValueConsumer interface {
	Consume(Value)
}

var testValuePlaceholder testValueConsumer = nilValueConsumer{}

type nilValueConsumer struct{}

func (c nilValueConsumer) Consume(Value) {}
