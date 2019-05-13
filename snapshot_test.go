package valf

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type snapshotTestSpec struct {
	Name           string
	Generate       func() (input, golden Value, modify func())
	ShoudPanic     bool
	SkipValueCheck bool
	ExtraCheck     func(t *testing.T, actual, golden *Value)
}

var snapshotTests = []snapshotTestSpec{
	{
		Name: "Bool",
		Generate: func() (Value, Value, func()) {
			return Bool(true), Bool(true), func() {}
		},
	},
	{
		Name: "Int",
		Generate: func() (Value, Value, func()) {
			return Int(42), Int(42), func() {}
		},
	},
	{
		Name: "Int64",
		Generate: func() (Value, Value, func()) {
			return Int64(42), Int64(42), func() {}
		},
	},
	{
		Name: "Int32",
		Generate: func() (Value, Value, func()) {
			return Int32(42), Int32(42), func() {}
		},
	},
	{
		Name: "Int16",
		Generate: func() (Value, Value, func()) {
			return Int16(42), Int16(42), func() {}
		},
	},
	{
		Name: "Int8",
		Generate: func() (Value, Value, func()) {
			return Int8(42), Int8(42), func() {}
		},
	},
	{
		Name: "Uint",
		Generate: func() (Value, Value, func()) {
			return Uint(42), Uint(42), func() {}
		},
	},
	{
		Name: "Uint64",
		Generate: func() (Value, Value, func()) {
			return Uint64(42), Uint64(42), func() {}
		},
	},
	{
		Name: "Uint32",
		Generate: func() (Value, Value, func()) {
			return Uint32(42), Uint32(42), func() {}
		},
	},
	{
		Name: "Uint16",
		Generate: func() (Value, Value, func()) {
			return Uint16(42), Uint16(42), func() {}
		},
	},
	{
		Name: "Uint8",
		Generate: func() (Value, Value, func()) {
			return Uint8(42), Uint8(42), func() {}
		},
	},
	{
		Name: "Float64",
		Generate: func() (Value, Value, func()) {
			return Float64(0.42), Float64(0.42), func() {}
		},
	},
	{
		Name: "Float32",
		Generate: func() (Value, Value, func()) {
			return Float32(0.42), Float32(0.42), func() {}
		},
	},
	{
		Name: "Duration",
		Generate: func() (Value, Value, func()) {
			return Duration(time.Second), Duration(time.Second), func() {}
		},
	},
	{
		Name: "Bytes",
		Generate: func() (Value, Value, func()) {
			v := []byte("test")
			return Bytes(v), ConstBytes([]byte("test")), func() {
				v[1] = 'o'
			}
		},
	},
	{
		Name: "String",
		Generate: func() (Value, Value, func()) {
			return String("test"), String("test"), func() {}
		},
	},
	{
		Name: "Strings",
		Generate: func() (Value, Value, func()) {
			v := []string{"a", "b"}
			return Strings(v), ConstStrings([]string{"a", "b"}), func() {
				v[1] = "c"
			}
		},
	},
	{
		Name: "ConstStrings",
		Generate: func() (Value, Value, func()) {
			return ConstStrings([]string{"a", "b"}), ConstStrings([]string{"a", "b"}), func() {}
		},
	},
	{
		Name: "Bools",
		Generate: func() (Value, Value, func()) {
			v := []bool{false, true}
			return Bools(v), ConstBools([]bool{false, true}), func() {
				v[1] = false
			}
		},
	},
	{
		Name: "Ints",
		Generate: func() (Value, Value, func()) {
			v := []int{1, 2, 3}
			return Ints(v), ConstInts([]int{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Ints8",
		Generate: func() (Value, Value, func()) {
			v := []int8{1, 2, 3}
			return Ints8(v), ConstInts8([]int8{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Ints16",
		Generate: func() (Value, Value, func()) {
			v := []int16{1, 2, 3}
			return Ints16(v), ConstInts16([]int16{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Ints32",
		Generate: func() (Value, Value, func()) {
			v := []int32{1, 2, 3}
			return Ints32(v), ConstInts32([]int32{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Ints64",
		Generate: func() (Value, Value, func()) {
			v := []int64{1, 2, 3}
			return Ints64(v), ConstInts64([]int64{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Uints",
		Generate: func() (Value, Value, func()) {
			v := []uint{1, 2, 3}
			return Uints(v), ConstUints([]uint{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Uints8",
		Generate: func() (Value, Value, func()) {
			v := []uint8{1, 2, 3}
			return Uints8(v), ConstUints8([]uint8{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Uints16",
		Generate: func() (Value, Value, func()) {
			v := []uint16{1, 2, 3}
			return Uints16(v), ConstUints16([]uint16{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Uints32",
		Generate: func() (Value, Value, func()) {
			v := []uint32{1, 2, 3}
			return Uints32(v), ConstUints32([]uint32{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Uints64",
		Generate: func() (Value, Value, func()) {
			v := []uint64{1, 2, 3}
			return Uints64(v), ConstUints64([]uint64{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "ConstBools",
		Generate: func() (Value, Value, func()) {
			return ConstBools([]bool{false, true}), ConstBools([]bool{false, true}), func() {}
		},
	},
	{
		Name: "ConstInts",
		Generate: func() (Value, Value, func()) {
			return ConstInts([]int{1, 2, 3}), ConstInts([]int{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstInts8",
		Generate: func() (Value, Value, func()) {
			return ConstInts8([]int8{1, 2, 3}), ConstInts8([]int8{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstInts16",
		Generate: func() (Value, Value, func()) {
			return ConstInts16([]int16{1, 2, 3}), ConstInts16([]int16{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstInts32",
		Generate: func() (Value, Value, func()) {
			return ConstInts32([]int32{1, 2, 3}), ConstInts32([]int32{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstInts64",
		Generate: func() (Value, Value, func()) {
			return ConstInts64([]int64{1, 2, 3}), ConstInts64([]int64{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstUints",
		Generate: func() (Value, Value, func()) {
			return ConstUints([]uint{1, 2, 3}), ConstUints([]uint{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstUints8",
		Generate: func() (Value, Value, func()) {
			return ConstUints8([]uint8{1, 2, 3}), ConstUints8([]uint8{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstUints16",
		Generate: func() (Value, Value, func()) {
			return ConstUints16([]uint16{1, 2, 3}), ConstUints16([]uint16{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstUints32",
		Generate: func() (Value, Value, func()) {
			return ConstUints32([]uint32{1, 2, 3}), ConstUints32([]uint32{1, 2, 3}), func() {}
		},
	},
	{
		Name: "ConstUints64",
		Generate: func() (Value, Value, func()) {
			return ConstUints64([]uint64{1, 2, 3}), ConstUints64([]uint64{1, 2, 3}), func() {}
		},
	},
	{
		Name: "Floats32",
		Generate: func() (Value, Value, func()) {
			v := []float32{0.42, -0.42}
			return Floats32(v), ConstFloats32([]float32{0.42, -0.42}), func() {
				v[1] = 0.21
			}
		},
	},
	{
		Name: "ConstFloats32",
		Generate: func() (Value, Value, func()) {
			return ConstFloats32([]float32{0.42, -0.42}), ConstFloats32([]float32{0.42, -0.42}), func() {}
		},
	},
	{
		Name: "Floats64",
		Generate: func() (Value, Value, func()) {
			v := []float64{0.42, -0.42}
			return Floats64(v), ConstFloats64([]float64{0.42, -0.42}), func() {
				v[1] = 0.21
			}
		},
	},
	{
		Name: "ConstFloats64",
		Generate: func() (Value, Value, func()) {
			return ConstFloats64([]float64{0.42, -0.42}), ConstFloats64([]float64{0.42, -0.42}), func() {}
		},
	},
	{
		Name: "Durations",
		Generate: func() (Value, Value, func()) {
			v := []time.Duration{time.Second, time.Hour}
			return Durations(v), ConstDurations([]time.Duration{time.Second, time.Hour}), func() {
				v[1] = time.Millisecond
			}
		},
	},
	{
		Name: "ConstDurations",
		Generate: func() (Value, Value, func()) {
			v := []time.Duration{time.Second, time.Hour}
			return ConstDurations(v), ConstDurations([]time.Duration{time.Second, time.Millisecond}), func() {
				v[1] = time.Millisecond
			}
		},
	},
	{
		Name: "Error",
		Generate: func() (Value, Value, func()) {
			return Error(context.Canceled), Error(context.Canceled), func() {}
		},
	},
	{
		Name: "Time",
		Generate: func() (Value, Value, func()) {
			return Time(time.Unix(42, 42)), Time(time.Unix(42, 42)), func() {}
		},
	},
	{
		Name: "Stringer",
		Generate: func() (Value, Value, func()) {
			v := new(testMutableStringer)
			*v = "test"
			return Stringer(v), String("test"), func() {
				*v = "other"
			}
		},
	},
	{
		Name: "ConstStringer",
		Generate: func() (Value, Value, func()) {
			v := new(testMutableStringer)
			*v = "test"
			return ConstStringer(v), ConstStringer(v), func() {
				v := new(testMutableStringer)
				*v = "test"
			}
		},
	},
	{
		Name: "Formatter",
		Generate: func() (Value, Value, func()) {
			v := new(int)
			*v = 42
			return FormatterRepr(v), String(fmt.Sprintf("%#v", v)), func() {
				*v = 21
			}
		},
	},
	{
		Name: "ConstFormatter",
		Generate: func() (Value, Value, func()) {
			v := new(int)
			*v = 42
			return ConstFormatterRepr(v), ConstFormatterRepr(v), func() {
				*v = 21
			}
		},
	},
	{
		Name: "Any",
		Generate: func() (Value, Value, func()) {
			v := new(emptyStruct)
			return Any(v), ConstAny(v), func() {
				v = new(emptyStruct)
			}
		},
		ShoudPanic: true,
	},
	{
		Name: "ConstAny",
		Generate: func() (Value, Value, func()) {
			v := new(emptyStruct)
			return ConstAny(v), ConstAny(v), func() {
				v = new(emptyStruct)
			}
		},
	},
	{
		Name: "AnyNil",
		Generate: func() (Value, Value, func()) {
			return Any(nil), ConstAny(nil), func() {}
		},
	},
	{
		Name: "AnySnapshotter",
		Generate: func() (Value, Value, func()) {
			v := testSnapshotter([]int{1, 2, 3})
			return Any(v), ConstAny([]int{1, 2, 3}), func() {
				v[1] = 4
			}
		},
	},
	{
		Name: "Array",
		Generate: func() (Value, Value, func()) {
			v := mockArray{Int(42), String("test")}
			return Array(v), ConstArray(arraySnapshot{items: []Value{Int(42), String("test")}}), func() {
				v[1] = String("other")
			}
		},
		ExtraCheck: func(t *testing.T, actual, golden *Value) {
			visitor := newMockArrayVisitor(t)
			actual.AcceptVisitor(visitor)
			require.Equal(t, true, visitor.visited)
			require.Equal(t, 2, len(visitor.value))
			require.Equal(t, Int(42), visitor.value[0])
			require.Equal(t, String("test"), visitor.value[1])
		},
	},
	{
		Name: "ConstArray",
		Generate: func() (Value, Value, func()) {
			v := mockArray{Int(42), String("test")}
			return ConstArray(v), ConstArray(mockArray{Int(42), String("other")}), func() {
				v[1] = String("other")
			}
		},
	},
	{
		Name: "ArrayNil",
		Generate: func() (Value, Value, func()) {
			return Array(nil), ConstArray(nil), func() {}
		},
	},
	{
		Name: "Object",
		Generate: func() (Value, Value, func()) {
			v := mockObject{"int": Int(42), "string": String("test")}
			return Object(v), Value{}, func() {
				v["string"] = String("other")
			}
		},
		ExtraCheck: func(t *testing.T, actual, golden *Value) {
			visitor := newMockObjectVisitor(t)
			actual.AcceptVisitor(visitor)
			require.Equal(t, true, visitor.visited)
			require.Equal(t, 2, len(visitor.value))
			require.Equal(t, Int(42), visitor.value["int"])
			require.Equal(t, String("test"), visitor.value["string"])
		},
		SkipValueCheck: true,
	},
	{
		Name: "ConstObject",
		Generate: func() (Value, Value, func()) {
			v := mockObject{"int": Int(42), "string": String("test")}
			golden := ConstObject(mockObject{"int": Int(42), "string": String("other")})
			return ConstObject(v), golden, func() {
				v["string"] = String("other")
			}
		},
	},
	{
		Name: "ObjectNil",
		Generate: func() (Value, Value, func()) {
			return Object(nil), ConstObject(nil), func() {}
		},
	},
	{
		Name:       "CorruptedValue",
		ShoudPanic: true,
		Generate: func() (Value, Value, func()) {
			return Value{bits: 255 & bitsMaskType}, Value{}, func() {}
		},
	},
}

type testMutableStringer string

func (s *testMutableStringer) String() string {
	return string(*s)
}

type testSnapshotter []int

func (s testSnapshotter) TakeSnapshot() interface{} {
	cc := make([]int, len(s))
	copy(cc, s)

	return cc
}

func TestSnapshotValues(t *testing.T) {
	for _, test := range snapshotTests {
		t.Run(test.Name, func(t *testing.T) {
			value, golden, modify := test.Generate()
			if test.ShoudPanic {
				require.Panics(t, func() {
					value.Snapshot()
				})
			} else {
				s := value.Snapshot()
				require.Equal(t, true, s.bits.Const())
				modify()
				if !test.SkipValueCheck {
					require.Equal(t, golden, s)
				}
				if test.ExtraCheck != nil {
					test.ExtraCheck(t, &s, &golden)
				}
			}
		})
	}
}
