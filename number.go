package numeric

import (
	"fmt"
	"math"
	"math/big"
)

// Number stores a numeric value
// using as its underlying storage a value implementing the Numeric interface.
type Number struct {
	storage Numeric
}

// Negative returns a Number that is the negative of n
// and automatically promotes it if required.
func (n Number) Negative() Number {
	storage := n.storage.Negate()
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Negate()
	}
	return Number{storage}
}

// TODO(fernandokm): avoid downcasting types in binary operations???
// eg: BigFloat.Add(Float) = BigFloat // Good (Float promoted to BigFloat)
//     Float.Add(BigFloat) = Float // Bad (BigFloat downcast to Float)

// Add returns a Number that is the sum of n and rhs
// and automatically promotes it if required.
func (n Number) Add(rhs Number) Number {
	storage := n.storage.Add(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Add(rhs.storage)
	}
	return Number{storage}
}

// Subtract returns a Number that is the difference of n and rhs
// and automatically promotes it if required.
func (n Number) Subtract(rhs Number) Number {
	storage := n.storage.Subtract(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Subtract(rhs.storage)
	}
	return Number{storage}
}

// Multiply returns a Number that is the product of n and rhs
// and automatically promotes it if required.
func (n Number) Multiply(rhs Number) Number {
	storage := n.storage.Multiply(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Multiply(rhs.storage)
	}
	return Number{storage}
}

// Divide returns a Number that is the quotient of n and rhs
// and automatically promotes it if required.
func (n Number) Divide(rhs Number) Number {
	storage := n.storage.Divide(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Divide(rhs.storage)
	}
	return Number{storage}
}

// CompareTo compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n Number) CompareTo(rhs Number) int {
	return n.storage.CompareTo(rhs.storage)
}

// Equals compares n to rhs
// and returns 1 if n is greater than rhs, -1 if it's smaller than rhs
// and 0 if the two values are equivalent.
func (n Number) Equals(rhs Number) bool {
	return n.CompareTo(rhs) == 0
}

// TODO(fernandokm): document behavior when overflow happens on conversions on Number and other Numeric types.
//                   document rounding behavior on Int(), Int64() and BigInt()

// Int returns an int with the value of n.
func (n Number) Int() int {
	return int(n.Int64())
}

// Int64 returns an int64 with the value of n.
func (n Number) Int64() int64 {
	return int64(n.storage.Float64())
}

// Float64 returns a float64 with the value of n.
func (n Number) Float64() float64 {
	return n.storage.Float64()
}

// BigInt returns a big.Int with the value of n.
func (n Number) BigInt() *big.Int {
	rat := n.BigRat()
	num, denom := rat.Num(), rat.Denom()
	return num.Div(num, denom)
}

// BigRat returns a big.Rat with the value of n.
func (n Number) BigRat() *big.Rat {
	return n.storage.BigRat()
}

// String returns a string representation of n.
func (n Number) String() string {
	return n.storage.String()
}

//TODO(fernandokm): disallow NaN and stop +/-inf and nan values from being created in Float

// NewNumberSafe attempts to convert value to a Number and returns an error if it failed.
func NewNumberSafe(value interface{}) (number Number, err error) {
	var n Numeric
	switch v := value.(type) {
	case int8:
		n = Float(v)
	case int16:
		n = Float(v)
	case int32:
		n = Float(v)
	case int64:
		n = Float(v)
	case int:
		n = Float(v)
	case uint8:
		n = Float(v)
	case uint16:
		n = Float(v)
	case uint32:
		n = Float(v)
	case uint64:
		n = Float(v)
	case uint:
		n = Float(v)
	case float32:
		n = Float(v)
	case float64:
		n = Float(v)

	case big.Rat:
		x := BigFloat(v)
		n = &x
	case big.Int:
		n = (*BigFloat)(new(big.Rat).SetInt(&v))

	case Numeric:
		n = v
	case Number:
		return v, nil
	default:
		return Number{}, fmt.Errorf("Unrecognized value %+v of type %T", v, v)
	}
	if n == Float(math.Inf(1)) || n == Float(math.Inf(-1)) {
		return Number{}, fmt.Errorf("Infinity is not a number!")
	}
	return Number{n}, nil
}

// NewNumber attempts to convert value to a Number and panics if it failed.
func NewNumber(value interface{}) Number {
	n, err := NewNumberSafe(value)
	if err != nil {
		panic(err.Error())
	}
	return n
}
