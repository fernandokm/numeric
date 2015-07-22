package numeric
func (n Integer) Negate() Numeric {
	return -n
}
func (n Integer) Zero() Numeric {
	return Integer(0)
}
func (n Integer) One() Numeric {
	return Integer(1)
}
func (n Float) Negate() Numeric {
	return -n
}
func (n Float) Zero() Numeric {
	return Float(0)
}
func (n Float) One() Numeric {
	return Float(1)
}
func (n UInteger) Negate() Numeric {
	return -n
}
func (n UInteger) Zero() Numeric {
	return UInteger(0)
}
func (n UInteger) One() Numeric {
	return UInteger(1)
}
