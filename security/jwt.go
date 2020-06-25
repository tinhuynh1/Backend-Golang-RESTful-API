package security

import (
	"DACN-GithubTrending/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const SECRET_KEY = "mr.hnqtin@gmail.com"

// GenToken ...
func GenToken(user model.User) (string, error) {
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	result, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return result, nil
}
