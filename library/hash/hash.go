package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt
func HashPassword(password string) (string, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CompareHashAndPassword compares the given hashed password with the provided password
func CompareHashAndPassword(hashedPassword, password string) error {
	// Compare hashed password with provided password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err // Passwords do not match
	}

	return nil // Passwords match
}
