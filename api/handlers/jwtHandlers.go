package handlers

func MainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	log.Println("User Name: ", claims["name"], "User ID: ", claims["jti"])

	return c.String(http.StatusOK, "you are ont top secret jwt page!")
}