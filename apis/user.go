package apis

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/middleware"
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
	// group.GET("", r.FindUsers)
	bearerVerify := middleware.CreateBearerVerify()
	group.GET("/:id", r.FindUser, bearerVerify)
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

	res, err := r.userUsecase.Register(body)
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error Creating User", err)
	}

	httpUtil.NewResponse(http.StatusCreated, res).WriteResponse(c)

	return
}

func (r *userRouter) Login(c echo.Context) (err error) {
	var body models.LoginBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}

	var errCode *echo.HTTPError
	res, err := r.userUsecase.Login(body)
	if err != nil {
		if err.Error() == "intenal_error" {
			errCode = echo.ErrInternalServerError
		} else {
			errCode = echo.ErrUnauthorized
		}

		return httpUtil.NewError(errCode.Code).WriteError(c, "Error Loggin in", err)
	}

	return httpUtil.NewResponse(http.StatusOK, res).WriteResponse(c)
}
