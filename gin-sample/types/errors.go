package types

import "errors"

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidUserOrPassword = errors.New("invalid User or Password")
)
