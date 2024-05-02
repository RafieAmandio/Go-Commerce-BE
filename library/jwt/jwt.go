package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// JWTSecret is the secret key used to sign JWT tokens
var JWTSecret = []byte("asdjaskdasjkdkjqwnejk")

// GenerateJWTToken generates a JWT token for the given customer ID
func GenerateJWTToken(customerID string) (string, error) {
	// Create a new JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["customer_id"] = customerID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Generate encoded token and return it
	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseJWTToken parses and validates a JWT token
func ParseJWTToken(tokenString string) (*jwt.Token, error) {
	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JWTSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Validate token
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	return token, nil
}
