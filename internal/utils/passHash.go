package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePass(hashedPass, pass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass)); err != nil {
		return err
	}
	return nil
}
