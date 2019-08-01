package middlewares

func SetAdminMiddlewares(g *echo.Group) {
	adminGroup.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	adminGroup.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "tarou" && password == "1234" {
			return true, nil
		}

		return false, nil
	}))
}