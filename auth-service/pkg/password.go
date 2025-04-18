package pkg

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSalt(length int) (string, error) {
	salt := make([]byte, length)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func HashWithSalt(password, salt string) string {
	hash := sha256.Sum256([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}

func GenerateHashPassword(password string) (string, string, error) {
	salt, err := GenerateSalt(16)
	if err != nil {
		return "", "", err
	}

	hash := HashWithSalt(password, salt)
	return hash, salt, nil
}

func VerifyHashPassword(password, hash, salt string) bool {
	return HashWithSalt(password, salt) == hash
}
