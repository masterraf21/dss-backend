package usecases

import (
	"github.com/emirpasic/gods/sets/hashset"
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

func extractLabel(ingredients []models.Ingredient) (res []string) {
	set := hashset.New()
	for _, ingredient := range ingredients {
		labels := ingredient.Labels
		for _, label := range labels {
			set.Add(label)
		}
	}

	res = make([]string, 0)
	for _, val := range set.Values() {
		res = append(res, val.(string))
	}

	return
}

func (u *menuUsecase) Create(body models.MenuBody) (res uint32, err error) {
	menu := models.Menu{
		Name:         body.Name,
		CalorieCount: body.CalorieCount,
		Recipe:       body.Recipe,
		PictureURL:   body.PictureURL,
		Ingredients:  body.Ingredients,
		Labels:       body.Labels,
	}

	res, err = u.Repo.Store(&menu)
	if err != nil {
		return
	}

	return
}

func (u *menuUsecase) GetAll() (res []models.Menu, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *menuUsecase) GetByID(id uint32) (res *models.Menu, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
