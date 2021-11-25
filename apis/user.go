package apis

import (
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/models"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

type userRouter struct {
	userUsecase models.UserUsecase
}

func NewUserRouter(usr models.UserUsecase) *userRouter {
	return &userRouter{
		userUsecase: usr,
	}
}

func (r *userRouter) Mount(group *echo.Group) {
	group.GET("", r.FindUsers)
	group.GET("/:id", r.FindUser)
	group.POST("/register", r.Register)
	group.POST("/login", r.Login)
}

func (r *userRouter) FindUsers(c echo.Context) (err error) {
	// data, err := r.userUsecase.
	return
}

func (r *userRouter) FindUser(c echo.Context) (err error) {
	return
}

func (r *userRouter) Register(c echo.Context) (err error) {
	var body models.UserBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}

	return
}

func (r *userRouter) Login(c echo.Context) (err error) {
	return
}
