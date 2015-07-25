package numeric

import (
	"math"
	"math/big"
)

func (n Float) Zero() Numeric {
	return Float(0)
}

func (n Float) One() Numeric {
	return Float(1)
}

func (n Float) Negate() Numeric {
	return -n
}

func (n Float) Add(rhs Numeric) Numeric {
	return n + Float(rhs.Float64())
}

func (n Float) Subtract(rhs Numeric) Numeric {
	return n - Float(rhs.Float64())
}

func (n Float) Multiply(rhs Numeric) Numeric {
	return n * Float(rhs.Float64())
}

func (n Float) Divide(rhs Numeric) Numeric {
	return n / Float(rhs.Float64())
}

func (n Float) Float64() float64 {
	return float64(n)
}

func (n Float) BigRat() *big.Rat {
	return new(big.Rat).SetFloat64(float64(n))
}

func (n Float) CompareTo(rhs Numeric) int {
	r := Float(rhs.Float64())
	if n > r {
		return 1
	} else if n < r {
		return -1
	}
	return 0
}

func (n Float) ShouldPromote() bool {
	return math.IsInf(float64(n), 0)
}

func (n Float) Promote() Numeric {
	return (*BigFloat)(n.BigRat())
}
