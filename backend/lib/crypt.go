package lib

import "golang.org/x/crypto/bcrypt"

func Encryption(plain string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 4)
	return hash, err
}
