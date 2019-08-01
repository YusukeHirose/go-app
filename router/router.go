package router

import "github.com/labstack/echo"

func Init() *echo.Echo {
	e := echo.New()

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	jwtGroup := e.Group("/jwt")

	return e
}
