package validator

import v10 "github.com/go-playground/validator/v10"

var validator = v10.New()

func Validate() *v10.Validate {
	return validator
}
