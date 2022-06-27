package model

import (
	"time"

	"github.com/go-playground/validator"
)

type CreateProductRequest struct {
	Name  string `form:"name" validate:"required"`
	Price int    `form:"price" validate:"required"`
	Stock int    `form:"stock" validate:"required"`
}

type CreateProductResponse struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type UpdateProductRequest struct {
	Name      string `form:"name"`
	Price     int    `form:"price"`
	Stock     int    `form:"stock"`
	UpdatedAt time.Time
}

type GetProductResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

var validate = validator.New()

func ValidateStruct(product CreateProductRequest) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(product)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
