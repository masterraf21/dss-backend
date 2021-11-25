package usecases

import (
	"errors"

	"github.com/masterraf21/dss-backend/models"
	authUtil "github.com/masterraf21/dss-backend/utils/auth"
)

type userUsecase struct {
	userRepo models.UserRepository
}

func NewUserUsecase(usr models.UserRepository) models.UserUsecase {
	return &userUsecase{
		userRepo: usr,
	}
}

func (u *userUsecase) Register(body models.UserBody) (id uint32, err error) {
	hash, err := authUtil.GeneratePassword(body.Password)
	if err != nil {
		return
	}
	user := models.User{
		Username:          body.Username,
		EncryptedPassword: hash,
		Email:             body.Email,
		PhoneNumber:       body.PhoneNumber,
	}

	id, err = u.userRepo.Store(&user)
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) GetByID(id uint32) (res *models.User, err error) {
	res, err = u.userRepo.GetByID(id)
	return
}

func (u *userUsecase) Login(body models.UserBody) (err error) {
	user, err := u.userRepo.GetByUsername(body.Username)
	var ok bool
	if user != nil {
		ok, err = authUtil.ComparePassword(user.EncryptedPassword, body.Password)
		if err != nil {
			return
		}
		if !ok {
			err = errors.New("Wrong Password")
			return
		}
	} else {
		err = errors.New("No username at system")
	}

	return
}
