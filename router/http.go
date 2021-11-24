package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/masterraf21/dss-backend/configs"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

func (h *Server) Start() *echo.Echo {
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

	listenerPort := fmt.Sprintf(":%d", configs.Server.Port)
	e.Logger.Fatal(e.Start(listenerPort))

	return e
}
