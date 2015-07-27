package numeric

import (
	"errors"
	"math"
	"math/big"
	"strconv"
)

var (
	posInfAsBigFloat = (*BigFloat)(new(big.Rat).Add(new(big.Rat).SetFloat64(math.MaxFloat64), new(big.Rat).SetFloat64(math.SmallestNonzeroFloat64)))
	negInfAsBigFloat = (*BigFloat)(new(big.Rat).Neg((*big.Rat)(posInfAsBigFloat)))
)

// New returns a Float equivalent to value.
func (n Float) New(value int64) Numeric {
	return Float(value)
}

// Negate returns a Float that is the negative of n.
func (n Float) Negate() Numeric {
	return -n
}

// Add returns a Float that is the sum of n and rhs.
func (n Float) Add(rhs Numeric) Numeric {
	f, _ := rhs.Float64()
	return n + Float(f)
}

// Subtract returns a Float that is the difference of n and rhs.
func (n Float) Subtract(rhs Numeric) Numeric {
	f, _ := rhs.Float64()
	return n - Float(f)
}

// Multiply returns a Float that is the product of n and rhs.
func (n Float) Multiply(rhs Numeric) Numeric {
	f, _ := rhs.Float64()
	return n * Float(f)
}

// Divide returns a Float that is the quotient of n and rhs.
func (n Float) Divide(rhs Numeric) Numeric {
	f, _ := rhs.Float64()
	return n / Float(f)
}

// Float64 returns a copy of the underlying value of this Float.
func (n Float) Float64() (float64, error) {
	return float64(n), nil
}

// BigRat converts this float to a big.Rat value and returns it.
// If the value is +/-Infinity or NaN, an error will be returned.
// For +Infinity, the result will be set to
// math.MaxFloat64 + math.SmallestNonzeroFloat64
// and for -Infinity, the negative of that will be returned.
// If n is NaN, the result will be 0.
func (n Float) BigRat() (*big.Rat, error) {
	if math.IsInf(float64(n), 0) {
		return new(big.Rat), errors.New("Cannot convert infinity to big.Rat")
	}
	if math.IsNaN(float64(n)) {
		return new(big.Rat), errors.New("Cannot convert NaN to big.Rat")
	}
	return new(big.Rat).SetFloat64(float64(n)), nil
}

// CompareTo compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n Float) CompareTo(rhs Numeric) int {
	f, _ := rhs.Float64()
	r := Float(f)
	if n > r {
		return 1
	} else if n < r {
		return -1
	}
	return 0
}

// ShouldPromote returns true only if n is +/-Infinity or NaN.
func (n Float) ShouldPromote() bool {
	return math.IsInf(float64(n), 0) || math.IsNaN(float64(n))
}

// Promote returns a BigFloat with the same value as n.
func (n Float) Promote() Numeric {
	if math.IsNaN(float64(n)) {
		return (*BigFloat)(big.NewRat(0, 1))
	}
	if math.IsInf(float64(n), 1) {
		return posInfAsBigFloat
	}
	if math.IsInf(float64(n), -1) {
		return negInfAsBigFloat
	}
	rat, _ := n.BigRat()
	return (*BigFloat)(rat)
}

// String returns a string representation of this Float.
func (n Float) String() string {
	return strconv.FormatFloat(float64(n), 'g', -1, 64)
}
