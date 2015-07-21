package numeric

import "fmt"

type Numeric interface {
	Add(rhs Numeric) Numeric
	Subtract(rhs Numeric) Numeric
	Multiply(rhs Numeric) Numeric
	Divide(rhs Numeric) Numeric
	AsInt64() int64
	AsUInt64() uint64
	AsFloat64() float64
}

type Integer int64
type UInteger uint64
type Float float64

type ConversionError struct {
	value interface{}
}

func (e ConversionError) Error() string {
	return fmt.Sprintf("Unrecognized value %v of type %T", e.value, e.value)
}

func SafeFrom(value interface{}) (n Numeric, err error) {
	switch v := value.(type) {
	case int8:
		n = Integer(v)
	case int16:
		n = Integer(v)
	case int32:
		n = Integer(v)
	case int64:
		n = Integer(v)
	case int:
		n = Integer(v)

	case uint8:
		n = UInteger(v)
	case uint16:
		n = UInteger(v)
	case uint32:
		n = UInteger(v)
	case uint64:
		n = UInteger(v)
	case uint:
		n = UInteger(v)

	case float32:
		n = Float(v)
	case float64:
		n = Float(v)
	case Numeric:
		n = v
	default:
		err = ConversionError{value}
	}
	return
}

func From(value interface{}) Numeric {
	n, err := SafeFrom(value)
	if err != nil {
		panic(err.Error())
	}
	return n
}
