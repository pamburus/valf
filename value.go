package valf

import (
	"fmt"
	"math"
	"reflect"
	"time"
	"unsafe"
)

// Value holds data of a specific type.
type Value struct {
	bits    bits
	vAny    interface{}
	vInt    int64
	vBytes  []byte
	vString string
}

// Type returns type of the value stored in v.
func (v Value) Type() Type {
	return v.bits.Type()
}

// Const returns true if the value stored in v is immutable.
func (v Value) Const() bool {
	return v.bits.Const()
}

// Snapshot returns a Value which can be safely stored for a long with guarantee that it won't be modified.
// The data of the value are copied if needed to achieve that guarantee.
func (v Value) Snapshot() Value {
	Snapshot(&v)

	return v
}

// AcceptVisitor interprets Value data according to its type and calls appropriate
// Visitor method.
func (v Value) AcceptVisitor(visitor Visitor) {
	switch v.bits.Type() {
	case TypeNone:
		visitor.VisitNone()
	case TypeAny:
		visitor.VisitAny(v.vAny)
	case TypeBool:
		visitor.VisitBool(v.vInt != 0)
	case TypeInt:
		visitor.VisitInt(int(v.vInt))
	case TypeInt8:
		visitor.VisitInt8(int8(v.vInt))
	case TypeInt16:
		visitor.VisitInt16(int16(v.vInt))
	case TypeInt32:
		visitor.VisitInt32(int32(v.vInt))
	case TypeInt64:
		visitor.VisitInt64(v.vInt)
	case TypeUint:
		visitor.VisitUint(uint(v.vInt))
	case TypeUint8:
		visitor.VisitUint8(uint8(v.vInt))
	case TypeUint16:
		visitor.VisitUint16(uint16(v.vInt))
	case TypeUint32:
		visitor.VisitUint32(uint32(v.vInt))
	case TypeUint64:
		visitor.VisitUint64(uint64(v.vInt))
	case TypeFloat32:
		visitor.VisitFloat32(math.Float32frombits(uint32(v.vInt)))
	case TypeFloat64:
		visitor.VisitFloat64(math.Float64frombits(uint64(v.vInt)))
	case TypeDuration:
		visitor.VisitDuration(time.Duration(v.vInt))
	case TypeError:
		if v.vAny != nil {
			visitor.VisitError(v.vAny.(error))
		} else {
			visitor.VisitError(nil)
		}
	case TypeTime:
		visitor.VisitTime(time.Unix(0, v.vInt).In(v.vAny.(*time.Location)))
	case TypeArray:
		if v.vAny != nil {
			visitor.VisitArray(v.vAny.(ValueArray))
		} else {
			visitor.VisitArray(nil)
		}
	case TypeObject:
		if v.vAny != nil {
			visitor.VisitObject(v.vAny.(ValueObject))
		} else {
			visitor.VisitObject(nil)
		}
	case TypeStringer:
		if v.vAny != nil {
			visitor.VisitString(v.vAny.(fmt.Stringer).String())
		} else {
			visitor.VisitAny(nil)
		}
	case TypeFormatter:
		visitor.VisitString(fmt.Sprintf(v.vString, v.vAny))
	case TypeBytes:
		visitor.VisitBytes(v.vBytes)
	case TypeString:
		visitor.VisitString(v.vString)
	case TypeStrings:
		visitor.VisitStrings(v.vAny.([]string))
	case TypeBools:
		visitor.VisitBools(*(*[]bool)(unsafe.Pointer(&v.vBytes)))
	case TypeInts:
		visitor.VisitInts(*(*[]int)(unsafe.Pointer(&v.vBytes)))
	case TypeInts8:
		visitor.VisitInts8(*(*[]int8)(unsafe.Pointer(&v.vBytes)))
	case TypeInts16:
		visitor.VisitInts16(*(*[]int16)(unsafe.Pointer(&v.vBytes)))
	case TypeInts32:
		visitor.VisitInts32(*(*[]int32)(unsafe.Pointer(&v.vBytes)))
	case TypeInts64:
		visitor.VisitInts64(*(*[]int64)(unsafe.Pointer(&v.vBytes)))
	case TypeUints:
		visitor.VisitUints(*(*[]uint)(unsafe.Pointer(&v.vBytes)))
	case TypeUints8:
		visitor.VisitUints8(*(*[]uint8)(unsafe.Pointer(&v.vBytes)))
	case TypeUints16:
		visitor.VisitUints16(*(*[]uint16)(unsafe.Pointer(&v.vBytes)))
	case TypeUints32:
		visitor.VisitUints32(*(*[]uint32)(unsafe.Pointer(&v.vBytes)))
	case TypeUints64:
		visitor.VisitUints64(*(*[]uint64)(unsafe.Pointer(&v.vBytes)))
	case TypeFloats32:
		visitor.VisitFloats32(*(*[]float32)(unsafe.Pointer(&v.vBytes)))
	case TypeFloats64:
		visitor.VisitFloats64(*(*[]float64)(unsafe.Pointer(&v.vBytes)))
	case TypeDurations:
		visitor.VisitDurations(*(*[]time.Duration)(unsafe.Pointer(&v.vBytes)))

	default:
		panic(fmt.Errorf("snapf: internal error: unhandled value type: %v", v.bits.Type()))
	}
}

