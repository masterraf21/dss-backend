package usecases

import (
	"fmt"
	"strings"

	"github.com/masterraf21/dss-backend/models"
)

type menuUsecase struct {
	Repo models.MenuRepository
}

// Newmodels.MenuUsecase willl create menu usecase
func NewMenuUsecase(mnr models.MenuRepository) models.MenuUsecase {
	return &menuUsecase{
		Repo: mnr,
	}
}

func extractLabel(ingredients string) (res []string) {
	res = strings.Split(ingredients, "--")
	return
}

// func extractLabel()

func (u *menuUsecase) Create(body models.MenuBody) (res uint32, err error) {
	menu := models.Menu{
		Name:        body.Name,
		Calorie:     body.Calorie,
		Recipe:      body.Recipe,
		PictureURL:  body.PictureURL,
		Ingredients: extractLabel(body.Ingredients),
		Labels:      body.Labels,
	}

	res, err = u.Repo.Store(&menu)
	if err != nil {
		return
	}

	return
}

func (u *menuUsecase) GetAll() (res []models.Menu, err error) {
	res, err = u.Repo.GetAll()
	if len(res) > 20 {
		res = res[:20]
	}
	return
}

func (u *menuUsecase) GetByID(id uint32) (res *models.Menu, err error) {
	res, err = u.Repo.GetByID(id)
	return
}

func (u *menuUsecase) BulkCreate(bodys []models.MenuBody) (res []uint32, err error) {
	// var menu models.Menu
	var menuPtr []*models.Menu
	menuPtr = make([]*models.Menu, 0)
	// fmt.Println(bodys)
	for _, body := range bodys {
		// var menu models.Menu
		menu := models.Menu{
			Name:        body.Name,
			Calorie:     body.Calorie,
			Recipe:      body.Recipe,
			PictureURL:  body.PictureURL,
			Ingredients: extractLabel(body.Ingredients),
			Labels:      body.Labels,
		}
		menuPtr = append(menuPtr, &menu)
	}

	// fmt.Println(menuPtr)

	for _, p := range menuPtr {
		fmt.Println(p)
	}

	res, err = u.Repo.BulkStore(menuPtr)
	if err != nil {
		return
	}

	return
}
