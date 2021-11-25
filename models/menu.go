package models

type LABEL string

const (
	HALAL     LABEL = "halal"
	VEGAN     LABEL = "vegan"
	NON_DAIRY LABEL = "non_dairy"
	NON_NUTS  LABEL = "non_nuts"
)

type (
	Menu struct {
		ID          uint32   `bson:"id_menu" json:"id_menu"`
		Name        string   `bson:"name" json:"name"`
		Calorie     int      `bson:"calorie" json:"calorie"`
		Recipe      string   `bson:"recipe" json:"recipe"`
		Ingredients []string `bson:"ingredients" json:"ingredients"`
		PictureURL  string   `bson:"picture_url" json:"picture_urL"`
		Labels      []string `bson:"labels" json:"labels"`
	}

	MenuBody struct {
		Name        string   `json:"name"`
		Calorie     int      `json:"calorie"`
		Recipe      string   `json:"recipe"`
		Ingredients []string `json:"ingredients"`
		PictureURL  string   `json:"picture_url"`
		Labels      []string `json:"labels"`
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
