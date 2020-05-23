package security

import (
	"DACN-GithubTrending/model"

	"github.com/dgrijalva/jwt-go"
)

// GenToken ...
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId:         user.UserId,
		Role:           user.Role,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", err
	}
	return result, err
}
