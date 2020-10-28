package users

import (
	"errors"
	"log"
	"sawali/pkg/toyerrs"

	"golang.org/x/crypto/bcrypt"
)

// PwArgs to satisfy ErrArgs
type PwArgs []string

// Len returns number of args
func (a PwArgs) Len() int {
	return len(a)
}

// HashAndSalt uses GenerateFromPassword to hash & salt pwd
func HashAndSalt(pwd []byte) (string, *toyerrs.Error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", toyerrs.New(
			"users.ComparePasswords",
			err,
			"PWDERR",
			// asterisks instead of actual user password
			// because we don't want to have it logged
			// or stored
			PwArgs{"*****"},
		)
	}

	return string(hash), nil
}

// ComparePasswords compares the password input by the user
// to the one that is stored in the DB.
func ComparePasswords(hash string, plain string) *toyerrs.Error {
	bh := []byte(hash)
	bp := []byte(plain)

	err := bcrypt.CompareHashAndPassword(bh, bp)
	if err != nil {
		log.Println(err)
		return toyerrs.New(
			"users.ComparePasswords",
			errors.New("invalid credentials"),
			"PWDERR",
			PwArgs{hash, plain},
		)
	}

	return nil
}
