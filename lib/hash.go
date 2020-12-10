package lib

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSalt(text string) (string, error) {
	salt := make([]byte, 24)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}

	hash, err := scrypt.Key([]byte(text), salt, 1<<14, 8, 1, 16)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash), nil
}
