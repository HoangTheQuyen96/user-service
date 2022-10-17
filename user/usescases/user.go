package usescases

import (
	"context"

	"github.com/HoangTheQuyen96/user-service/domain"
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
		Id:       "1",
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		Password: user.Password,
	})

	if err != nil {
		return nil, err
	}
	return userCreated, nil
}
