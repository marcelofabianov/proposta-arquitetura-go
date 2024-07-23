package domain

import "errors"

var (
	ErrUserNotFound = errors.New("error_user_not_found")
	ErrEmailExists  = errors.New("error_email_exists")
)
