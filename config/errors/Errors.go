package errors

import "errors"

var (
    ErrUserAlreadyExists = errors.New("user with this email already exists")
    ErrInvalidPassword = errors.New("invalid password")
)