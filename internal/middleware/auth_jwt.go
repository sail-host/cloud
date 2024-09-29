package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sail-host/cloud/internal/app/dto"
	"github.com/sail-host/cloud/internal/app/repository"
	jwtUtil "github.com/sail-host/cloud/internal/utils/jwt"
)

func AuthJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		// Check if the token follows "Bearer <token>" format
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		tokenString := parts[1]

		// Parse the token
		j := jwtUtil.NewJWT()
		token, err := j.ParseToken(tokenString)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		if token.BaseClaims.ID == 0 || token.BaseClaims.Token == "" {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		authRepo := repository.NewIAuthRepo()
		auth, err := authRepo.GetAuthTokenByToken(token.BaseClaims.Token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		if auth.ID == 0 {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		if auth.ExpiredAt.Before(time.Now()) {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Token expired"})
		}

		userRepo := repository.NewIUserRepo()
		user, err := userRepo.GetUserByID(strconv.FormatUint(uint64(auth.UserID), 10))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, dto.BaseError{Status: "error", Message: "Unauthorized"})
		}

		c.Set("user", user)
		return next(c)
	}
}
