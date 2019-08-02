package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret main page!")
}
