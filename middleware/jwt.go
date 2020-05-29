package middleware
import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo"
	"DACN-GithubTrending/model"
	"DACN-GithubTrending/security"
)
func JWTMiddleware() echo.MiddlewareFunc{
	config := middleware.JWTConfig{
		Claims: &model.JwtCustomClaims{},
		SigningKey: security.SECRET_KEY,	
	}
	return middleware.JWTWithConfig(config)
}