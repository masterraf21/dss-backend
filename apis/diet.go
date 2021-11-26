package apis

import (
	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/middleware"
	"github.com/masterraf21/dss-backend/models"
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
	bearerVerify := middleware.CreateBearerVerify()

	group.POST("", r.FindDiet, bearerVerify)
}

func (r *dietRouter) FindDiet(c echo.Context) (err error) {
	return
}
