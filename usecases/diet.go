package usecases

import (
	"errors"
	"fmt"
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

func countREE(
	gender models.GENDER, weight float32, height float32, age int) (result float32) {
	if gender == models.FEMALE {
		result = 655.1 + 13.8*weight + 5.0*height - 6.8*float32(age)
	} else if gender == models.MALE {
		result = 66.5 + 9.6*weight + 1.8*height - 4.7*float32(age)
	}

	return
}

func countCE(ree float32, af models.ACTIVITY_FACTOR) (result float32) {
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

func countDCR(ce float32, dietType *models.DietType) (result float32) {
	// fmt.Println(dietType)
	switch dietType.Operation {
	case models.MINUS:
		result = ce - dietType.Amount
	case models.PLUS:
		result = ce - dietType.Amount
	}
	return
}

func (u *dietUsecase) findDCR(body models.DietPlanBody) (dcr float32, err error) {
	dietType, err := u.dietTypeRepository.GetByID(body.DietTypeID)
	if err != nil {
		return
	}
	if dietType == nil {
		err = errors.New("Diet Type not available")
		return
	}
	ree := countREE(body.Gender, body.Weight, body.Height, body.Age)
	ce := countCE(ree, body.ActivityFactor)
	dcr = countDCR(ce, dietType)

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

func subsetSumRangex(input []models.Menu, n int, a int, b int) (res [][3]models.Menu) {
	res = make([][3]models.Menu, 0)

	sort.Slice(input, func(i, j int) bool {
		return input[i].Calorie < input[j].Calorie
	})
	// fmt.Println()
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			first := input[i].Calorie + input[j].Calorie
			l := 0
			r := n - 1
			for l <= r {
				m := l + (r-l)/2
				// fmt.Printf("l: %d r:%d m: %d\n", l, r, m)

				second := first + input[m].Calorie
				if second < a {
					l = m + 1
				} else if second > b {
					r = m - 1
				} else if second >= a {
					fmt.Println("MASUK")
					for second <= b && m < n {
						fmt.Printf("m: %d r: %d l: %d\n", m, r, l)
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

func subsetSumRange(inp []models.Menu, n int, a int, b int) (res [][3]models.Menu) {
	res = make([][3]models.Menu, 0)

	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n-2; k++ {
				cc := inp[i].Calorie + inp[j].Calorie + inp[k].Calorie
				if (cc >= a) && (cc <= b) {
					res = append(res, [3]models.Menu{inp[i], inp[j], inp[k]})
				}
			}
		}
	}

	return
}

func (u *dietUsecase) FindDietPlan(body models.DietPlanBody) (err error) {
	duration := body.Duration
	dcr, err := u.findDCR(body)
	if err != nil {
		return
	}
	dcrUpper := int(math.Round(float64(dcr))) + 100
	dcrBottom := int(math.Round(float64(dcr))) - 100
	// fmt.Println("MASUK")
	menus, err := u.menuRepository.GetAll()
	if err != nil {
		return
	}

	// fmt.Printf("Upper DCR: %d", dcrUpper)
	// fmt.Printf("Lower DCR: %d", dcrBottom)
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

	var menuAll []models.MenuPerDay
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

	dietPlan := models.DietPlan{
		Type:      dietType,
		Duration:  duration,
		Weight:    body.Weight,
		StartDate: startDate,
		EndDate:   endDate,
		MenusAll:  menuAll,
	}

	err = u.userUsecase.UpdateDietPlan(body.UserID, &dietPlan)

	return
}
