package main

import (
	"DACN-GithubTrending/db"
	"DACN-GithubTrending/handler"
	"DACN-GithubTrending/helper"
	log "DACN-GithubTrending/log"
	"DACN-GithubTrending/repository/repo_impl"
	"DACN-GithubTrending/router"
	"fmt"
	"os"
	"time"

	"github.com/labstack/echo"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}
func main() {
	sql := &db.Sql{
		Host:     "database",
		Port:     5432,
		Username: "postgres",
		Password: "123456",
		Dbname:   "golang_project",
	}
	sql.Connect()
	defer sql.Close()

	e := echo.New()

	structValidator := helper.NewStructValidator()
	structValidator.RegisterValidate()

	e.Validator = structValidator
	userHandler := handler.UserHandler{
		UserRepo: repo_impl.NewUserRepo(sql),
	}
	repoHandler := handler.RepoHandler{
		GithubRepo: repo_impl.NewGithubRepo(sql),
	}
	api := router.API{
		Echo:        e,
		UserHandler: userHandler,
		RepoHandler: repoHandler,
	}
	api.SetupRouter()
	go scheduleUpdateTrending(15*time.Second, repoHandler)
	e.Logger.Fatal(e.Start(":3000"))
}
func scheduleUpdateTrending(timeSchedule time.Duration, handler handler.RepoHandler) {
	ticker := time.NewTicker(timeSchedule)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Checking from github...")
				helper.CrawlRepo(handler.GithubRepo)
			}
		}
	}()
}
