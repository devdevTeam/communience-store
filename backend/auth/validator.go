package auth

import (
	"errors"
	"regexp"
)

var mailreg = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func signupValidator(mail, password, userName string) error {
	if !mailValidator(mail) {
		return errors.New("Error: mail format is invalid")
	}
	if !passwordValidator(password) {
		return errors.New("Error: password format is invalid")
	}
	if !userNameValidator(userName) {
		return errors.New("Error: userName format is invalid")
	}
	return nil
}

func mailValidator(mail string) bool {
	if !mailreg.MatchString(mail) {
		return false
	}
	return true
}

func passwordValidator(password string) bool {
	if len(password) < 8 {
		return false
	}
	return true
}

func userNameValidator(userName string) bool {
	if len(userName) == 0 {
		return false
	}
	return true
}
