package models

import (
	"errors"

	"github.com/badoux/checkmail"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"nome,omitempty"`
	Nick     string `json:"nick,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"senha,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("The name is required and cannot be blank")
	}

	if user.Nick == "" {
		return errors.New("The Nick is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("The Email is required and cannot be blank")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("The email entered is invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("The Password is required and cannot be blank")
	}

	return nil
}
