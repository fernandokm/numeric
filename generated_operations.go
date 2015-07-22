package numeric
func (lhs Integer) Subtract(rhs Numeric) Numeric {
	return lhs - Integer(rhs.AsInt64())
}
func (lhs Integer) Multiply(rhs Numeric) Numeric {
	return lhs * Integer(rhs.AsInt64())
}
func (lhs Integer) Divide(rhs Numeric) Numeric {
	return lhs / Integer(rhs.AsInt64())
}
func (lhs Integer) Add(rhs Numeric) Numeric {
	return lhs + Integer(rhs.AsInt64())
}
func (lhs Float) Subtract(rhs Numeric) Numeric {
	return lhs - Float(rhs.AsFloat64())
}
func (lhs Float) Multiply(rhs Numeric) Numeric {
	return lhs * Float(rhs.AsFloat64())
}
func (lhs Float) Divide(rhs Numeric) Numeric {
	return lhs / Float(rhs.AsFloat64())
}
func (lhs Float) Add(rhs Numeric) Numeric {
	return lhs + Float(rhs.AsFloat64())
}
func (lhs UInteger) Subtract(rhs Numeric) Numeric {
	return lhs - UInteger(rhs.AsUInt64())
}
func (lhs UInteger) Multiply(rhs Numeric) Numeric {
	return lhs * UInteger(rhs.AsUInt64())
}
func (lhs UInteger) Divide(rhs Numeric) Numeric {
	return lhs / UInteger(rhs.AsUInt64())
}
func (lhs UInteger) Add(rhs Numeric) Numeric {
	return lhs + UInteger(rhs.AsUInt64())
}
