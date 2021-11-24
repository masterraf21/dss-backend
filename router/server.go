package router

import (
	"context"

	"github.com/masterraf21/dss-backend/models"
)

type Server struct {
	DietUsecase models.DietUsecase
	MenuUsecase models.MenuUsecase
	UserUsecase models.UserUsecase
}

// NewServer will create handler
func NewServer(ctx context.Context) *Server {
	return &Server{}
}
