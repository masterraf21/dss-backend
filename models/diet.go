package models

type (
	GENDER                int
	ACTIVITY_FACTOR       int
	CALCULATION_OPERATION int
)

const (
	FEMALE GENDER = iota
	MALE
)

const (
	SEDENTARY ACTIVITY_FACTOR = iota
	LIGHTLY_ACTIVE
	MODERATELY_ACTIVE
	VERY_ACTIVE
	EXTRA_ACTIVE
)

const (
	PLUS CALCULATION_OPERATION = iota
	MINUS
)

type (
	DietPlan struct {
		Type      *DietType    `json:"type" bson:"type"`
		Duration  int          `json:"duration" bson:"duration"`
		Weight    float32      `json:"weight" bson:"weight"`
		StartDate string       `json:"start_date" bson:"start_date"`
		EndDate   string       `json:"end_date" bson:"end_date"`
		Menu      []MenuPerDay `json:"menu" bson:"menu"`
		Calorie   float32
	}

	DietType struct {
		ID          uint32                `json:"id_diet_type" bson:"id_diet_type"`
		Name        string                `json:"name" bson:"name"`
		Description string                `json:"description" bson:"description"`
		Operation   CALCULATION_OPERATION `json:"operation" bson:"operation"`
		Amount      float32               `json:"amount" bson:"amount"`
	}

	DietTypeBody struct {
		Name        string                `json:"name"`
		Description string                `json:"description"`
		Operation   CALCULATION_OPERATION `json:"operation"`
		Amount      float32               `json:"amount"`
	}

	MenuPerWeek struct {
		MenuPerDays [7]MenuPerDay `json:"menu_per_day" bson:"menu_per_day"`
	}
	MenuPerDay struct {
		Date   string `json:"data"`
		First  *Menu  `json:"first" bson:"first"`
		Second *Menu  `json:"second" bson:"second"`
		Third  *Menu  `json:"third" bson:"third"`
	}

	DietPlanBody struct {
		Weight         float32         `json:"weight"`
		Height         float32         `json:"height"`
		Gender         GENDER          `json:"gender"`
		Age            int             `json:"age"`
		ActivityFactor ACTIVITY_FACTOR `json:"activity_factor"`
		DietTypeID     uint32          `json:"diet_type_id"`
	}

	DietUsecase interface {
		// REE: Resting Enery Expenditure
		CountREE(gender GENDER, weight float32, height float32, age int) float32
		// CE: Calorie Expenditure
		CountCE(ree float32, af ACTIVITY_FACTOR) float32
		// DCR: Daily Calorie Recommendation
		CountDCR(ce float32, dietType DietType) float32
		FindMenu(dcr float32, duration int)
	}

	DietTypeUsecase interface {
		Create(body DietTypeBody) (uint32, error)
		GetAll() ([]DietType, error)
		GetByID(id uint32) (*DietType, error)
	}

	// DietRepository interface{}

	DietTypeRepository interface {
		Store(dietType *DietType) (uint32, error)
		BulkStore(dietTypes []*DietType) ([]uint32, error)
		GetAll() ([]DietType, error)
		GetByID(id uint32) (*DietType, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}
)
