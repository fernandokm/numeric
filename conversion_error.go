package numeric

import "fmt"

type ConversionError struct {
	value interface{}
}

func (e ConversionError) Error() string {
	return fmt.Sprintf("Unrecognized value %v of type %T", e.value, e.value)
}
