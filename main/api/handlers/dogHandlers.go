package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func AddDog(c echo.Context) error {
	dog := Dog{}

	defer c.Request().Body.Close()

	err := json.NewDecoder(c.Request().Body).Decode(&dog)
	if err != nil {
		log.Printf("Faild processing  addDog: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is dog: %#v", dog)
	return c.String(http.StatusOK, "We got the dog!")
}
