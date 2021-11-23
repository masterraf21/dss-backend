package models

// Counter represents counter for
type Counter struct {
	UserID uint32 `bson:"id_user" json:"id_user"`
}

// CounterRepository repo for counter
type CounterRepository interface {
	Get(collectionName string, identifier string) (uint32, error)
}
