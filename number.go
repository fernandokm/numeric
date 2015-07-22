package numeric

type Number struct {
	storage Numeric
}

func SafeFrom(value interface{}) (number Number, err error) {
	var n Numeric
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
	case Number:
		return v, nil
	default:
		err = ConversionError{value}
	}
	if n != nil {
		number = Number{n}
	}
	return
}

func From(value interface{}) Number {
	n, err := SafeFrom(value)
	if err != nil {
		panic(err.Error())
	}
	return n
}
