/*
Package numeric contains structs and interfaces that represent numeric types.
It creates wrappers around numeric types, giving them a common interface
that can be used to handle numbers (see Numeric interface) and a type that
provides user-friendly interaction (see Number type). This package is not
intended to replace the builtin types (int, float64, ...), but rather to
make number manipulation easier in cases where it's necessary to support a
wide range of approximated numbers without the performance penalty of
always using an arbitrary precision type (big.Int and big.Rat).
It's important to note that since one of the Numeric types is Float,
calculations that require exact numbers or decimal precision should not
use this package.
All types in this package are immutable.


Numeric interface

The default implementations of the Numeric interface (see Float and BigFloat)
are aliases to the builtin types (float64 and big.Rat) and, therefore, one may
obtain a Numeric value by simply casting to the Numeric type:
  x := 15.0 // float64
  y := numeric.Float(x) // numeric.Float
  fmt.Println(y) // => 15

However, Numeric types are merely wrappers around their builtin counterparts
and don't provide support for automatic conversion between Numeric types, e.g.,
when a Float value reaches infinity and needs to be converted to a BigFloat.


Number type

In order to get automatic promotion from a smaller type (e.g. Float) to a
larger type (e.g. BigFloat), the Number type should be used.


Usage

A value of Number type can be created through the NewNumber() and NewNumberSafe()
methods. Both methods accept an arbitrary type (interface{}), and attempt
to convert it to a Number value. If such a conversion does not exist,
NewNumberSafe() will return an error and NewNumber() will panic.
  x, err := numeric.NewNumberSafe("string")
  if err != nil {
    handleError(err)
  }

*/
package numeric
