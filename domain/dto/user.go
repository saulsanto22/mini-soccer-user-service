package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	PhoneNum string    `json:"phone_number"`
	Username string    `json:"username"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterResponse struct {
	User        UserResponse
	Name        string `json:"name"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	ConfrimPass string `json:"confrim_pass"`
	PhoneNumber string `json:"phone_number"`
	RoleId      uint
}

type RegisterRequest struct {
	Name        string `json:"name" validate:"required"`
	Username    string `json:"username" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	ConfrimPass string `json:"confrim_pass" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	RoleId      uint
}

type UpdateRequest struct {
	Name        string  `json:"name" validate:"required"`
	Username    string  `json:"username" validate:"required"`
	Email       string  `json:"email" validate:"required,email"`
	Password    *string `json:"password,omitempty"`
	ConfrimPass *string `json:"confrim_pass,omitempty"`
	PhoneNumber string  `json:"phone_number" validate:"required"`
	RoleId      uint
}
