package numeric_test

import (
	. "github.com/fernandokm/numeric"
	"math"
	"math/big"
	"testing"
)

func TestNumberOp_Negate(t *testing.T) {
	assertEquals(NewNumber(15), NewNumber(-15.0).Negative(), t)
}

func TestNumberOp_Promotion(t *testing.T) {
	max := NewNumber(math.MaxFloat64).Multiply(NewNumber(2)).Divide(NewNumber(2))
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

func TestNumberComparison(t *testing.T) {
	compare := func(x, y interface{}) int {
		return NewNumber(x).CompareTo(NewNumber(y))
	}

	assert(compare(50, 70) == -1, "50 >= 70", t)
	assert(compare(21, 21.50) == -1, "21 >= 21.5", t)
	assert(compare(50.5, 50.5) == 0, "50.5 != 50.5", t)
	assert(compare(78.6, 70.2) == 1, "78.6 <= 70.2", t)
}

func TestNumberEquality(t *testing.T) {
	assert(NewNumber(20).Equals(NewNumber(20.0)), "20 != 20", t)
	assert(NewNumber(51.01).Equals(NewNumber(51.01)), "51.01 != 51.01", t)
}

func TestNumberConv_Int64(t *testing.T) {
	x, err := NewNumber(50.0).Int64()
	assert(err == nil, "Expected nil error", t)
	assertEquals(x, int64(50), t)

	x, err = NewNumber(math.Inf(1)).Int64()
	assert(err != nil, "Expected error with +infinity float", t)
	assert(x == math.MaxInt64, "Expected +infinity to be converted to math.MaxInt64", t)

	x, err = NewNumber(math.Inf(-1)).Int64()
	assert(err != nil, "Expected error with -infinity float", t)
	assert(x == math.MinInt64, "Expected -infinity to be converted to math.MinInt64", t)
}

func TestNumberConv_Float64(t *testing.T) {
	assertEquals(NewNumber(50).Float64(), 50.0, t)
	max := NewNumber(math.MaxFloat64)
	large := max.Add(max)
	small := large.Negative()
	assert(math.IsInf(large.Float64(), 1), "Expected +inf", t)
	assert(math.IsInf(small.Float64(), -1), "Expected -inf", t)
}

func TestNumberConv_BigInt(t *testing.T) {
	assert(new(big.Int).Sub(NewNumber(10).BigInt(), big.NewInt(10)).Sign() == 0, "10 != 10", t)
}

func TestNumberConv_BigRat(t *testing.T) {
	assert(new(big.Rat).Sub(bigRat(10), NewNumber(10).BigRat()).Sign() == 0, "10 != 10", t)
}

func TestNumberNewSafe(t *testing.T) {
	n, err := NewNumberSafe(10)
	assert(n.Float64() == 10, "10 != 10", t)
	assert(err == nil, "Expected nil error", t)

	_, err = NewNumberSafe("hey")
	assert(err != nil, "Expected error", t)
}
