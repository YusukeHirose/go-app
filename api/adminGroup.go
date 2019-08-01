package api

func AdminGroup(g *echo.Group) {
	g.GET("/main", handlers.MainAdmin())
}