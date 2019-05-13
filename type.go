package valf

// Type defines the value type stored in the Value.
type Type byte

// Valid values for Type.
const (
	TypeNone Type = iota
	TypeAny
	TypeBool
	TypeInt
	TypeInt8
	TypeInt16
	TypeInt32
	TypeInt64
	TypeUint
	TypeUint8
	TypeUint16
	TypeUint32
	TypeUint64
	TypeFloat32
	TypeFloat64
	TypeDuration
	TypeError
	TypeTime
	TypeString

	TypeBytes
	TypeBools
	TypeInts
	TypeInts8
	TypeInts16
	TypeInts32
	TypeInts64
	TypeUints
	TypeUints8
	TypeUints16
	TypeUints32
	TypeUints64
	TypeFloats32
	TypeFloats64
	TypeDurations
	TypeStrings

	TypeArray
	TypeObject
	TypeStringer
	TypeFormatter
)
