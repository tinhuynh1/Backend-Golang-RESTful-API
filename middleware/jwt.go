package middleware

import (
	"DACN-GithubTrending/model"
	"DACN-GithubTrending/security"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {
	config := middleware.JWTConfig{
		Claims:     &model.JwtCustomClaims{},
		SigningKey: []byte(security.SECRET_KEY),
	}
	return middleware.JWTWithConfig(config)
}
