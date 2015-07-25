package numeric

import (
	"math/big"
)

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
}

type Float float64
type BigFloat big.Rat
