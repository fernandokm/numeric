package numeric

import (
	"math/big"
)

type Number struct {
	storage Numeric
}

func (n *Number) Negative() Number {
	storage := n.storage.Negate()
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Negate()
	}
	return Number{storage}
}

// TODO(fernandokm): avoid downcasting types in binary operations???
// eg: BigFloat.Add(Float) = BigFloat // Good (Float promoted to BigFloat)
//     Float.Add(BigFloat) = Float // Bad (BigFloat downcast to Float)

func (n *Number) Add(rhs Number) Number {
	storage := n.storage.Add(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Add(rhs.storage)
	}
	return Number{storage}
}

func (n *Number) Subtract(rhs Number) Number {
	storage := n.storage.Subtract(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Subtract(rhs.storage)
	}
	return Number{storage}
}

func (n *Number) Multiply(rhs Number) Number {
	storage := n.storage.Multiply(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Multiply(rhs.storage)
	}
	return Number{storage}
}

func (n *Number) Divide(rhs Number) Number {
	storage := n.storage.Divide(rhs.storage)
	for storage.ShouldPromote() {
		storage = n.storage.Promote().Divide(rhs.storage)
	}
	return Number{storage}
}

func (n *Number) CompareTo(rhs Number) int {
	return n.storage.CompareTo(rhs.storage)
}

func (n *Number) EqualTo(rhs Number) bool {
	return n.CompareTo(rhs) == 0
}

func (n *Number) Int() int {
	return int(n.Int64())
}

func (n *Number) Int64() int64 {
	return int64(n.storage.Float64())
}

func (n *Number) Float64() float64 {
	return n.storage.Float64()
}

func (n *Number) BigInt() *big.Int {
	rat := n.BigRat()
	num, denom := rat.Num(), rat.Denom()
	return num.Div(num, denom)
}

func (n *Number) BigRat() *big.Rat {
	return n.storage.BigRat()
}

func SafeFrom(value interface{}) (number Number, err error) {
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
		return Number{}, ConversionError{value}
	}
	return Number{n}, nil
}

func From(value interface{}) Number {
	n, err := SafeFrom(value)
	if err != nil {
		panic(err.Error())
	}
	return n
}
