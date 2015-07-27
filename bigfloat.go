package numeric

import (
	"errors"
	"math"
	"math/big"
	"strings"
)

// New returns a *BigFloat equivalent to value.
func (n *BigFloat) New(value int64) Numeric {
	return (*BigFloat)(big.NewRat(value, 1))
}

// Negate returns a *BigFloat that is the negative of n.
func (n *BigFloat) Negate() Numeric {
	rat := new(big.Rat)
	return (*BigFloat)(rat.Neg((*big.Rat)(n)))
}

// Add returns a *BigFloat that is the sum of n and rhs.
func (n *BigFloat) Add(rhs Numeric) Numeric {
	rat, _ := rhs.BigRat()
	return (*BigFloat)(rat.Add(rat, (*big.Rat)(n)))
}

// Subtract returns a *BigFloat that is the difference of n and rhs.
func (n *BigFloat) Subtract(rhs Numeric) Numeric {
	rat, _ := rhs.BigRat()
	return (*BigFloat)(rat.Sub((*big.Rat)(n), rat))
}

// Multiply  returns a *BigFloat that is the product of n and rhs.
func (n *BigFloat) Multiply(rhs Numeric) Numeric {
	rat, _ := rhs.BigRat()
	return (*BigFloat)(rat.Mul(rat, (*big.Rat)(n)))
}

// Divide returns a *BigFloat that is the quotient of n and rhs.
func (n *BigFloat) Divide(rhs Numeric) Numeric {
	rat, _ := rhs.BigRat()
	return (*BigFloat)(rat.Quo((*big.Rat)(n), rat))
}

// Float64 converts this BigFloat to a float64 value and returns it.
// If the value is too large or too small, an error will be returned and
// the result of the conversion will be +/- infinity, respectively.
func (n *BigFloat) Float64() (float64, error) {
	val, _ := ((*big.Rat)(n)).Float64()
	if math.IsInf(val, 0) {
		return val, errors.New("Value is infinity")
	}
	return val, nil
}

// BigRat returns a copy of the underlying value of this BigFloat.
func (n *BigFloat) BigRat() (*big.Rat, error) {
	return new(big.Rat).Set((*big.Rat)(n)), nil
}

// CompareTo compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n *BigFloat) CompareTo(rhs Numeric) int {
	r, _ := rhs.BigRat()
	return r.Sub((*big.Rat)(n), r).Sign()
}

// ShouldPromote returns false for *BigFloat.
func (n *BigFloat) ShouldPromote() bool {
	return false
}

// Promote returns a copy of n.
func (n *BigFloat) Promote() Numeric {
	rat, _ := n.BigRat()
	return (*BigFloat)(rat)
}

// String returns a string representation of this BigFloat.
func (n *BigFloat) String() string {
	rat := (*big.Rat)(n)
	if rat.IsInt() {
		return rat.FloatString(0)
	}
	return strings.TrimRight(rat.FloatString(5), "0")
}
