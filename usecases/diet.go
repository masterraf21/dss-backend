package usecases

import (
	"math"
	"sort"

	"github.com/masterraf21/dss-backend/models"
)

type dietUsecase struct {
	dietTypeRepository models.DietTypeRepository
	menuRepository     models.MenuRepository
}

// NewDietUsecase will create usecase
func NewDietUsecase(
	dtsr models.DietTypeRepository, mr models.MenuRepository,
) models.DietUsecase {
	return &dietUsecase{
		dietTypeRepository: dtsr,
		menuRepository:     mr,
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

func (u *dietUsecase) CountDCR(ce float32, dietType *models.DietType) (result float32) {
	switch dietType.Operation {
	case models.MINUS:
		result = ce - dietType.Amount
	case models.PLUS:
		result = ce - dietType.Amount
	}
	return
}

func (u *dietUsecase) FindDCR(body models.DietPlanBody) (dcr float32, err error) {
	dietType, err := u.dietTypeRepository.GetByID(body.DietTypeID)
	if err != nil {
		return
	}
	ree := u.CountREE(body.Gender, body.Weight, body.Height, body.Age)
	ce := u.CountCE(ree, body.ActivityFactor)
	dcr = u.CountDCR(ce, dietType)

	return
}

type idCalorie struct {
	ID      uint32
	Calorie int
}

func extractIDCalorie(menus []models.Menu) (res []idCalorie) {
	res = make([]idCalorie, 0)
	for _, menu := range menus {
		res = append(res, idCalorie{
			ID:      menu.ID,
			Calorie: menu.CalorieCount,
		})
	}

	return
}

func subsetSumRange(input []idCalorie, n int, a int, b int) (res [][3]uint32) {
	res = make([][3]uint32, 0)

	sort.Slice(input, func(i, j int) bool {
		return input[i].Calorie < input[j].Calorie
	})

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			first := input[i].Calorie + input[j].Calorie
			if (first >= a) && (first <= b) {
				max := j + 1
				min := n - 2
				for min <= max {
					mid := min + (max-1)/2
					second := first + input[mid].Calorie
					if (second >= a) && (second <= b) {
						res = append(res, [3]uint32{input[i].ID, input[j].ID, input[mid].ID})
					}
					if second > b {
						min = mid + 1
					} else if second < a {
						max = mid - 1
					}
				}

			}
		}
	}

	return
}

func (u *dietUsecase) FindDietPlan(body models.DietPlanBody) (res *models.DietPlan, err error) {
	duration := body.Duration
	dcr, err := u.FindDCR(body)
	if err != nil {
		return
	}
	dcrUpper := int(math.Round(float64(dcr))) + 100
	dcrBottom := int(math.Round(float64(dcr))) - 100

	menus, err := u.menuRepository.GetAll()
	if err != nil {
		return
	}

	data := extractIDCalorie(menus)

	planMenuIDs := subsetSumRange(data, len(data), dcrBottom, dcrUpper)

	return
}
