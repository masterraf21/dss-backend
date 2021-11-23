package models

type (
	Menu struct {
		Name         string       `bson:"name" json:"name"`
		CalorieCount float32      `bson:"calorie_count" json:"calorie_count"`
		Recipe       string       `bson:"recipe" json:"recipe"`
		Ingredients  []Ingredient `bson:"ingredients" json:"ingredients"`
		PictureURL   string       `bson:"picture_url" json:"picture_urL"`
	}

	Ingredient struct {
		Name        string  `bson:"name" json:"name"`
		Measurement float32 `bson:"measurement" json:"measurement"`
		Unit        string  `bson:"unit" json:"unit"`
	}
)
