package main

import (
	"net/http"
	"os"
	"time"

	// "github.com/gokomodo/library/panics/middleware/echohandler"
	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/masterraf21/dss-backend/configs"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
	"github.com/masterraf21/dss-backend/utils/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
)

// Server represents server
type Server struct {
	Instance    *mongo.Database
	Port        string
	ServerReady chan bool
}

func main() {
	instance := mongodb.ConfigureMongo()
	serverReady := make(chan bool)
	server := Server{
		Instance:    instance,
		Port:        configs.Server.Port,
		ServerReady: serverReady,
	}
	server.Start()
}

func (s *Server) Start() {
	port := configs.Server.Port
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(
		// middleware.Logger,
		// echohandler.Capture,
		em.CORSWithConfig(em.CORSConfig{
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
		// helper.Recover(),
		// ddEcho.Middleware(),
	)
	e.Debug = os.Getenv("DEBUG") == "1"

	e.GET("/info", func(c echo.Context) (err error) {
		return httpUtil.NewResponse(
			http.StatusOK,
			"",
			map[string]interface{}{
				"version": configs.Build.Version,
				"commit":  configs.Build.Commit,
				"upsince": time.Now().Format(time.RFC3339),
				"uptime":  time.Since(time.Now()).String(),
			},
		).WriteResponse(c)
	})
}
