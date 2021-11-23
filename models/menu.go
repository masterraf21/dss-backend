package models

type (
	Menu struct {
		Name              string       `bson:"name" json:"name"`
		CalorieCount      float32      `bson:"calorie_count" json:"calorie_count"`
		Recipe            string       `bson:"recipe" json:"recipe"`
		Ingredients       []Ingredient `bson:"ingredients" json:"ingredients"`
		PictureURL        string       `bson:"picture_url" json:"picture_urL"`
		IngredientsLabels []string     `bson:"ingredients_label" json:"ingredients_label"`
	}

	Ingredient struct {
		Name         string   `bson:"name" json:"name"`
		Measurement  float32  `bson:"measurement" json:"measurement"`
		Unit         string   `bson:"unit" json:"unit"`
		Labels       []*Label `bson:"labels:" json:"labels"`
		LabelsString []string `bson:"labels_string" json:"labels_string"`
	}

	Label struct {
		Name        string `bson:"name" json:"name"`
		Description string `bson:"description" json:"description"`
	}

	MenuUsecase interface{}

	MenuRepository interface{}
)
