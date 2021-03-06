package models

import "fmt"

type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
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
