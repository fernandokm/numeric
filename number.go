package numeric

import (
	"errors"
	"fmt"
	"math"
	"math/big"
)

// Number stores a numeric value
// using a Numeric type as its underlying storage.
// It can be created through the NewNumber and NewNumberSafe methods.
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

// Int64 returns an int64 with the value of n.
// If n cannot be represented as an int64, an error will be returned
// and the result of Int64() will be a sensible attempt to represent n.
func (n Number) Int64() (int64, error) {
	f, err := n.storage.Float64()
	val := int64(f)
	if f > math.MaxInt64 {
		val = math.MaxInt64
		err = errors.New("Value too large")
	} else if f < math.MinInt64 {
		val = math.MinInt64
		err = errors.New("Value too small")
	}
	return val, err
}

// Float64 returns a float64 with the value of n.
// If the value is invalid, too large or too small,
// the values NaN and +/-Infinity may be returned.
func (n Number) Float64() float64 {
	f, _ := n.storage.Float64()
	return f
}

// BigInt returns a *big.Int with the value of n.
func (n Number) BigInt() *big.Int {
	rat := n.BigRat()
	num, denom := rat.Num(), rat.Denom()
	return num.Quo(num, denom)
}

// BigRat returns a big.Rat with the value of n.
func (n Number) BigRat() *big.Rat {
	r, _ := n.storage.BigRat()
	return r
}

// String returns a string representation of n.
func (n Number) String() string {
	return n.storage.String()
}

// NewNumberSafe attempts to convert a value to a Number and returns an error if it fails.
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
	return Number{n}, nil
}

// NewNumber attempts to convert a value to a Number and panics if it fails.
func NewNumber(value interface{}) Number {
	n, err := NewNumberSafe(value)
	if err != nil {
		panic(err.Error())
	}
	return n
}
