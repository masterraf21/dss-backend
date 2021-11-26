package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/masterraf21/dss-backend/models"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

type menuRouter struct {
	MenuUsecase models.MenuUsecase
}

func NewMenuRouter(mus models.MenuUsecase) *menuRouter {
	return &menuRouter{
		MenuUsecase: mus,
	}
}

func (r *menuRouter) Mount(group *echo.Group) {
	group.POST("", r.CreateMenu)
	group.POST("/bulk", r.BulkCreate)
	group.GET("", r.FindMenus)
	group.GET("/:id", r.FindMenu)
}

func (r *menuRouter) FindMenus(c echo.Context) (err error) {
	data, err := r.MenuUsecase.GetAll()
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Failed to get Menus", err)
	}

	return httpUtil.NewResponse(http.StatusOK, data).WriteResponse(c)
}

func (r *menuRouter) FindMenu(c echo.Context) (err error) {
	menuId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Invalid Menu ID", err)
	}

	res, err := r.MenuUsecase.GetByID(uint32(menuId))
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error getting menu by id", err)
	}

	return httpUtil.NewResponse(http.StatusOK, res).WriteResponse(c)
}

func (r *menuRouter) CreateMenu(c echo.Context) (err error) {
	var body models.MenuBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}

	id, err := r.MenuUsecase.Create(body)
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error creating menu", err)
	}

	return httpUtil.NewResponse(http.StatusCreated, id).WriteResponse(c)
}

func (r *menuRouter) BulkCreate(c echo.Context) (err error) {
	var body []models.MenuBody
	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		return httpUtil.NewError(echo.ErrBadRequest.Code).WriteError(c, "Error decoding body", err)
	}
	// fmt.Println(len(body))
	// for _, bod := range body {
	// 	fmt.Println(bod.Name)
	// }

	ids, err := r.MenuUsecase.BulkCreate(body)
	if err != nil {
		return httpUtil.NewError(echo.ErrInternalServerError.Code).WriteError(c, "Error bulk creating menu", err)
	}

	return httpUtil.NewResponse(http.StatusCreated, ids).WriteResponse(c)
}
