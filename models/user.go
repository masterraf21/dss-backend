package models

type (
	User struct {
		ID                uint32 `bson:"id_user" json:"id_user"`
		Username          string `bson:"username" json:"username"`
		EncryptedPassword string `bson:"encrypted_password" json:"encrypted_password"`
		Email             string `bson:"email" json:"email"`
		PhoneNumber       string `bson:"phone_number" json:"phone_number"`
		// OnDiet            bool      `bson:"on_diet" json:"on_diet"`
		CurrentPlan *DietPlan `bson:"current_plan" json:"current_plan"`
	}

	UserUsecase interface {
		RegisterUser(user User) (uint32, error)
	}

	UserRepository interface {
		Store(user *User) (uint32, error)
		BulkStore(users []User) ([]uint32, error)
		GetAll() ([]User, error)
		GetByID(id uint32) (*User, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}
)
