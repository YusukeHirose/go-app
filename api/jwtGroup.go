package api

import "github.com/labstack/echo"

func JwtGroup(g *echo.Echo) {
	g.GET("/main", handlers.MainJwt())
}
