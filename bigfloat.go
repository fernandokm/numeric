package numeric

import (
	"math/big"
)

// Zero returns a *BigFloat with the value of 0.
func (n *BigFloat) Zero() Numeric {
	return (*BigFloat)(big.NewRat(0, 1))
}

// One returns a *BigFloat with the value of 1.
func (n *BigFloat) One() Numeric {
	return (*BigFloat)(big.NewRat(1, 1))
}

// Negate returns a *BigFloat that is the negative of n.
func (n *BigFloat) Negate() Numeric {
	rat := new(big.Rat)
	return (*BigFloat)(rat.Neg((*big.Rat)(n)))
}

// Add returns a *BigFloat that is the sum of n and rhs.
func (n *BigFloat) Add(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Add(rat, (*big.Rat)(n)))
}

// Subtract returns a *BigFloat that is the difference of n and rhs.
func (n *BigFloat) Subtract(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Sub((*big.Rat)(n), rat))
}

// Multiply  returns a *BigFloat that is the product of n and rhs.
func (n *BigFloat) Multiply(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Mul(rat, (*big.Rat)(n)))
}

// Divide returns a *BigFloat that is the quotient of n and rhs.
func (n *BigFloat) Divide(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Quo((*big.Rat)(n), rat))
}

// Float64 converts this BigFloat to a float64 value and returns it.
func (n *BigFloat) Float64() float64 {
	val, _ := ((*big.Rat)(n)).Float64()
	return val
}

// BigRat returns a copy of the underlying value of this BigFloat.
func (n *BigFloat) BigRat() *big.Rat {
	// rat := new(big.Rat)
	// *rat = big.Rat(*n)
	// return rat
	return new(big.Rat).Set((*big.Rat)(n))
}

// CompareTo compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n *BigFloat) CompareTo(rhs Numeric) int {
	r := rhs.BigRat()
	return r.Sub((*big.Rat)(n), r).Sign()
}

// ShouldPromote returns false for *BigFloat.
func (n *BigFloat) ShouldPromote() bool {
	return false
}

// Promote returns a copy of n.
func (n *BigFloat) Promote() Numeric {
	return (*BigFloat)(n.BigRat())
}

// String returns a string representation of this BigFloat.
func (n *BigFloat) String() string {
	// TODO(fernandokm): improve this
	return (*big.Rat)(n).FloatString(5)
}
