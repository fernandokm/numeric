package numeric

import "math"

func Sqrt(n Number) Number {
	return Number{Float(math.Sqrt(n.AsFloat64()))}
}
