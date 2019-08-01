package middlewares

func SetCookieMiddlewares(g *echo.Echo) {
	cookieGroup.Use(checkCookie)
}