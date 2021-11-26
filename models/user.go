package models

type (
	User struct {
		ID                uint32 `bson:"id_user" json:"id_user"`
		Username          string `bson:"username" json:"username"`
		EncryptedPassword string `bson:"encrypted_password" json:"encrypted_password"`
		Email             string `bson:"email" json:"email"`
		PhoneNumber       string `bson:"phone_number" json:"phone_number"`
		// OnDiet            bool      `bson:"on_diet" json:"on_diet"`
		CurrentPlan *DietPlan `bson:"diet_plan" json:"diet_plan"`
	}

	LoginRespose struct {
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
		ExpireIn     int64  `json:"expire_in"`
		UserID       uint32 `json:"user_id"`
	}

	UserBody struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	LoginBody struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	UserUsecase interface {
		Register(body UserBody) (uint32, error)
		GetByID(id uint32) (*User, error)
		Login(body LoginBody) (*LoginRespose, error)
		UpdateDietPlan(id uint32, plan *DietPlan) error
	}

	UserRepository interface {
		Store(user *User) (uint32, error)
		BulkStore(users []*User) ([]uint32, error)
		GetAll() ([]User, error)
		GetByID(id uint32) (*User, error)
		GetByUsername(username string) (*User, error)
		UpdateArbitrary(id uint32, key string, value interface{}) error
	}
)
