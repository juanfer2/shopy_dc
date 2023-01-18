package validations

type DefaultValidator struct{}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}
