package users

import (
	"errors"
	"sawali/pkg/toyerrs"

	"golang.org/x/crypto/bcrypt"
)

// PwArgs to satisfy ErrArgs
type PwArgs []string

// Len returns number of args
func (a PwArgs) Len() int {
	return len(a)
}

// HashPassword uses GenerateFromPassword to hash & salt given password
func HashPassword(pwd string) (string, *toyerrs.Error) {
	err := &toyerrs.Error{}

	bytes, e := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if e != nil {
		err = toyerrs.New(
			"users.ComparePasswords",
			e,
			"PWDERR",
			// asterisks instead of actual user password
			// because we don't want to have it logged
			// or stored
			PwArgs{"****"},
		)
	}

	return string(bytes), err
}

// ComparePasswords compares the password input by the user
// to the one that is stored in the DB.
func ComparePasswords(hash, pwd string) *toyerrs.Error {
	err := &toyerrs.Error{}
	bh := []byte(hash)
	bp := []byte(pwd)

	e := bcrypt.CompareHashAndPassword(bh, bp)
	if e != nil {
		err = toyerrs.New(
			"users.ComparePasswords",
			errors.New("invalid credentials"),
			"PWDERR",
			PwArgs{hash, pwd},
		)
	}

	return err
}
