package security
import (
	"DACN-GithubTrending/model"

	"github.com/dgrijalva/jwt-go"
)

func GenToken(user model.User) (string, error){
	claims := &model.JwtCustomClaims{
		UserId: user.UserId,
		Role: user.Role,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWhitClaims(jwt.SigninMethodES256, claims)
	result, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil{
		return "", err
	}
	
}