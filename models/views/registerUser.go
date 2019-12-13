package views

import (
	"errors"
	"regexp"
)

type RegisterUser struct {
	Email    string
	Password string
}

func (u *RegisterUser) Validate() error {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if u.Email == "" || !emailRegex.MatchString(u.Email) {
		return errors.New("Invalid input for email")
	}
	if u.Password == "" {
		return errors.New("Invalid input for password")
	} else {
		if len(u.Password) < 6 {
			return errors.New("Password minimum is 6 characters")
		}
	}
	return nil
}
