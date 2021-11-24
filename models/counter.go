package models

// Counter represents counter for
type Counter struct {
	UserID     uint32 `bson:"id_user" json:"id_user"`
	MenuID     uint32 `bson:"id_menu" json:"id_menu"`
	DietTypeID uint32 `bson:"id_diet_type" json:"id_diet_type"`
}

// CounterRepository repo for counter
type CounterRepository interface {
	Get(collectionName string, identifier string) (uint32, error)
}
