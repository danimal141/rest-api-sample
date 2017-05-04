package router

import (
	"github.com/danimal141/rest-api-sample/api/v1"
	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e1 := e.Group("/api/v1")
	e1.GET("/users", v1.UsersIndex)
	e1.GET("/users/:user_id", v1.UsersShow)
	return e
}
