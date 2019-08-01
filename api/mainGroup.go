package api

import "github.com/labstack/echo"

func MainGroup(e *echo.Echo) {
	e.GET("/login", handlers.Login())
	e.GET("/", handlers.getHelloWorld())
	e.GET("/cats/:data", handlers.GetCats())
	e.POST("/cats", handlers.AddCat())
	e.POST("/dogs", handlers.AddDog())
	e.POST("/monkeys", handlers.AddMonkey())
}
