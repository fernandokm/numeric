package numeric

import (
	"math/big"
)

//TODO: change Zero() and One methods to New(int32)

// Numeric is an interface that represents any numeric type,
// such as float64. Default implementations exist for the types
// float64 (type Float) and big.Rat(type BigFloat).
//
// Methods Zero() and One() should return the values 0 and 1.
// The Negate(), Add(), Subtract(), Multiply() and Divide() methods
// are aliases for the corresponding operations.
// To compare two Numeric values, the method CompareTo() should be called.
// For inputs x and y, it should return:
//		x.CompareTo(y)
//      // 0 if x == y
//			// 1 if x > y
//			// -1 if x < y
// The available conversion methods are Float64() and BigRat().
// Numeric also supports string conversion through the Stringer interface.
//
// Promote() and ShouldPromote() are part of a value promotion system.
// If a value is to big or small to be adequately represented in its
// Numeric type, it can be promoted (converted) to a type that has a
// larger capacity. If this happens, the method ShouldPromote() should
// return true and the method Promote() should return the promoted
// value.
type Numeric interface {
	Zero() Numeric
	One() Numeric
	Negate() Numeric
	Add(rhs Numeric) Numeric
	Subtract(rhs Numeric) Numeric
	Multiply(rhs Numeric) Numeric
	Divide(rhs Numeric) Numeric
	Float64() float64
	BigRat() *big.Rat
	CompareTo(rhs Numeric) int
	ShouldPromote() bool
	Promote() Numeric
	String() string
}

// Float is an implementation of the Numeric interface for the float64 type.
// It can be promoted to a BigFloat
// and will return true on ShouldPromote() if it's equal to +/-Infinity.
type Float float64

// BigFloat is an implementation of the Numeric interface for the big.Rat type.
// It cannot be promoted.
type BigFloat big.Rat
