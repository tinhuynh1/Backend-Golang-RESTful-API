package middleware
import (
	"DACN-GithubTrending/model"
	"DACN-GithubTrending/model/req"
	"net/http"

	"github.com/labstack/echo"
)
func ISAdmin() echo.MiddlewareFunc  {
	return func(next echo.HandlerFunc) echo.HandlerFunc  {
		return func(c echo.Context) error {
			//handle logic
			req := req.ReqSignIn{}
			if err := c.Bind(&req); err != nil{
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message: err.Error(),
					Data: nil,
				})
			}
			if req.Email != "admin@gmail.com"{
				return c.JSON(http.StatusBadRequest, model.Response{
					StatusCode: http.StatusBadRequest,
					Message: "Bạn không có quyền gọi api này",
					Data: nil,
				})
			}
			return next(c)
		}
	}
}