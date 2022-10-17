package usescases

import (
	"context"
	"time"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/google/uuid"
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
	userCreated, err := u.userRepo.CreateUser(ctx, &domain.User{
		Id:         uuid.New().String(),
		Name:       user.Name,
		Email:      user.Email,
		Phone:      user.Phone,
		Password:   user.Password,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	})

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
