package models

import "github.com/go-playground/validator/v10"

type CreateProductParams struct {
	UserId      int32  `json:"userId" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int32  `json:"price" validate:"required"`
}

func (c *CreateProductParams) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		return err
	}
	return nil
}
