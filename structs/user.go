package structs

import (
	uuid "github.com/google/uuid"
)

type UserWithoutId struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserId struct {
	ID   uuid.UUID `json:"id"`
	Role int32     `json:"role"`
}

type User struct {
	UserId
	UserWithoutId
}

type UserToken struct {
	Email string `json:"email"`
	Token string `json:"token"`
	Role  int32  `json:"role"`
}

type UserResponse struct {
	Auth UserToken `json:"auth"`
}
