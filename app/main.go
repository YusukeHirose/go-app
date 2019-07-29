package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!!")
}

func getCats(c echo.Context) error {
	catName := c.QueryParam("name")
	catType := c.QueryParam("type")

	dataType := c.Param("data")

	if dataType == "string" {
		return c.String(http.StatusOK, fmt.Sprintf("Cat name is: %s\nand his type is: %s\n", catName, catType))
	}

	if dataType == "json" {
		return c.JSON(http.StatusOK, map[string]string{
			"name":    catName,
			"catType": catType,
		})
	}

	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": "You need to lets us know if you want json or string data",
	})

}

func main() {
	e := echo.New()
	e.GET("/", getHelloWorld)
	e.GET("/cats/:data", getCats)
	e.Logger.Fatal(e.Start(":8080"))
}
