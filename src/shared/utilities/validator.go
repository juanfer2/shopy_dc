package utilities

import (
	"github.com/juanfer2/shopy_dc_golang/src/shared/utilities/validations"
)

func ValidatorStruct(s any) []error {
	return validations.ValidateStruct(s)
}
