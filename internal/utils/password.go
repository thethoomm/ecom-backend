package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	passwordBytes := []byte(password)

	bytes, err := bcrypt.GenerateFromPassword(passwordBytes, 10)

	hashedStringPassword := string(bytes)

	return hashedStringPassword, err
}

func VerifyPassword(password, hash string) bool {
	passwordBytes := []byte(password)
	hashedPasswordBytes := []byte(hash)

	err := bcrypt.CompareHashAndPassword(passwordBytes, hashedPasswordBytes)

	return err == nil
}
