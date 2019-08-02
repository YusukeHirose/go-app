package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!!")
}
