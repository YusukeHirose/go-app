package handlers

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Monkey struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddMonkey(c echo.Context) error {
	monkey := Monkey{}

	err := c.Bind(&monkey)
	if err != nil {
		log.Printf("Faild processing  addMonkey: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is dog: %#v", monkey)
	return c.String(http.StatusOK, "We got the monkey!")
}
