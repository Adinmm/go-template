package dto

import (
	"github.com/go-ozzo/ozzo-validation/v4"
)

type LoginInput struct {
	Name     string `json:"name" validate:"required,min=3"`
	Password string `json:"password" validate:"required,min=8"`
}


func (i LoginInput) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Name,
			validation.Required.Error("Name is required"),
			validation.Length(3, 0).Error("Name must be at least 3 characters long"),
		),
		validation.Field(&i.Password,
			validation.Required.Error("Password is required"),
			validation.Length(8, 0).Error("Password must be at least 8 characters long"),
		),
	)
}


