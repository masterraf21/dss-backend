package models

type (
	User struct {
		Username          string `bson:"username" json:"username"`
		EncryptedPassword string `bson:"encrypted_password" json:"encrypted_password"`
		Email             string `bson:"email" json:"email"`
		PhoneNumber       string `bson:"phone_number" json:"phone_number"`
	}

	UserUsecase interface{}

	UserRepository interface{}
)
