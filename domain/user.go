package domain

import (
	"context"
	"time"
)

type User struct {
	Id         string    `json:"id" gorm:"column:id;index:id;PRIMARY_KEY;"`
	Name       string    `json:"name" gorm:"column:name;"`
	Email      string    `json:"email" gorm:"column:email;"`
	Phone      string    `json:"string" gorm:"column:phone;"`
	Password   string    `json:"password" gorm:"column:password;"`
	CreateTime time.Time `json:"create_time" gorm:"column:create_time;"`
	UpdateTime time.Time `json:"update_time" gorm:"column:update_time;"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"string"`
	Password string `json:"password"`
}

type UserUsecase interface {
	Register(context.Context, *CreateUserRequest) (*User, error)
}

type UserRepository interface {
	CreateUser(context.Context, *User) (*User, error)
}
