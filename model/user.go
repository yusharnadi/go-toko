package model

import "time"

type CreateUserRequest struct {
	Name     string `form:"name" validate:"required"`
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

type CreateUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UpdateUserRequest struct {
	Name      string `form:"name" validate:"required"`
	Email     string `form:"email" validate:"required"`
	Password  string `form:"password" validate:"required"`
	UpdatedAt time.Time
}

type GetUserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
