package domain

import (
	"errors"
	"fmt"
)

var (
	ErrUserPasswordHashFailed   = errors.New("error_user_password_hash_failed")
	ErrUserPersistNewUserFailed = errors.New("error_user_persist_new_user_failed")
)

func IsErrUserPasswordHashFailed(err error) bool {
	return errors.Is(err, ErrUserPasswordHashFailed)
}

func IsErrUserPersistNewUserFailed(err error) bool {
	return errors.Is(err, ErrUserPersistNewUserFailed)
}

func GetErrUserPasswordHashFailed(err error) error {
	fmt.Println("Error: ", err)

	return ErrUserPasswordHashFailed
}

func GetErrUserPersistNewUserFailed(err error) error {
	fmt.Println("Error: ", err)

	return ErrUserPersistNewUserFailed
}
