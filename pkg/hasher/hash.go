package hasher

import (
	"github.com/thvinhtruong/legoha/app/apperror"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the hash password before going in the db.
func HashPassword(input string) (string, error) {
	if len(input) == 0 {
		return "", apperror.ErrorEmptyField
	}

	post_password := []byte(input)

	hashed, err := bcrypt.GenerateFromPassword(post_password, 12)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// ComparePassword compares 2 passwords.
func ComparePassword(user_input string, from_db string) error {
	if (len(user_input) == 0) || (len(from_db) == 0) {
		return apperror.ErrorEmptyField
	}

	pass_1 := []byte(user_input)
	pass_2 := []byte(from_db)

	err := bcrypt.CompareHashAndPassword(pass_1, pass_2)

	return err
}
