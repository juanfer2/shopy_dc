package validations

import (
	"fmt"
	"regexp"
)

// Regular expression to validate email address.
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// EmailValidator checks if string is a valid email address.
type EmailValidator struct {
}

func (v EmailValidator) Validate(val any) (bool, error) {
	if !mailRe.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}
	return true, nil
}
