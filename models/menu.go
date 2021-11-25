package models

type (
	Menu struct {
		ID                uint32       `bson:"id_menu" json:"id_menu"`
		Name              string       `bson:"name" json:"name"`
		CalorieCount      float32      `bson:"calorie_count" json:"calorie_count"`
		Recipe            string       `bson:"recipe" json:"recipe"`
		Ingredients       []Ingredient `bson:"ingredients" json:"ingredients"`
		PictureURL        string       `bson:"picture_url" json:"picture_urL"`
		IngredientsLabels []string     `bson:"ingredients_label" json:"ingredients_label"`
	}

	MenuBody struct {
		Name         string       `json:"name"`
		CalorieCount float32      `json:"calorie_count"`
		Recipe       string       `json:"recipe"`
		Ingredients  []Ingredient `json:"ingredients"`
		PictureURL   string       `json:"picture_urL"`
	}

	Ingredient struct {
		Name        string   `bson:"name" json:"name"`
		Measurement float32  `bson:"measurement" json:"measurement"`
		Unit        string   `bson:"unit" json:"unit"`
		Labels      []string `bson:"labels:" json:"labels"`
		// LabelsString []string `bson:"labels_string" json:"labels_string"`
	}

	Label struct {
		Name        string `bson:"name" json:"name"`
		Description string `bson:"description" json:"description"`
	}

	MenuUsecase interface {
		Create(body MenuBody) (uint32, error)
		GetAll() ([]Menu, error)
		GetByID(id uint32) (*Menu, error)
	}

	MenuRepository interface {
		Store(menu *Menu) (uint32, error)
		BulkStore(menus []*Menu) ([]uint32, error)
		GetAll() ([]Menu, error)
		GetByID(id uint32) (*Menu, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}
)
