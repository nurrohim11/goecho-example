package api

import (
	"echo-example/src/handlers"

	"github.com/labstack/echo"
)

func MainGroup(e *echo.Echo) {
	e.GET("/cats/:data", handlers.GetCats)
	e.POST("/cats", handlers.AddCat)

}

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin)
}
