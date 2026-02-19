package go_auth_api_tests

import (
	"errors"
	"strings"
)

func Register(email, password string) error {
	if email == "" || password == "" {
		return errors.New("email or password is empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @")
	}
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	return nil
}
