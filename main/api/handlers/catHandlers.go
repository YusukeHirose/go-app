package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func GetCats(c echo.Context) error {
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

func AddCat(c echo.Context) error {
	cat := Cat{}

	// 最後に呼ぶ処理をdeferで記述
	defer c.Request().Body.Close()

	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("Faild reading the request body for addCats: %s", err)
		return c.String(http.StatusInternalServerError, "err")
	}

	err = json.Unmarshal(b, &cat)
	if err != nil {
		log.Printf("Faild unmarshaling in addCats: %s", err)
		return c.String(http.StatusInternalServerError, "")
	}

	log.Printf("this is cat: %#v", cat)
	return c.String(http.StatusOK, "We got the cat!")
}
