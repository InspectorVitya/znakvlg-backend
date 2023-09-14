package model

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password")
	ErrUserBlocked     = errors.New("user blocked")
	ErrInternalService = errors.New("err services")
)
