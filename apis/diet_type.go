package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/models"

	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

type dietTypeRouter struct {
	DietTypeUsecase models.DietTypeUsecase
}

func NewDietTypeRouter(dtu models.DietTypeUsecase) *dietTypeRouter {
	return &dietTypeRouter{
		DietTypeUsecase: dtu,
	}
}

func (r *dietTypeRouter) Mount(group *echo.Group) {
	group.POST("", r.CreateDietType)
	group.GET("", r.FindDietTypes)
	group.GET("/:id", r.FindDietType)
}

func (r *dietTypeRouter) FindDietTypes(c echo.Context) (err error) {
	data, err := r.DietTypeUsecase.GetAll()
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(
			c, "Error getting menu types", err)
	}

	return httpUtil.NewResponse(http.StatusOK, data).WriteResponse(c)
}

func (r *dietTypeRouter) FindDietType(c echo.Context) (err error) {
	dietTypeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Invalid Diet Type ID", err)
	}

	data, err := r.DietTypeUsecase.GetByID(uint32(dietTypeID))
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c,
			"Errror getting Diet Type by ID", err)
	}

	return httpUtil.NewResponse(http.StatusOK, data).WriteResponse(c)
}

func (r *dietTypeRouter) CreateDietType(c echo.Context) (err error) {
	var body models.DietTypeBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}

	id, err := r.DietTypeUsecase.Create(body)
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error creating diet type", err)
	}

	return httpUtil.NewResponse(http.StatusCreated, id).WriteResponse(c)
}
