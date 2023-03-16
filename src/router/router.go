package router

import (
	"echo-example/src/api"
	"echo-example/src/api/middlewares"
	"echo-example/src/vo"
	"net/http"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	//create groups
	adminGroup := e.Group("/admin")

	//set all middlewares
	middlewares.SetMainMiddleWares(e)
	middlewares.SetAdminMiddlewares(adminGroup)

	// Route / to handler function
	e.GET("/health-check", func(c echo.Context) error {
		resp := vo.HealthCheckResponse{
			Message: "Everything is good!",
		}
		return c.JSON(http.StatusOK, resp)
	})

	//set main routes
	api.MainGroup(e)
	//set groupRoutes
	api.AdminGroup(adminGroup)

	return e
}
