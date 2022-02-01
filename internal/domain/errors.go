package domain

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
	Err   error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("validation error field: %s, error: %s", e.Field, e.Err)
}

func NewValidationError(field string, err error) error {
	return AppError{
		ValidationError{
			Field: field,
			Err:   err,
		},
	}
}

type AppError struct {
	Err error
}

func (e AppError) Error() string {
	return e.Err.Error()
}

// func newError(e error) error {

// }

var (
	ErrUserNotFound      error = AppError{errors.New("user doesn't exists")}
	ErrSessionNotFound   error = AppError{errors.New("user doesn't exists")}
	ErrUserAlreadyExists error = AppError{errors.New("user with such email already exists")}
	ErrInvalidEmail      error = AppError{&ValidationError{
		Field: "email",
		Err:   errors.New("invalid email address"),
	}}
	ErrValidation error = AppError{errors.New("validation error")}
	ErrInternal   error = AppError{errors.New("internal error")}
)
