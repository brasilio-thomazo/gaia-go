package security

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
)

var (
	memory  = uint32(64 * 1024)
	threads = uint8(2)
	len     = uint32(32)
	times   = uint32(3)
	cost    = int(4)
)

func getSalt() ([]byte, error) {
	b64 := os.Getenv("ARGON2_SALT")
	if b64 == "" {
		return nil, fmt.Errorf("ARGON2_SALT is not set")
	}
	return base64.RawStdEncoding.DecodeString(b64)
}

func Argon2HashPassword(password string) (string, error) {
	salt, err := getSalt()
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, times, memory, threads, len)
	return string(hash), nil
}

func BcryptHashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Argon2CheckPassword(password string, hash string) bool {
	salt, err := getSalt()
	if err != nil {
		log.Println("error getting salt:", err)
		return false
	}
	hashToCompare := argon2.IDKey([]byte(password), salt, times, memory, threads, len)
	return string(hashToCompare) == hash
}

func BcryptCheckPassword(password string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
