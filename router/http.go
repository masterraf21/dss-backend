package router

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
	"github.com/masterraf21/dss-backend/apis"
	"github.com/masterraf21/dss-backend/configs"
	httpUtil "github.com/masterraf21/dss-backend/utils/http"
)

func (h *Handler) HTTPStart() *echo.Echo {
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
			map[string]interface{}{
				"version": configs.Build.Version,
				"commit":  configs.Build.Commit,
				"upsince": time.Now().Format(time.RFC3339),
				"uptime":  time.Since(time.Now()).String(),
			},
		).WriteResponse(c)
	})

	menuRouter := apis.NewMenuRouter(h.MenuUsecase)
	menuRouter.Mount(e.Group("/menu"))

	dietTypeRouter := apis.NewDietTypeRouter(h.DietTypeUsecase)
	dietTypeRouter.Mount(e.Group("/diet_type"))

	userRouter := apis.NewUserRouter(h.UserUsecase)
	userRouter.Mount(e.Group("/user"))

	dietRouter := apis.NewDietRouter(h.DietUsecase)
	dietRouter.Mount(e.Group("/diet"))

	// userRouter := apis.NewUserRouter(h.UserUsecase)

	listenerPort := fmt.Sprintf("0.0.0.0:%d", configs.Server.Port)
	e.Logger.Fatal(e.Start(listenerPort))

	return e
}
