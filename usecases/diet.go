package usecases

import (
	"github.com/masterraf21/dss-backend/models"
)

type dietUsecase struct {
	dietTypeRepository models.DietTypeRepository
}

// NewDietUsecase will create usecase
func NewDietUsecase(
	dtsr models.DietTypeRepository,
) models.DietUsecase {
	return &dietUsecase{
		dietTypeRepository: dtsr,
	}
}

func (u *dietUsecase) CountREE(
	gender models.GENDER, weight float32, height float32, age int) (result float32) {
	if gender == models.FEMALE {
		result = 655.1 + 13.8*weight + 5.0*height - 6.8*float32(age)
	} else if gender == models.MALE {
		result = 66.5 + 9.6*weight + 1.8*height - 4.7*float32(age)
	}

	return
}

func (u *dietUsecase) CountCE(ree float32, af models.ACTIVITY_FACTOR) (result float32) {
	switch af {
	case models.SEDENTARY:
		result = ree * 1.2
	case models.LIGHTLY_ACTIVE:
		result = ree * 1.375
	case models.MODERATELY_ACTIVE:
		result = ree * 1.55
	case models.VERY_ACTIVE:
		result = ree * 1.725
	case models.EXTRA_ACTIVE:
		result = ree * 1.9
	}

	return
}

func (u *dietUsecase) CountDCR(ce float32, dietType models.DietType) (result float32) {
	switch dietType.Operation {
	case models.MINUS:
		result = ce - dietType.Amount
	case models.PLUS:
		result = ce - dietType.Amount
	}
	return
}

func (u *dietUsecase) FindMenu(dcr float32, duration int) {
}
