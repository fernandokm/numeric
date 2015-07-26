package numeric_test

import (
	. "github.com/fernandokm/numeric"
	"math"
	"testing"
)

var (
	floatBinaryTestInput = [][2]Float{
		{-5, -7},
		{6.1, -9},
		{-62, 9.7},
		{10.3, 34.9},
		{10.05, 15.2},
		{19, 18098.075},
	}
	posInf = math.Inf(1)
	negInf = math.Inf(-1)
	nan    = math.NaN()
)

func TestFloatLiterals(t *testing.T) {
	f0 := Float(0)
	f1 := Float(1)
	assertEquals(f0, f0.Zero(), t)
	assertEquals(f1, f0.One(), t)
}

func TestFloatOp_Negate(t *testing.T) {
	assertEquals(Float(-7), Float(7).Negate(), t)
	assertEquals(Float(9.7), Float(-9.7).Negate(), t)
}

func TestFloatOp_Add(t *testing.T) {
	for _, in := range floatBinaryTestInput {
		assertEquals(in[0]+in[1], in[0].Add(in[1]), t)
	}
}

func TestFloatOp_Subtract(t *testing.T) {
	for _, in := range floatBinaryTestInput {
		assertEquals(in[0]-in[1], in[0].Subtract(in[1]), t)
	}
}

func TestFloatOp_Multiply(t *testing.T) {
	for _, in := range floatBinaryTestInput {
		assertEquals(in[0]*in[1], in[0].Multiply(in[1]), t)
	}
}

func TestFloatOp_Divide(t *testing.T) {
	for _, in := range floatBinaryTestInput {
		assertEquals(in[0]/in[1], in[0].Divide(in[1]), t)
	}
}

func TestFloatConv_Float(t *testing.T) {
	assertEquals(15.0, Float(15.0).Float64(), t)
	assertEquals(-7.2, Float(-7.2).Float64(), t)
	assertEquals(27.9, Float(27.9).Float64(), t)
}

func TestFloatConv_BigRat(t *testing.T) {
	assertEquals(bigRat(15), Float(15).BigRat(), t)
	assertEquals(bigRat(-7.2), Float(-7.2).BigRat(), t)
	assertEquals(bigRat(27.9), Float(27.9).BigRat(), t)
}

func TestFloatComparison(t *testing.T) {
	msg := "Invalid float comparison "
	compare := func(x, y Float) int { return x.CompareTo(y) }
	assert(compare(10, 5) > 0, msg+"10>5", t)
	assert(compare(7, 7) == 0, msg+"7==7", t)
	assert(compare(-10, -5) < 0, msg+"-10<-5", t)
}

func TestFloatPromotion_ShouldPromote(t *testing.T) {
	assert(Float(posInf).ShouldPromote(), "Float(+infinity).ShouldPromote() should be true", t)
	assert(Float(negInf).ShouldPromote(), "Float(-infinity).ShouldPromote() should be true", t)

	assert(!Float(nan).ShouldPromote(), "Float(nan).ShouldPromote() should be false", t)
	assert(!Float(7.0).ShouldPromote(), "Float(7.0).ShouldPromote() should be false", t)
}

func TestFloatPromotion_Promote(t *testing.T) {
	assertEquals(bigFloat(7), Float(7).Promote(), t)
	assertEquals(bigFloat(-2.6), Float(-2.6).Promote(), t)
}
