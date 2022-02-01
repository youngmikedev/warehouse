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

var (
	ErrUserNotFound            = errors.New("user doesn't exists")
	ErrSessionNotFound         = errors.New("user doesn't exists")
	ErrUserAlreadyExists       = errors.New("user with such email already exists")
	ErrInvalidEmail      error = &ValidationError{
		Field: "email",
		Err:   errors.New("invalid email address"),
	}
	ErrValidation = errors.New("validation error")
)
