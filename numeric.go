package numeric

import (
	"math/big"
)

// Numeric is an interface that represents any numeric type,
// such as float64. Default implementations exist for the types
// float64 (type Float) and big.Rat(type BigFloat).
//
// Method New() should return the value provided.
// The Negate(), Add(), Subtract(), Multiply() and Divide() methods
// are aliases for the corresponding operations.
// To compare two Numeric values, the method CompareTo() should be called.
// For inputs x and y, it should return:
//		x.CompareTo(y)
//			// 0 if x == y
//			// 1 if x > y
//			// -1 if x < y
// The available conversion methods are Float64() and BigRat().
// Both of these methods return the result and an error, which should represent
// overflow or invalid conversion errors. If an error occurs, the methods
// should attempt to return a sensible value, but may fallback to the default
// zero value.
// Numeric also supports string conversion through the Stringer interface.
//
// Promote() and ShouldPromote() are part of a value promotion system.
// If a value is to big or small to be adequately represented in its
// Numeric type, it can be promoted (converted) to a type that has a
// larger capacity. If this happens, the method ShouldPromote() should
// return true and the method Promote() should return the promoted
// value.
type Numeric interface {
	New(value int64) Numeric
	Negate() Numeric
	Add(rhs Numeric) Numeric
	Subtract(rhs Numeric) Numeric
	Multiply(rhs Numeric) Numeric
	Divide(rhs Numeric) Numeric
	Float64() (float64, error)
	BigRat() (*big.Rat, error)
	CompareTo(rhs Numeric) int
	ShouldPromote() bool
	Promote() Numeric
	String() string
}

// Float is an implementation of the Numeric interface for the float64 type.
// It can be promoted to a BigFloat
// and will return true on ShouldPromote() if it's equal to +/-Infinity or NaN.
type Float float64

// BigFloat is an implementation of the Numeric interface for the big.Rat type.
// It cannot be promoted.
type BigFloat big.Rat
