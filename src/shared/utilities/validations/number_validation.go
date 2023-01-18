package validations

import (
	"fmt"
)

type NumberValidator struct {
	Min int
	Max int
}

func (v NumberValidator) Validate(val any) (bool, error) {
	num := val.(int)
	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}
	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}
	return true, nil
}
