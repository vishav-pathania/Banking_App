package Error

import (
	"errors"
)

type ValidationErr struct {
	err             error
	statusCode      int
	specificMessage string
}

func NewValidationErr(specificMessage string) *ValidationErr {
	verror := errors.New("validation error")
	return &ValidationErr{
		err:             verror,
		statusCode:      400,
		specificMessage: specificMessage,
	}
}
