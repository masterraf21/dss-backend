package apis

import (
	"github.com/labstack/echo"
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
	group.POST("", r.FindDiet)
}

func (r *dietRouter) FindDiet(c echo.Context) (err error) {
	return
}
