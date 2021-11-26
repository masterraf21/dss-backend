package usecases

import (
	"errors"

	"github.com/masterraf21/dss-backend/auth"
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

func (u *userUsecase) Login(body models.LoginBody) (res *models.LoginRespose, err error) {
	user, err := u.userRepo.GetByUsername(body.Username)
	if err != nil {
		err = errors.New("internal_error")
		return
	}

	// fmt.Println("%v\n", user)
	var ok bool
	var token *auth.TokenDetails

	if user != nil {
		ok, err = authUtil.ComparePassword(user.EncryptedPassword, body.Password)
		if err != nil {
			return
		}
		if !ok {
			err = errors.New("Wrong Password")
			return
		}
		token, err = auth.GenerateToken(user)
		if err != nil {
			return
		}

		res = &models.LoginRespose{
			Token:        token.AccessToken,
			RefreshToken: token.RefreshToken,
			ExpireIn:     token.AtExpires,
			UserID:       user.ID,
		}
	} else {
		err = errors.New("No username at system")
		return
	}

	return
}

func (u *userUsecase) UpdateDietPlan(id uint32, plan *models.DietPlan) error {
	err := u.userRepo.UpdateArbitrary(id, "diet_plan", plan)
	return err
}
