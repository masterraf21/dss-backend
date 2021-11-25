package router

import (
	"context"

	"github.com/masterraf21/dss-backend/models"
	repoMongo "github.com/masterraf21/dss-backend/repositories/mongodb"
	"github.com/masterraf21/dss-backend/usecases"
	"github.com/masterraf21/dss-backend/utils/mongodb"
)

type Server struct {
	DietTypeUsecase models.DietTypeUsecase
	DietUsecase     models.DietUsecase
	MenuUsecase     models.MenuUsecase
	UserUsecase     models.UserUsecase
}

// NewServer will create handler
func NewServer(ctx context.Context) *Server {
	instance := mongodb.ConfigureMongo()
	counterRepo := repoMongo.NewCounterRepo(instance)
	menuRepo := repoMongo.NewMenuRepo(instance, counterRepo)
	dietTypeRepo := repoMongo.NewDietTypeRepository(instance, counterRepo)

	dietUsecase := usecases.NewDietUsecase(dietTypeRepo)
	dietTypeUsecase := usecases.NewDietTypeUsecase(dietTypeRepo)
	menuUsecase := usecases.NewMenuUsecase(menuRepo)

	return &Server{
		DietTypeUsecase: dietTypeUsecase,
		DietUsecase:     dietUsecase,
		MenuUsecase:     menuUsecase,
	}
}
