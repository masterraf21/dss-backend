package usecases

import "github.com/masterraf21/dss-backend/models"

type dietTypeUsecase struct {
	Repo models.DietTypeRepository
}

func NewDietTypeUsecase(dtr models.DietTypeRepository) models.DietTypeUsecase {
	return &dietTypeUsecase{
		Repo: dtr,
	}
}

func (u *dietTypeUsecase) Create(body models.DietTypeBody) (id uint32, err error) {
	dietType := models.DietType{
		Name:        body.Name,
		Description: body.Description,
		Operation:   body.Operation,
		Amount:      body.Amount,
	}

	id, err = u.Repo.Store(&dietType)
	if err != nil {
		return
	}

	return
}

func (u *dietTypeUsecase) GetAll() (res []models.DietType, err error) {
	res, err = u.Repo.GetAll()
	return
}

func (u *dietTypeUsecase) GetByID(id uint32) (res *models.DietType, err error) {
	res, err = u.Repo.GetByID(id)
	return
}
