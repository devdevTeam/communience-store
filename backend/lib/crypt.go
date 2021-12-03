package lib

import "golang.org/x/crypto/bcrypt"

func Encryption(plain string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plain), 4)
	return hash, err
}

func HashChecker(plain string, hash []byte) error {
	Bplain := []byte(plain)
	err := bcrypt.CompareHashAndPassword(hash, Bplain)
	return err
}
