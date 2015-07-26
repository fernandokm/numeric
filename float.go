package numeric

import (
	"math"
	"math/big"
	"strconv"
)

// Zero returns a Float with value 0.
func (n Float) Zero() Numeric {
	return Float(0)
}

// One returns a Float with value 1.
func (n Float) One() Numeric {
	return Float(1)
}

// Negate returns a Float that is the negative of n.
func (n Float) Negate() Numeric {
	return -n
}

// Add returns a Float that is the sum of n and rhs.
func (n Float) Add(rhs Numeric) Numeric {
	return n + Float(rhs.Float64())
}

// Subtract returns a Float that is the difference of n and rhs.
func (n Float) Subtract(rhs Numeric) Numeric {
	return n - Float(rhs.Float64())
}

// Multiply returns a Float that is the product of n and rhs.
func (n Float) Multiply(rhs Numeric) Numeric {
	return n * Float(rhs.Float64())
}

// Divide returns a Float that is the quotient of n and rhs.
func (n Float) Divide(rhs Numeric) Numeric {
	return n / Float(rhs.Float64())
}

// Float64 returns a copy of the underlying value of this Float.
func (n Float) Float64() float64 {
	return float64(n)
}

// BigRat converts this float to a big.Rat value and returns it.
func (n Float) BigRat() *big.Rat {
	return new(big.Rat).SetFloat64(float64(n))
}

// CompareTo compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n Float) CompareTo(rhs Numeric) int {
	r := Float(rhs.Float64())
	if n > r {
		return 1
	} else if n < r {
		return -1
	}
	return 0
}

// ShouldPromote returns true only if n is +/-Infinity.
func (n Float) ShouldPromote() bool {
	return math.IsInf(float64(n), 0)
}

// Promote returns a BigFloat with the same value as n.
func (n Float) Promote() Numeric {
	return (*BigFloat)(n.BigRat())
}

// String returns a string representation of this Float.
func (n Float) String() string {
	return strconv.FormatFloat(float64(n), 'g', -1, 64)
}
