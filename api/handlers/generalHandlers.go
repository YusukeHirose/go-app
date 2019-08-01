package handlers

func GetHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!!")
}