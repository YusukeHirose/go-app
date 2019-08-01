package api

func CookieGroup(g *echo.Group) {
	g.GET("/main", handlers.MainCookie())
}