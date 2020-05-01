package repo_impl

import (
	"DACN-GithubTrending/banana"
	"DACN-GithubTrending/db"
	"DACN-GithubTrending/log"
	"DACN-GithubTrending/model"
	"DACN-GithubTrending/model/req"
	"DACN-GithubTrending/repository"
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type UserRepoImpl struct {
	sql *db.Sql
}

func NewUserRepo(sql *db.Sql) repository.UserRepo {
	return &UserRepoImpl{
		sql: sql,
	}
}
func (u UserRepoImpl) SaveUser(context context.Context, user model.User) (model.User, error) {
	statement := `
	INSERT INTO users(user_id, email, password, role, full_name, created_at, updated_at)
	VALUES(:user_id, :email, :password, :role, :full_name, :created_at, :updated_at)
	`
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := u.sql.Db.NamedExecContext(context, statement, user)
	if err != nil {
		log.Error(err.Error())
		if err, ok := err.(*pq.Error); ok {
			if err.Code.Name() == "unique_violation" {
				return user, banana.UserConflic
			}
		}
		return user, banana.SignUpFall
	}

	return user, nil
}
func (u *UserRepoImpl) CheckLogin(context context.Context, loginReq req.ReqSignIn) (model.User, error) {
	var user = model.User{}
	err := u.sql.Db.GetContext(context, &user, "SELECT * FROM users WHERE email=$1", loginReq.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, banana.UserNotFound
		}
		log.Error(err.Error())
		return user, err
	}
	return user, nil
}
