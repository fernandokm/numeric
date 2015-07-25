package numeric

import (
	"math/big"
)

type Number struct {
	storage Numeric
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
