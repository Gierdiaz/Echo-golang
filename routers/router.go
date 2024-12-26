package routers

import (
	"github.com/Gierdiaz/Echo-golang/handlers"
	"github.com/labstack/echo/v4"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", handlers.GetUsers)
	e.POST("/users", handlers.CreateUser)

	return e
}