// Bool returns a new Value with the given bool.
func Bool(v bool) Value {
	var tmp int64
	if v {
		tmp = 1
	}

	return Value{bits: bits(TypeBool) | bitsConst, vInt: tmp}
}

// Int returns a new Value with the given int.
func Int(v int) Value {
	return Value{bits: bits(TypeInt) | bitsConst, vInt: int64(v)}
}

// Int64 returns a new Value with the given int64.
func Int64(v int64) Value {
	return Value{bits: bits(TypeInt64) | bitsConst, vInt: v}
}

// Int32 returns a new Value with the given int32.
func Int32(v int32) Value {
	return Value{bits: bits(TypeInt32) | bitsConst, vInt: int64(v)}
}

// Int16 returns a new Value with the given int16.
func Int16(v int16) Value {
	return Value{bits: bits(TypeInt16) | bitsConst, vInt: int64(v)}
}

// Int8 returns a new Value with the given int.8
func Int8(v int8) Value {
	return Value{bits: bits(TypeInt8) | bitsConst, vInt: int64(v)}
}

// Uint returns a new Value with the given uint.
func Uint(v uint) Value {
	return Value{bits: bits(TypeUint) | bitsConst, vInt: int64(v)}
}

// Uint64 returns a new Value with the given uint64.
func Uint64(v uint64) Value {
	return Value{bits: bits(TypeUint64) | bitsConst, vInt: int64(v)}
}

// Uint32 returns a new Value with the given uint32.
func Uint32(v uint32) Value {
	return Value{bits: bits(TypeUint32) | bitsConst, vInt: int64(v)}
}

// Uint16 returns a new Value with the given uint16.
func Uint16(v uint16) Value {
	return Value{bits: bits(TypeUint16) | bitsConst, vInt: int64(v)}
}

// Uint8 returns a new Value with the given uint8.
func Uint8(v uint8) Value {
	return Value{bits: bits(TypeUint8) | bitsConst, vInt: int64(v)}
}

// Float64 returns a new Value with the given float64.
func Float64(v float64) Value {
	return Value{bits: bits(TypeFloat64) | bitsConst, vInt: int64(math.Float64bits(v))}
}

// Float32 returns a new Value with the given float32.
func Float32(v float32) Value {
	return Value{bits: bits(TypeFloat32) | bitsConst, vInt: int64(math.Float32bits(v))}
}

