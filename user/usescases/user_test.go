package usescases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/HoangTheQuyen96/user-service/domain/mocks"
	"github.com/HoangTheQuyen96/user-service/user/usescases"
	"github.com/golang/mock/gomock"
)

func TestRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserRepo := mocks.NewMockUserRepository(ctrl)

	mockCreateUserRequest := domain.CreateUserRequest{
		Name:     "Hoang",
		Email:    "HoangTheQuyen01@gmail.com",
		Phone:    "0123456789",
		Password: "123456",
	}
	t.Run("Fail", func(t *testing.T) {
		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))

		uc := usescases.NewUserUsecase(mockUserRepo)

		_, err := uc.Register(context.TODO(), &mockCreateUserRequest)

		if err != nil {
			t.Logf("Should be return error when create to dbs user error")
		}
	})
	t.Run("Success", func(t *testing.T) {
		mockCreateUserRequest := domain.CreateUserRequest{
			Name:     "Hoang",
			Email:    "HoangTheQuyen01@gmail.com",
			Phone:    "0123456789",
			Password: "123456",
		}

		user := domain.User{
			Id:       "1",
			Name:     "Hoang",
			Email:    "HoangTheQuyen01@gmail.com",
			Phone:    "0123456789",
			Password: "$2a$04$xeqUeIq9EhnWi3OSkiJtA.v4Bd6UEcRWXkqcaZb4cpdDZKfXlx3A2",
		}

		mockUserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&user, nil)

		uc := usescases.NewUserUsecase(mockUserRepo)

		userCreated, err := uc.Register(context.TODO(), &mockCreateUserRequest)

		if err == nil && userCreated != nil {
			t.Logf("Should be return user created when create to dbs user success")
		}
	})
}
