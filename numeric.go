package numeric

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
