package service

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/model"
	"github.com/sail-host/cloud/internal/global"
	"github.com/sail-host/cloud/internal/utils/hash"
)

type AuthService struct{}

type IAuthService interface {
	Login(c echo.Context, request dto.LoginRequest) (*dto.LoginResponse, *dto.BaseError)
	Register(c echo.Context, request dto.RegisterRequest) (*dto.BaseResponse, *dto.BaseError)
	CheckUserFirstTime(c echo.Context) (bool, *dto.BaseError)
}

func NewIAuthService() IAuthService {
	return &AuthService{}
}

func (s *AuthService) Login(c echo.Context, request dto.LoginRequest) (*dto.LoginResponse, *dto.BaseError) {
	var response dto.LoginResponse
	var baseError dto.BaseError

	user, err := userRepo.GetUserByEmail(request.Email)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "User not found or invalid password"
		return nil, &baseError
	}

	if !hash.CheckPasswordHash(request.Password, user.Password) {
		baseError.Status = "error"
		baseError.Message = "User not found or invalid password"
		return nil, &baseError
	}

	// Generate token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["role"] = user.Role
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(global.CONF.System.EncryptKey))
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Error generating token"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "Login successful"
	response.Token = tokenString

	return &response, nil
}

func (s *AuthService) Register(c echo.Context, request dto.RegisterRequest) (*dto.BaseResponse, *dto.BaseError) {
	var response dto.BaseResponse
	var baseError dto.BaseError

	// Validate request
	if err := c.Validate(request); err != nil {
		baseError.Status = "error"
		baseError.Message = "Validation failed"
		return nil, &baseError
	}

	// Check database users exists
	exist, err := s.CheckUserFirstTime(c)
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Database error"
		return nil, &baseError
	}

	if exist {
		baseError.Status = "error"
		baseError.Message = "User already exists"
		return nil, &baseError
	}

	hashedPassword, errBase := hash.HashPassword(request.Password)
	if errBase != nil {
		baseError.Status = "error"
		baseError.Message = "Hashing error"
		return nil, &baseError
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	errBase = userRepo.CreateUser(&user)
	if errBase != nil {
		baseError.Status = "error"
		baseError.Message = "Database error"
		return nil, &baseError
	}

	response.Status = "success"
	response.Message = "user created successfully"

	return &response, nil
}

func (s *AuthService) CheckUserFirstTime(c echo.Context) (bool, *dto.BaseError) {
	var baseError dto.BaseError

	exist, err := userRepo.CheckUsersExist()
	if err != nil {
		baseError.Status = "error"
		baseError.Message = "Database error"
		return false, &baseError
	}
	return exist, nil
}
