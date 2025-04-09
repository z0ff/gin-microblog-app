package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword はパスワードをハッシュ化する
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePassword はハッシュ化されたパスワードと比較する
func ComparePassword(hash string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
