package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func ComparePassword(hashedPassword string, passwordStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordStr))
}
