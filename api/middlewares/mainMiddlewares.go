package middlewares

import (
	"github.com/labstack/echo"
)

func SetMainMiddlewares(e *echo.Echo) {
	e.Use(ServerHeader)
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("NotReallyHeader", "thisHaveNoMeaning")
		return next(c)
	}
}
