package models

import "fmt"

type User struct {
	Id        int
	FirstName string
	LastName  string
	Email     string
	Username  string
	Password  string
}

func (u *User) ValidUser() error {
	if u.FirstName == "" {
		return fmt.Errorf("fistname missing")
	}
	if u.LastName == "" {
		return fmt.Errorf("lastname missing")
	}
	if u.Email == "" {
		return fmt.Errorf("email missing")
	}
	if u.Username == "" {
		return fmt.Errorf("username missing")
	}
	if u.Password == "" {
		return fmt.Errorf("password missing")
	}
	if err := u.validPassword(); err != nil {
		return err
	}
	// otherwise
	return nil
}

func (u *User) validPassword() error {
	if len(u.Password) < 5 {
		return fmt.Errorf("invalid password")
	}
	// otherwise
	return nil
}
