package middleware

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GenerateJWTToken(user, email, gender, role string) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user
	claims["email"] = email
	claims["gender"] = gender
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Generate encoded token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

var IsLoggedIn = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte("secret"),
})

func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["ADMIN"].(bool)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func IsRaqin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		isAdmin := claims["RAQIN"].(bool)
		if !isAdmin {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
