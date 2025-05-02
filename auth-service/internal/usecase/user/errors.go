package userusecase

import (
	"errors"
)

var (
	ErrUserNotFound      = errors.New("user_not_found")
	ErrUserNotActive     = errors.New("user_not_active")
	ErrPasswordIncorrect = errors.New("password_is_incorrect")
)
