package models

import "github.com/go-playground/validator/v10"

type RegisterPayload struct {
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"firstName", validate:"required"`
	LastName  string `json:"firstName", validate:"required"`
}

func (c *RegisterPayload) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
