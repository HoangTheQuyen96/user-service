package domain

import (
	"context"
)

type User struct {
	Id         string `json:"id" gorm:"column:id;index:id;PRIMARY_KEY;"`
	Name       string `json:"name" gorm:"column:name;"`
	Email      string `json:"email" gorm:"column:email;"`
	Phone      string `json:"phone" gorm:"column:phone;"`
	Password   string `json:"password" gorm:"column:password;"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserUsecase interface {
	Register(context.Context, *CreateUserRequest) (*User, error)
}

type UserRepository interface {
	CreateUser(context.Context, *User) (*User, error)
}
