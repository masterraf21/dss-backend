package apis

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/models"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

type dietRouter struct {
	DietUsecase models.DietUsecase
}

func NewDietRouter(dtr models.DietUsecase) *dietRouter {
	return &dietRouter{
		DietUsecase: dtr,
	}
}

func (r *dietRouter) Mount(group *echo.Group) {
	// bearerVerify := middleware.CreateBearerVerify()

	// group.GET("/test", r.Test, bearerVerify)
	group.POST("", r.FindDiet)
}

func (r *dietRouter) FindDiet(c echo.Context) (err error) {
	var body models.DietPlanBody

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}

	err = r.DietUsecase.FindDietPlan(body)
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error finding plan", err)
	}

	return httpUtil.NewResponse(http.StatusOK, nil).WriteResponse(c)
}

func (r *dietRouter) Test(c echo.Context) (err error) {
	return
}
