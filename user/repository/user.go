package repository

import (
	"context"

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
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
