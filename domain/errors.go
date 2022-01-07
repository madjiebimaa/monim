package domain

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("your request item is not found")
	ErrConflict            = errors.New("your item already exist")
	ErrBadParamInput       = errors.New("given param is not valid")
)
