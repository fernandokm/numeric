package numeric_test

import (
	// "fmt"
	. "github.com/fernandokm/numeric"
	"math/big"
	"testing"
)

var (
	bigFloatBinaryTestInput = [][2]*BigFloat{
		{bigFloat(-5), bigFloat(-7)},
		{bigFloat(6.1), bigFloat(-9)},
		{bigFloat(-62), bigFloat(9.7)},
		{bigFloat(10.3), bigFloat(34.9)},
		{bigFloat(10.05), bigFloat(15.2)},
		{bigFloat(19), bigFloat(18098.075)},
	}
)

func addBigFloat(x, y *BigFloat) *BigFloat {
	v1, v2 := (*big.Rat)(x), (*big.Rat)(y)
	return (*BigFloat)(new(big.Rat).Add(v1, v2))
}

func subBigFloat(x, y *BigFloat) *BigFloat {
	v1, v2 := (*big.Rat)(x), (*big.Rat)(y)
	// fmt.Printf("%s - %s = %s\n", v1, v2, new(big.Rat).Sub(v1, v2))
	// fmt.Printf("%v=%v\t\t%v=%v\n", x, v1, y, v2)
	return (*BigFloat)(new(big.Rat).Sub(v1, v2))
}

func mulBigFloat(x, y *BigFloat) *BigFloat {
	v1, v2 := (*big.Rat)(x), (*big.Rat)(y)
	return (*BigFloat)(new(big.Rat).Mul(v1, v2))
}

func divBigFloat(x, y *BigFloat) *BigFloat {
	v1, v2 := (*big.Rat)(x), (*big.Rat)(y)
	return (*BigFloat)(new(big.Rat).Quo(v1, v2))
}

func TestBigFloatLiterals(t *testing.T) {
	f0 := bigFloat(0)
	f1 := bigFloat(1)
	assertEquals(f0, f0.Zero(), t)
	assertEquals(f1, f0.One(), t)
}

func TestBigFloatOp_Negate(t *testing.T) {
	assertEquals(bigFloat(-7), bigFloat(7).Negate(), t)
	assertEquals(bigFloat(9.7), bigFloat(-9.7).Negate(), t)
}

func TestBigFloatOp_Add(t *testing.T) {
	for _, in := range bigFloatBinaryTestInput {
		// t.Logf("Values: %v\t%v", in[0], in[1])
		assertEquals(addBigFloat(in[0], in[1]), in[0].Add(in[1]), t)
	}
}

func TestBigFloatOp_Subtract(t *testing.T) {
	for _, in := range bigFloatBinaryTestInput {
		// t.Logf("Values: %v\t%v", in[0], in[1])
		assertEquals(subBigFloat(in[0], in[1]), in[0].Subtract(in[1]), t)
	}
}

func TestBigFloatOp_Multiply(t *testing.T) {
	for _, in := range bigFloatBinaryTestInput {
		assertEquals(mulBigFloat(in[0], in[1]), in[0].Multiply(in[1]), t)
	}
}

func TestBigFloatOp_Divide(t *testing.T) {
	for _, in := range bigFloatBinaryTestInput {
		assertEquals(divBigFloat(in[0], in[1]), in[0].Divide(in[1]), t)
	}
}

func TestBigFloatConv_Float(t *testing.T) {
	assertEquals(15.0, bigFloat(15.0).Float64(), t)
	assertEquals(-7.2, bigFloat(-7.2).Float64(), t)
	assertEquals(27.9, bigFloat(27.9).Float64(), t)
}

func TestBigFloatConv_BigRat(t *testing.T) {
	assertEquals(bigRat(15), bigFloat(15).BigRat(), t)
	assertEquals(bigRat(-7.2), bigFloat(-7.2).BigRat(), t)
	assertEquals(bigRat(27.9), bigFloat(27.9).BigRat(), t)
	x := bigFloat(20)
	rat := x.BigRat()
	rat.Neg(rat)
	assert(x.Float64() == 20, "Changing output of BigFloat.BigRat() shouldn't alter BigFloat", t)
}

func TestBigFloatComparison(t *testing.T) {
	msg := "Invalid float comparison "
	compare := func(x, y float64) int { return bigFloat(x).CompareTo(bigFloat(y)) }
	assert(compare(10, 5) > 0, msg+"10>5", t)
	assert(compare(7, 7) == 0, msg+"7==7", t)
	assert(compare(-10, -5) < 0, msg+"-10<-5", t)
}

func TestBigFloatPromotion_ShouldPromote(t *testing.T) {
	assert(!bigFloat(posInf).ShouldPromote(), "Float(+infinity).ShouldPromote() should be false", t)
	assert(!bigFloat(negInf).ShouldPromote(), "Float(-infinity).ShouldPromote() should be false", t)
	assert(!bigFloat(nan).ShouldPromote(), "Float(nan).ShouldPromote() should be false", t)
	assert(!bigFloat(7.0).ShouldPromote(), "Float(7.0).ShouldPromote() should be false", t)
}

func TestBigFloatPromotion_Promote(t *testing.T) {
	assertEquals(bigFloat(7), bigFloat(7).Promote(), t)
	assertEquals(bigFloat(-2.6), bigFloat(-2.6).Promote(), t)
}
