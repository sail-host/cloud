package service

import (
	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
)

type AuthService struct{}

type IAuthService interface {
	Login(c echo.Context, request dto.LoginRequest) (*dto.LoginResponse, error)
}

func NewIAuthService() IAuthService {
	return &AuthService{}
}

func (s *AuthService) Login(c echo.Context, request dto.LoginRequest) (*dto.LoginResponse, error) {
	var response dto.LoginResponse

	response.Status = "success"
	response.Message = "Login successful"
	response.Token = "1234567890"

	return &response, nil
}
