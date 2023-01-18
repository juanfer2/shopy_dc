package validations

import (
	"fmt"
	"regexp"
)

/// var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

type RegexValidation struct {
	Expression string
}

func (v RegexValidation) Validate(val any) (bool, error) {
	fmt.Println(">>>>>>>>>>>>>>>>><<v")
	fmt.Println(v)
	r := regexp.MustCompile(v.Expression)
	if !r.MatchString(val.(string)) {
		return false, fmt.Errorf("is not regex valid")
	}

	return true, nil
}
