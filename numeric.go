package numeric

import "reflect"

type Numeric interface {
	Negate() Numeric
	Add(rhs Numeric) Numeric
	Subtract(rhs Numeric) Numeric
	Multiply(rhs Numeric) Numeric
	Divide(rhs Numeric) Numeric
	AsInt64() int64
	AsUInt64() uint64
	AsFloat64() float64
	Zero() Numeric
	One() Numeric
	Compare(rhs Numeric) int
}

type Integer int64
type UInteger uint64
type Float float64

func compareFloat(x, y float64) int {
	if x > y {
		return 1
	}
	if x < y {
		return -1
	}
	return 0
}

func compareInt(x, y int64) int {
	if x > y {
		return 1
	}
	if x < y {
		return -1
	}
	return 0
}

func compareUInt(x, y uint64) int {
	if x > y {
		return 1
	}
	if x < y {
		return -1
	}
	return 0
}

func (n Integer) Compare(rhs Numeric) int {
	if reflect.TypeOf(rhs) == floatType {
		return compareFloat(n.AsFloat64(), rhs.AsFloat64())
	}
	return compareInt(n.AsInt64(), rhs.AsInt64())
}

func (n UInteger) Compare(rhs Numeric) int {
	t := reflect.TypeOf(rhs)
	if t == floatType {
		return compareFloat(n.AsFloat64(), rhs.AsFloat64())
	}
	if t == integerType && rhs.AsInt64() > 0 {
		return -1
	}
	return compareUInt(n.AsUInt64(), rhs.AsUInt64())
}

func (n Float) Compare(rhs Numeric) int {
	return compareFloat(n.AsFloat64(), rhs.AsFloat64())
}
