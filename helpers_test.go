package numeric_test

import (
	. "github.com/fernandokm/numeric"
	"math/big"
	"reflect"
	"testing"
)

func equals(x, y interface{}) bool {
	return x == y
}
func bigRatEquals(x, y interface{}) bool {
	v1, ok1 := x.(*big.Rat)
	v2, ok2 := y.(*big.Rat)
	return ok1 && ok2 && (new(big.Rat).Sub(v1, v2).Sign() == 0)
}
func bigFloatEquals(x, y interface{}) bool {
	v1, ok1 := x.(*BigFloat)
	v2, ok2 := y.(*BigFloat)
	return ok1 && ok2 && (v1.CompareTo(v2) == 0)
}
func numberEquals(x, y interface{}) bool {
	v1, ok1 := x.(Number)
	v2, ok2 := y.(Number)
	return ok1 && ok2 && (v1.Equals(v2))
}

func assertEquals(expected, actual interface{}, t *testing.T) (success bool) {
	if reflect.TypeOf(expected) == reflect.TypeOf(actual) {
		var compare func(x, y interface{}) bool
		switch expected.(type) {
		case *BigFloat:
			compare = bigFloatEquals
		case *big.Rat:
			compare = bigRatEquals
		case Number:
			compare = numberEquals
		default:
			compare = equals
		}
		if compare(expected, actual) {
			return true
		}
	}
	t.Logf("Expected %s (type %T), got %s (type %T)", expected, expected, actual, actual)
	t.Fail()
	return false
}

func assert(condition bool, msg string, t *testing.T) bool {
	if !condition {
		t.Log(msg)
		t.Fail()
	}
	return condition
}

func bigRat(x float64) *big.Rat {
	return new(big.Rat).SetFloat64(x)
}

func bigFloat(x float64) *BigFloat {
	return (*BigFloat)(bigRat(x))
}
