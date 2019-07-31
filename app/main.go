package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

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
	return c.String(http.StatusOK, "you are on the secret main page!")
}

func mainCookie(c echo.Context) error {
	return c.String(http.StatusOK, "you are on the secret cookie main page!")
}

func login(c echo.Context) error {
	username := c.QueryParam("username")
	password := c.QueryParam("password")

	if username == "tarou" && password == "1234" {
		cookie := &http.Cookie{}
		cookie.Name = "sessionID"
		cookie.Value = "some_string"
		cookie.Expires = time.Now().Add(48 * time.Hour)
		c.SetCookie(cookie)

		return c.String(http.StatusOK, "You were logged in!")
	}

	return c.String(http.StatusUnauthorized, "Your username and passward were wrong!")
}

func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "BlueBot/1.0")
		c.Response().Header().Set("NotReallyHeader", "thisHaveNoMeaning")
		return next(c)
	}
}

func checkCookie(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("sessionID")
		if err != nil {
			if strings.Contains(err.Error(), "named cookie not present") {
				return c.String(http.StatusUnauthorized, "You dont have the right cookie")
			}
			log.Println(err)
			return err
		}
		if cookie.Value == "some_string" {
			return next(c)
		}

		return c.String(http.StatusUnauthorized, "You dont have the right cookie, cookie")
	}
}

func main() {
	e := echo.New()

	e.Use(ServerHeader)

	adminGroup := e.Group("/admin")
	cookieGroup := e.Group("/cookie")
	// document通りでも警告出る。
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "tarou" && password == "1234" {
			return true, nil
		}

		return false, nil
	}))

	cookieGroup.Use(checkCookie)

	adminGroup.GET("/main", mainAdmin)

	cookieGroup.GET("/main", mainCookie)

	e.GET("/login", login)
	e.GET("/", getHelloWorld)
	e.GET("/cats/:data", getCats)
	e.POST("/cats", addCat)
	e.POST("/dogs", addDog)
	e.POST("/monkeys", addMonkey)
	e.Logger.Fatal(e.Start(":8080"))
}
