package usescases

import (
	"context"
	"time"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(userRepo domain.UserRepository) domain.UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (u *userUsecase) Register(ctx context.Context, user *domain.CreateUserRequest) (*domain.User, error) {
	pwdHash, err := u.hashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	userCreated, err := u.userRepo.CreateUser(ctx, &domain.User{
		Id:         uuid.New().String(),
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Password:   pwdHash,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}

// hash password
func (h *userUsecase) hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 4)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}
