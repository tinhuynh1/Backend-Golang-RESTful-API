package main

import (
	"DACN-GithubTrending/db"
	"DACN-GithubTrending/handler"
	log "DACN-GithubTrending/log"
	"DACN-GithubTrending/repository/repo_impl"
	"DACN-GithubTrending/router"
	"fmt"
	"os"

	"github.com/labstack/echo"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}
func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		Password: "123456",
		Dbname:   "golang_project",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
	}
	api.SetupRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
