package router

import (
	"github.com/masterraf21/dss-backend/models"
	repoMongo "github.com/masterraf21/dss-backend/repositories/mongodb"
	"github.com/masterraf21/dss-backend/usecases"
	"github.com/masterraf21/dss-backend/utils/mongodb"
)

type Handler struct {
	DietTypeUsecase models.DietTypeUsecase
	DietUsecase     models.DietUsecase
	MenuUsecase     models.MenuUsecase
	UserUsecase     models.UserUsecase
}

// NewServer will create handler
func NewHandler() *Handler {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	menuRepo := repoMongo.NewMenuRepo(instance, counterRepo)
	dietTypeRepo := repoMongo.NewDietTypeRepository(instance, counterRepo)
	userRepo := repoMongo.NewUserRepo(instance, counterRepo)

	userUsecase := usecases.NewUserUsecase(userRepo)
	dietUsecase := usecases.NewDietUsecase(dietTypeRepo, menuRepo, userUsecase)
	dietTypeUsecase := usecases.NewDietTypeUsecase(dietTypeRepo)
	menuUsecase := usecases.NewMenuUsecase(menuRepo)

	return &Handler{
		DietTypeUsecase: dietTypeUsecase,
		DietUsecase:     dietUsecase,
		MenuUsecase:     menuUsecase,
		UserUsecase:     userUsecase,
	}
}
