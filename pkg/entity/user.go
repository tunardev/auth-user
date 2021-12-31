package entity

import (
	"errors"

	"gopkg.in/asaskevich/govalidator.v9"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDTO struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (user User) IsValidRegister() error {

	if user.Username == "" {
		return errors.New("username is required")
	}

	if len(user.Username) < 1 {
		return errors.New("username must be at least 1 character")
	}

	if len(user.Username) > 20 {
		return errors.New("username is too long")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if !govalidator.IsEmail(user.Email) {
		return errors.New("email is invalid")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (user User) IsValidLogin() error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if !govalidator.IsEmail(user.Email) {
		return errors.New("email is invalid")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	if len(user.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}

func (user User) ToDTO() UserDTO {
	return UserDTO{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
