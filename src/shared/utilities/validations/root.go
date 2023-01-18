package validations

import (
	"fmt"
	"reflect"
	"strings"
)

// Name of the struct tag used in examples.
const tagName = "validate"

type Validator interface {
	Validate(any) (bool, error)
}

// Returns validator struct corresponding to validation type
func getValidatorFromTag(tag string) Validator {
	fmt.Println("   tag")
	fmt.Println("    ", tag)
	args := strings.Split(tag, ",")
	switch args[0] {
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		return EmailValidator{}
	case "regex":
		validator := RegexValidation{}
		fmt.Sscanf(strings.Join(args[1:], ","), "expression=%s", &validator.Expression)
		return validator
	}
	return DefaultValidator{}
}

// Performs actual data validation using validator definitions on the struct
func ValidateStruct(s interface{}) []error {
	errs := []error{}
	// ValueOf returns a Value representing the run-time data
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		// Get the field tag value
		tag := v.Type().Field(i).Tag.Get(tagName)

		fmt.Println("=========tag===============")
		fmt.Println(v)
		fmt.Println(tag)
		fmt.Println(tagName)
		fmt.Println(v.Type().Field(i).Tag.Get(tagName))
		// Skip if tag is not defined or ignored
		if tag == "" || tag == "-" {
			continue
		}
		// Get a validator that corresponds to a tag
		validator := getValidatorFromTag(tag)
		// Perform validation
		valid, err := validator.Validate(v.Field(i).Interface())
		// Append error to results
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}
	return errs
}
