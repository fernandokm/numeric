package numeric

import "reflect"

var (
	integerType  = reflect.TypeOf(Integer(0))
	uintegerType = reflect.TypeOf(UInteger(0))
	floatType    = reflect.TypeOf(Float(0))
)

func order(v1, v2 Number) (Number, Number) {
	t1, t2 := reflect.TypeOf(v1.storage), reflect.TypeOf(v2.storage)
	if (t2 == floatType) || (t1 == uintegerType && t2 == integerType) {
		return v2, v1
	}
	return v1, v2
}

func (n Number) Compare(value Number) int {
	return n.storage.Compare(value.storage)
}

func (n Number) GreaterThan(value Number) bool {
	return n.Compare(value) > 0
}

func (n Number) GreaterThanOrEqualTo(value Number) bool {
	return n.Compare(value) >= 0
}

func (n Number) LessThan(value Number) bool {
	return n.Compare(value) < 0
}

func (n Number) LessThanOrEqualTo(value Number) bool {
	return n.Compare(value) <= 0
}

func (n Number) EqualTo(value Number) bool {
	return n.Compare(value) == 0
}

func (n Number) Negate() {
	n.storage = n.storage.Negate()
}

func (n Number) Negative() Number {
	return Number{n.storage.Negate()}
}

func (n Number) Add(value Number) {
	n, value = order(n, value)
	n.storage = n.storage.Add(value.storage)
}

func (n Number) Plus(value Number) Number {
	n, value = order(n, value)
	return Number{n.storage.Add(value.storage)}
}

func (n Number) Subtract(value Number) {
	n, value = order(n, value)
	n.storage = n.storage.Subtract(value.storage)
}

func (n Number) Minus(value Number) Number {
	n, value = order(n, value)
	return Number{n.storage.Subtract(value.storage)}
}

func (n Number) Multiply(value Number) {
	n, value = order(n, value)
	n.storage = n.storage.Multiply(value.storage)
}

func (n Number) Times(value Number) Number {
	n, value = order(n, value)
	return Number{n.storage.Multiply(value.storage)}
}

func (n Number) Divide(value Number) {
	n, value = order(n, value)
	n.storage = n.storage.Divide(value.storage)
}

func (n Number) Divided(value Number) Number {
	n, value = order(n, value)
	return Number{n.storage.Divide(value.storage)}
}

func (n Number) AsInt64() int64 {
	return n.storage.AsInt64()
}

func (n Number) AsInt32() int32 {
	return int32(n.storage.AsInt64())
}

func (n Number) AsInt16() int16 {
	return int16(n.storage.AsInt64())
}

func (n Number) AsInt8() int8 {
	return int8(n.storage.AsInt64())
}

func (n Number) AsUInt64() uint64 {
	return n.storage.AsUInt64()
}

func (n Number) AsUInt32() uint32 {
	return uint32(n.storage.AsUInt64())
}

func (n Number) AsUInt16() uint16 {
	return uint16(n.storage.AsUInt64())
}

func (n Number) AsUInt8() uint8 {
	return uint8(n.storage.AsUInt64())
}

func (n Number) AsFloat32() float32 {
	return float32(n.storage.AsFloat64())
}

func (n Number) AsFloat64() float64 {
	return n.storage.AsFloat64()
}
