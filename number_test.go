package numeric_test

import (
	. "github.com/fernandokm/numeric"
	"math"
	"testing"
)

func TestNumberOp_Negate(t *testing.T) {
	assertEquals(NewNumber(15), NewNumber(-15.0).Negative(), t)
}

func TestNumberOp_Promotion(t *testing.T) {
	max := NewNumber(math.MaxFloat64).Multiply(NewNumber(2)).Divide(NewNumber(2)) // make sure max is promoted to BigFloat
	x := max.Multiply(max).Multiply(NewNumber(2))
	y := max.Divide(NewNumber(2))
	// result = (x + y) * (x - y) = x*x - y*y
	result := x.Add(y).Multiply(x.Subtract(y))
	expected := x.Multiply(x).Subtract(y.Multiply(y))
	assertEquals(expected, result, t)
	assert(!result.Multiply(x).Equals(expected), "Expected non equal numbers", t)
}

func TestNumberOp_Add(t *testing.T) {
	assertEquals(NewNumber(17), NewNumber(10).Add(NewNumber(7.0)), t)
	assertEquals(NewNumber(25), NewNumber(100).Add(NewNumber(-75)), t)
}

func TestNumberOp_Subtract(t *testing.T) {
	assertEquals(NewNumber(17), NewNumber(10).Subtract(NewNumber(-7.0)), t)
	assertEquals(NewNumber(25), NewNumber(100).Subtract(NewNumber(75)), t)
}

func TestNumberOp_Multiply(t *testing.T) {
	assertEquals(NewNumber(20), NewNumber(4).Multiply(NewNumber(5)), t)
	assertEquals(NewNumber(25), NewNumber(-5).Multiply(NewNumber(-5.0)), t)
	assertEquals(NewNumber(-50), NewNumber(-5).Multiply(NewNumber(10)), t)
}

func TestNumberOp_Divide(t *testing.T) {
	assertEquals(NewNumber(20), NewNumber(4).Divide(NewNumber(1).Divide(NewNumber(5))), t)
	assertEquals(NewNumber(5), NewNumber(25).Divide(NewNumber(5.0)), t)
	assertEquals(NewNumber(-10), NewNumber(100).Divide(NewNumber(-10)), t)
}
