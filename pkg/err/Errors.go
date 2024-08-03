package errors

import "errors"

var (
    ErrUserAlreadyExists = errors.New("user with this email already exists")
    ErrInvalidPassword = errors.New("invalid password")
    ErrTweetIsNotOfTheUser = errors.New("tweet is not of the user")
    ErrTweetNotFound = errors.New("tweet not found")
)