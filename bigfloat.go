package numeric

import (
	"math/big"
)

func (n *BigFloat) Zero() Numeric {
	return (*BigFloat)(big.NewRat(0, 1))
}

func (n *BigFloat) One() Numeric {
	return (*BigFloat)(big.NewRat(1, 1))
}

func (n *BigFloat) Negate() Numeric {
	rat := new(big.Rat)
	return (*BigFloat)(rat.Neg((*big.Rat)(n)))
}

func (n *BigFloat) Add(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Add(rat, (*big.Rat)(n)))
}

func (n *BigFloat) Subtract(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Sub(rat, (*big.Rat)(n)))
}

func (n *BigFloat) Multiply(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Mul(rat, (*big.Rat)(n)))
}

func (n *BigFloat) Divide(rhs Numeric) Numeric {
	rat := rhs.BigRat()
	return (*BigFloat)(rat.Quo(rat, (*big.Rat)(n)))
}

func (n *BigFloat) Float64() float64 {
	val, _ := ((*big.Rat)(n)).Float64()
	return val
}

func (n *BigFloat) BigRat() *big.Rat {
	var rat *big.Rat
	*rat = *n.BigRat()
	return rat
}

func (n *BigFloat) Compare(rhs Numeric) int {
	r := rhs.BigRat()
	return r.Sub((*big.Rat)(n), r).Sign()
}

func (n *BigFloat) ShouldPromote() bool {
	return false
}

func (n *BigFloat) Promote() Numeric {
	return (*BigFloat)(n.BigRat())
}