// Duration returns a new Value with the given time.Duration.
func Duration(v time.Duration) Value {
	return Value{bits: bits(TypeDuration) | bitsConst, vInt: int64(v)}
}

// Bytes returns a new Value with the given slice of bytes.
func Bytes(v []byte) Value {
	return Value{bits: bits(TypeBytes), vBytes: v}
}

// String returns a new Value with the given string.
func String(v string) Value {
	return Value{bits: bits(TypeString) | bitsConst, vString: v}
}

// Strings returns a new Value with the given slice of strings.
func Strings(v []string) Value {
	return Value{bits: bits(TypeStrings), vAny: v}
}

// ConstStrings returns a new Value with the given slice of strings.
func ConstStrings(v []string) Value {
	return Value{bits: bits(TypeStrings) | bitsConst, vAny: v}
}

// Bools returns a new Value with the given slice of bools.
func Bools(v []bool) Value {
	return Value{bits: bits(TypeBools), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Ints returns a new Value with the given slice of ints.
func Ints(v []int) Value {
	return Value{bits: bits(TypeInts), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Ints8 returns a new Value with the given slice of 8-bit ints.
func Ints8(v []int8) Value {
	return Value{bits: bits(TypeInts8), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Ints16 returns a new Value with the given slice of 16-bit ints.
func Ints16(v []int16) Value {
	return Value{bits: bits(TypeInts16), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Ints32 returns a new Value with the given slice of 32-bit ints.
func Ints32(v []int32) Value {
	return Value{bits: bits(TypeInts32), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Ints64 returns a new Value with the given slice of 64-bit ints.
func Ints64(v []int64) Value {
	return Value{bits: bits(TypeInts64), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Uints returns a new Value with the given slice of uints.
func Uints(v []uint) Value {
	return Value{bits: bits(TypeUints), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Uints8 returns a new Value with the given slice of 8-bit uints.
func Uints8(v []uint8) Value {
	return Value{bits: bits(TypeUints8), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Uints16 returns a new Value with the given slice of 16-bit uints.
func Uints16(v []uint16) Value {
	return Value{bits: bits(TypeUints16), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Uints32 returns a new Value with the given slice of 32-bit uints.
func Uints32(v []uint32) Value {
	return Value{bits: bits(TypeUints32), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Uints64 returns a new Value with the given slice of 64-bit uints.
func Uints64(v []uint64) Value {
	return Value{bits: bits(TypeUints64), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Floats32 returns a new Value with the given slice of 32-bit floats.
func Floats32(v []float32) Value {
	return Value{bits: bits(TypeFloats32), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Floats64 returns a new Value with the given slice of 64-biy floats.
func Floats64(v []float64) Value {
	return Value{bits: bits(TypeFloats64), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Durations returns a new Value with the given slice of time.Duration.
func Durations(v []time.Duration) Value {
	return Value{bits: bits(TypeDurations), vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstBytes returns a new Value with the given slice of bytes.
//
// Call ConstBytes if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstBytes(v []byte) Value {
	return Value{bits: bits(TypeBytes) | bitsConst, vBytes: v}
}

// ConstBools returns a new Value with the given slice of bools.
//
// Call ConstBools if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstBools(v []bool) Value {
	return Value{bits: bits(TypeBools) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstInts returns a new Value with the given slice of ints.
//
// Call ConstInts if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts(v []int) Value {
	return Value{bits: bits(TypeInts) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstInts8 returns a new Value with the given slice of 8-bit ints.
//
// Call ConstInts8 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts8(v []int8) Value {
	return Value{bits: bits(TypeInts8) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstInts16 returns a new Value with the given slice of 16-bit ints.
//
// Call ConstInts16 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts16(v []int16) Value {
	return Value{bits: bits(TypeInts16) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstInts32 returns a new Value with the given slice of 32-bit ints.
//
// Call ConstInts32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts32(v []int32) Value {
	return Value{bits: bits(TypeInts32) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstInts64 returns a new Value with the given slice of 64-bit ints.
//
// Call ConstInts64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstInts64(v []int64) Value {
	return Value{bits: bits(TypeInts64) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstUints returns a new Value with the given slice of uints.
//
// Call ConstUints if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints(v []uint) Value {
	return Value{bits: bits(TypeUints) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstUints8 returns a new Value with the given slice of 8-bit uints.
//
// Call ConstUints8 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints8(v []uint8) Value {
	return Value{bits: bits(TypeUints8) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstUints16 returns a new Value with the given slice of 16-bit uints.
//
// Call ConstUints16 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints16(v []uint16) Value {
	return Value{bits: bits(TypeUints16) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstUints32 returns a new Value with the given slice of 32-bit uints.
//
// Call ConstUints32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints32(v []uint32) Value {
	return Value{bits: bits(TypeUints32) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstUints64 returns a new Value with the given slice of 64-bit uints.
//
// Call ConstUints64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstUints64(v []uint64) Value {
	return Value{bits: bits(TypeUints64) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstFloats32 returns a new Value with the given slice of 32-bit floats.
//
// Call ConstFloats32 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstFloats32(v []float32) Value {
	return Value{bits: bits(TypeFloats32) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstFloats64 returns a new Value with the given slice of 64-bit floats.
//
// Call ConstFloats64 if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstFloats64(v []float64) Value {
	return Value{bits: bits(TypeFloats64) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// ConstDurations returns a new Value with the given slice of time.Duration.
//
// Call ConstDurations if your array is const. It has significantly less impact
// on the calling goroutine.
//
func ConstDurations(v []time.Duration) Value {
	return Value{bits: bits(TypeDurations) | bitsConst, vBytes: *(*[]byte)(unsafe.Pointer(&v))}
}

// Error returns a new Value with the given error.
func Error(v error) Value {
	return Value{bits: bits(TypeError) | bitsConst, vAny: v}
}

// Time returns a new Value with the given time.Time.
func Time(v time.Time) Value {
	return Value{bits: bits(TypeTime) | bitsConst, vInt: v.UnixNano(), vAny: v.Location()}
}

// Array returns a new Value with the given ArrayEncoder.
func Array(v ValueArray) Value {
	if v == nil {
		return ConstArray(v)
	}

	return Value{bits: bits(TypeArray), vAny: v}
}

// ConstArray returns a new Value with the given ArrayEncoder.
func ConstArray(v ValueArray) Value {
	return Value{bits: bits(TypeArray) | bitsConst, vAny: v}
}

// Object returns a new Value with the given ObjectEncoder.
func Object(v ValueObject) Value {
	if v == nil {
		return ConstObject(v)
	}

	return Value{bits: bits(TypeObject), vAny: v}
}

// ConstObject returns a new Value with the given ObjectEncoder.
func ConstObject(v ValueObject) Value {
	return Value{bits: bits(TypeObject) | bitsConst, vAny: v}
}

// Stringer returns a new Value with the given Stringer.
func Stringer(v fmt.Stringer) Value {
	if v == nil {
		return ConstStringer(v)
	}

	return Value{bits: bits(TypeStringer), vAny: v}
}

// ConstStringer returns a new Value with the given Stringer.
// Call ConstStringer if your object is const. It has significantly less
// impact on the calling goroutine.
func ConstStringer(v fmt.Stringer) Value {
	return Value{bits: bits(TypeStringer) | bitsConst, vAny: v}
}

// Formatter returns a new Value with the given verb and interface to
// Valueformat.
func Formatter(verb string, v interface{}) Value {
	return Value{bits: bits(TypeFormatter), vString: verb, vAny: v}
}

// FormatterRepr returns a new Value with the given interface to format.
// ValueIt uses the predefined verb "%#v" (a Go-syntax representation of the value).
func FormatterRepr(v interface{}) Value {
	return Formatter("%#v", v)
}

// ConstFormatter returns a new Value with the given verb and interface
// Valueto format.
//
// Call ConstFormatter if your object is const. It has significantly less
// impact on the calling goroutine.
//
func ConstFormatter(verb string, v interface{}) Value {
	return Value{bits: bits(TypeFormatter) | bitsConst, vString: verb, vAny: v}
}

// ConstFormatterRepr returns a new Value with the given interface to
// format. It uses the predefined verb "%#v" (a Go-syntax representation of
// Valuethe value).
//
// Call ConstFormatterV if your object is const. It has significantly less
// impact on the calling goroutine.
//
func ConstFormatterRepr(v interface{}) Value {
	return ConstFormatter("%#v", v)
}

// Any returns a new Value with the given value of any type. It tries
// to choose the best way to represent value a Value.
//
// Note that Any is not able to choose ConstX methods. Use specific
// functions for better performance.
func Any(v interface{}) Value {
	if v == nil {
		return Value{bits: bits(TypeAny) | bitsConst}
	}

	switch rv := v.(type) {
	case bool:
		return Bool(rv)
	case int:
		return Int(rv)
	case int64:
		return Int64(rv)
	case int32:
		return Int32(rv)
	case int16:
		return Int16(rv)
	case int8:
		return Int8(rv)
	case uint:
		return Uint(rv)
	case uint64:
		return Uint64(rv)
	case uint32:
		return Uint32(rv)
	case uint16:
		return Uint16(rv)
	case uint8:
		return Uint8(rv)
	case float64:
		return Float64(rv)
	case float32:
		return Float32(rv)
	case time.Time:
		return Time(rv)
	case time.Duration:
		return Duration(rv)
	case error:
		return Error(rv)
	case ValueArray:
		return Array(rv)
	case ValueObject:
		return Object(rv)
	case []byte:
		return Bytes(rv)
	case []string:
		return Strings(rv)
	case []bool:
		return Bools(rv)
	case []int:
		return Ints(rv)
	case []int64:
		return Ints64(rv)
	case []int32:
		return Ints32(rv)
	case []int16:
		return Ints16(rv)
	case []int8:
		return Ints8(rv)
	case []uint:
		return Uints(rv)
	case []uint64:
		return Uints64(rv)
	case []uint32:
		return Uints32(rv)
	case []uint16:
		return Uints16(rv)
	case []float64:
		return Floats64(rv)
	case []float32:
		return Floats32(rv)
	case []time.Duration:
		return Durations(rv)
	case string:
		return String(rv)
	case fmt.Stringer:
		return Stringer(rv)

	default:
		switch reflect.TypeOf(rv).Kind() {
		case reflect.String:
			return String(reflect.ValueOf(rv).String())
		case reflect.Bool:
			return Bool(reflect.ValueOf(rv).Bool())
		case reflect.Int:
			return Int(int(reflect.ValueOf(rv).Int()))
		case reflect.Int8:
			return Int8(int8(reflect.ValueOf(rv).Int()))
		case reflect.Int16:
			return Int16(int16(reflect.ValueOf(rv).Int()))
		case reflect.Int32:
			return Int32(int32(reflect.ValueOf(rv).Int()))
		case reflect.Int64:
			return Int64(int64(reflect.ValueOf(rv).Int()))
		case reflect.Uint:
			return Uint(uint(reflect.ValueOf(rv).Uint()))
		case reflect.Uint8:
			return Uint8(uint8(reflect.ValueOf(rv).Uint()))
		case reflect.Uint16:
			return Uint16(uint16(reflect.ValueOf(rv).Uint()))
		case reflect.Uint32:
			return Uint32(uint32(reflect.ValueOf(rv).Uint()))
		case reflect.Uint64:
			return Uint64(uint64(reflect.ValueOf(rv).Uint()))
		case reflect.Float32:
			return Float32(float32(reflect.ValueOf(rv).Float()))
		case reflect.Float64:
			return Float64(float64(reflect.ValueOf(rv).Float()))
		}
	}

	return Value{bits: bits(TypeAny), vAny: v}
}

// ConstAny returns a new Value with the given value of any type. It tries
// to choose the best way to represent value a Value assuming that
// provided value will is immutable and won't change in the future.
func ConstAny(v interface{}) Value {
	if v == nil {
		return Value{bits: bits(TypeAny) | bitsConst}
	}

	switch rv := v.(type) {
	case bool:
		return Bool(rv)
	case int:
		return Int(rv)
	case int64:
		return Int64(rv)
	case int32:
		return Int32(rv)
	case int16:
		return Int16(rv)
	case int8:
		return Int8(rv)
	case uint:
		return Uint(rv)
	case uint64:
		return Uint64(rv)
	case uint32:
		return Uint32(rv)
	case uint16:
		return Uint16(rv)
	case uint8:
		return Uint8(rv)
	case float64:
		return Float64(rv)
	case float32:
		return Float32(rv)
	case time.Time:
		return Time(rv)
	case time.Duration:
		return Duration(rv)
	case error:
		return Error(rv)
	case ValueArray:
		return ConstArray(rv)
	case ValueObject:
		return ConstObject(rv)
	case []byte:
		return ConstBytes(rv)
	case []string:
		return ConstStrings(rv)
	case []bool:
		return ConstBools(rv)
	case []int:
		return ConstInts(rv)
	case []int64:
		return ConstInts64(rv)
	case []int32:
		return ConstInts32(rv)
	case []int16:
		return ConstInts16(rv)
	case []int8:
		return ConstInts8(rv)
	case []uint:
		return ConstUints(rv)
	case []uint64:
		return ConstUints64(rv)
	case []uint32:
		return ConstUints32(rv)
	case []uint16:
		return ConstUints16(rv)
	case []float64:
		return ConstFloats64(rv)
	case []float32:
		return ConstFloats32(rv)
	case []time.Duration:
		return ConstDurations(rv)
	case string:
		return String(rv)
	case fmt.Stringer:
		return ConstStringer(rv)

	default:
		switch reflect.TypeOf(rv).Kind() {
		case reflect.String:
			return String(reflect.ValueOf(rv).String())
		case reflect.Bool:
			return Bool(reflect.ValueOf(rv).Bool())
		case reflect.Int:
			return Int(int(reflect.ValueOf(rv).Int()))
		case reflect.Int8:
			return Int8(int8(reflect.ValueOf(rv).Int()))
		case reflect.Int16:
			return Int16(int16(reflect.ValueOf(rv).Int()))
		case reflect.Int32:
			return Int32(int32(reflect.ValueOf(rv).Int()))
		case reflect.Int64:
			return Int64(int64(reflect.ValueOf(rv).Int()))
		case reflect.Uint:
			return Uint(uint(reflect.ValueOf(rv).Uint()))
		case reflect.Uint8:
			return Uint8(uint8(reflect.ValueOf(rv).Uint()))
		case reflect.Uint16:
			return Uint16(uint16(reflect.ValueOf(rv).Uint()))
		case reflect.Uint32:
			return Uint32(uint32(reflect.ValueOf(rv).Uint()))
		case reflect.Uint64:
			return Uint64(uint64(reflect.ValueOf(rv).Uint()))
		case reflect.Float32:
			return Float32(float32(reflect.ValueOf(rv).Float()))
		case reflect.Float64:
			return Float64(float64(reflect.ValueOf(rv).Float()))
		}
	}

	return Value{bits: bits(TypeAny) | bitsConst, vAny: v}
}

type bits byte

const (
	bitsMaskType bits = (1 << 7) - 1
	bitsConst    bits = 1 << 7
)

func (b bits) Type() Type {
	return Type(b & bitsMaskType)
}

func (b bits) Const() bool {
	return b&bitsConst != 0
}
