package customvalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/okeefem2/simple_bank/internal/currency"
)

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if c, ok := fl.Field().Interface().(string); ok {
		return currency.IsSupportedCurrency(c)
	}
	return false
}
