package repository

import (
	"context"
	"fmt"

	"github.com/HoangTheQuyen96/user-service/domain"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	fmt.Println("hihihehe")
	fmt.Println(user)
	return nil, nil
}
