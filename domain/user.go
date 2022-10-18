package domain

import (
	"context"

	"github.com/dgrijalva/jwt-go"
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

type LoginRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}
type UserUsecase interface {
	Register(context.Context, *CreateUserRequest) (*User, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByPhoneOrEmail(ctx context.Context, emailOrPhone string) (*User, error)
}
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type UserClaims struct {
	Id          string `json:"id"`
	Username    string `json:"username"`
	AccessUuid  string `json:"access_uuid"`
	RefreshUuid string `json:"refresh_uuid"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	jwt.StandardClaims
}
