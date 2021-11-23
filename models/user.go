package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	User struct {
		ID                uint32 `bson:"id_user" json:"id_user"`
		Username          string `bson:"username" json:"username"`
		EncryptedPassword string `bson:"encrypted_password" json:"encrypted_password"`
		Email             string `bson:"email" json:"email"`
		PhoneNumber       string `bson:"phone_number" json:"phone_number"`
		OnDiet            bool   `bson:"on_diet" json:"on_diet"`
	}

	// UserBody struct {
	// 	Username    string `json:"username"`
	// 	Password    string `json:"encrypted_password"`
	// 	Email       string `json:"email"`
	// 	PhoneNumber string `json:"phone_number"`
	// }

	UserUsecase interface {
		RegisterUser(user User) (uint32, error)
	}

	UserRepository interface {
		Store(user *User) (primitive.ObjectID, error)
	}
)
