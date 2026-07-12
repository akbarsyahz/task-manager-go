package security

import (
	"crypto/rand"
	"crypto/subtle"

	dtoM "taskManager/db/model"

	"golang.org/x/crypto/argon2"
)

var saltLength = 16

// GenerateSalt for Generating the salt
func GenerateSalt() ([]byte, error) {
	salt := make([]byte, saltLength)

	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}

	return salt, nil
}

// HashArgon2Register for Registarion hasing
func HashArgon2Register(password string, salt []byte) ([]byte, error) {
	hashPass := argon2.IDKey([]byte(password), salt, 20, 64*1024, 4, 32)
	return hashPass, nil
}

// HashArgon2Pass for Login
func HashArgon2Pass(password string, salt []byte) ([]byte, error) {
	hashPass := argon2.IDKey([]byte(password), salt, 20, 64*1024, 4, 32)
	return hashPass, nil
}

// ComparingPassword for comparing password
func ComparingPassword(password string, user dtoM.LoginUser) (bool, error) {
	result := false

	newHash, err := HashArgon2Pass(password, user.Salted)
	if err != nil {
		return false, err
	}

	match := subtle.ConstantTimeCompare(user.Hash, newHash)
	if match == 1 {
		result = true
	}

	return result, nil
}
