package usescases

import (
	"context"
	"time"

	"github.com/HoangTheQuyen96/user-service/domain"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *userUsecase) checkPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *userUsecase) Login(ctx context.Context, user *domain.LoginRequest) (*domain.LoginResponse, error) {
	userFind, err := u.userRepo.GetUserByPhoneOrEmail(ctx, user.UserName)

	if err != nil {
		return nil, err
	}

	err = u.checkPasswordHash(userFind.Password, user.Password)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Username or password is incorrect. Please, verify and resubmit.")
	}

	tokenDetails, err := u.generateToken(userFind)
	if err != nil {
		return nil, err
	}

	return &domain.LoginResponse{
		AccessToken:  tokenDetails.AccessToken,
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: tokenDetails.RefreshToken,
	}, nil
}

// generate token
func (u *userUsecase) generateToken(user *domain.User) (*domain.TokenDetails, error) {
	var jwtKeyAccess = []byte("secret")
	var jwtKeyRefesh = []byte("secret")

	tokenDetails := &domain.TokenDetails{
		AccessUuid:  uuid.New().String(),
		RefreshUuid: uuid.New().String(),
		AtExpires:   time.Now().Add(30 * time.Minute).Unix(),
		RtExpires:   time.Now().Add(24 * time.Hour).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.UserClaims{
		Id:         user.Id,
		Username:   user.Phone,
		AccessUuid: tokenDetails.AccessUuid,
		Phone:      user.Phone,
		Email:      user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenDetails.AtExpires,
		},
	})

	accessTokenString, err := accessToken.SignedString(jwtKeyAccess)

	if err != nil {
		return nil, err
	}

	tokenDetails.AccessToken = accessTokenString

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &domain.UserClaims{
		Id:          user.Id,
		Username:    user.Phone,
		RefreshUuid: tokenDetails.RefreshUuid,
		Phone:       user.Phone,
		Email:       user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenDetails.AtExpires,
		},
	})

	refreshTokenString, err := refreshToken.SignedString(jwtKeyRefesh)
	if err != nil {
		return nil, err
	}
	tokenDetails.RefreshToken = refreshTokenString

	return tokenDetails, nil
}
