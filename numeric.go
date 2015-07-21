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

func (e *ConversionError) Error() string {
	return fmt.Sprintf("Unrecognized value %v of type %T", e.value, e.value)
}

func from(value interface{}) (n Numeric, err ConversionError) {
	switch v := value.(type) {
	case int8, int16, int32, int64:
		n = v.(Integer)
	case uint8, uint16, uint32, uint64:
		n = v.(UInteger)
	case float32, float64:
		n = v.(Float)
	case Numeric:
		n = v
	default:
		err = ConversionError{value}
	}
	return
}
