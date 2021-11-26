package usecases

import (
	"math"
	"sort"
	"time"

	"github.com/masterraf21/dss-backend/models"
)

type dietUsecase struct {
	dietTypeRepository models.DietTypeRepository
	menuRepository     models.MenuRepository
	userUsecase        models.UserUsecase
}

// NewDietUsecase will create usecase
func NewDietUsecase(
	dtsr models.DietTypeRepository, mr models.MenuRepository, usr models.UserUsecase,
) models.DietUsecase {
	return &dietUsecase{
		dietTypeRepository: dtsr,
		menuRepository:     mr,
		userUsecase:        usr,
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
			Calorie: menu.Calorie,
		})
	}

	return
}

func subsetSumRange(input []models.Menu, n int, a int, b int) (res [][3]models.Menu) {
	res = make([][3]models.Menu, 0)

	sort.Slice(input, func(i, j int) bool {
		return input[i].Calorie < input[j].Calorie
	})

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			first := input[i].Calorie + input[j].Calorie
			l := 0
			r := n - 1
			for l <= r {
				m := l + (r-1)/2
				second := first + input[m].Calorie
				if second < a {
					l = m + 1
				} else if second > b {
					r = m - 1
				} else if second >= a {
					for second <= b {
						second = first + input[m].Calorie
						if m != i && m != j {
							res = append(res, [3]models.Menu{input[i], input[j], input[m]})
						}
						m++
					}
				}
			}

		}
	}

	return
}

func (u *dietUsecase) FindDietPlan(body models.DietPlanBody, userID uint32) (res *models.DietPlan, err error) {
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

	// user, err := u.userRepository.GetByID(userID)
	// if err != nil {
	// 	return
	// }

	// data := extractIDCalorie(menus)

	planMenusRaw := subsetSumRange(menus, len(menus), dcrBottom, dcrUpper)
	var planMenus [][3]models.Menu

	if len(planMenusRaw) < duration {
		padNumber := duration - len(planMenusRaw)
		planMenus = make([][3]models.Menu, duration)
		if len(planMenusRaw) > 0 {
			for i := 0; i < duration; i++ {
				if i >= len(planMenusRaw) {
					planMenus[i] = planMenusRaw[i-padNumber]
				} else {
					planMenus[i] = planMenusRaw[i]
				}
			}
		}
	} else {
		planMenus = make([][3]models.Menu, duration)
		for i := 0; i < duration; i++ {
			planMenus[i] = planMenusRaw[i]
		}
	}

	dietType, err := u.dietTypeRepository.GetByID(body.DietTypeID)
	if err != nil {
		return
	}

	menuAll := []models.MenuPerDay{}
	menuAll = make([]models.MenuPerDay, 0)

	FORMAT := "2006-January-02"
	now := time.Now().Local()
	startDate := now.Format(FORMAT)
	endDate := now.Add(time.Hour * time.Duration(24*duration)).Format(FORMAT)

	for i := 0; i < len(planMenus); i++ {
		dayPlus := now.Add(time.Hour * time.Duration(24*i))
		date := dayPlus.Format(FORMAT)
		menuAll = append(menuAll, models.MenuPerDay{
			Date: date,
			Menu: planMenus[i],
		})
	}

	// err = u.userRepository.UpdateArbitrary(userID,"diet_plan",)
	err = u.userUsecase.UpdateDietPlan(userID, &models.DietPlan{
		Type:      dietType,
		Duration:  duration,
		Weight:    body.Weight,
		StartDate: startDate,
		EndDate:   endDate,
		MenusAll:  menuAll,
	})

	return
}
