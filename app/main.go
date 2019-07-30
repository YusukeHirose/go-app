package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Cat struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Dog struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Monkey struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

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

func addCat(c echo.Context) error {
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

func addDog(c echo.Context) error {
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

func addMonkey(c echo.Context) error {
	monkey := Monkey{}

	err := c.Bind(&monkey)
	if err != nil {
		log.Printf("Faild processing  addMonkey: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	log.Printf("this is dog: %#v", monkey)
	return c.String(http.StatusOK, "We got the monkey!")
}

func mainAdmin(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret main pate!")
}

func main() {
	e := echo.New()

	g := e.Group("/admin")
	// document通りでも警告出る。
	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	g.GET("/main", mainAdmin)

	e.GET("/", getHelloWorld)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/monkeys", addMonkey)
	e.Logger.Fatal(e.Start(":8080"))
}
