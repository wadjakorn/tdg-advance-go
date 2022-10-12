package main

import (
	"net/http"
	"os"
	"restapi2/users"

	"github.com/labstack/echo/v4"
)

func main() {
	port := os.Getenv("PORT")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, User API!")
	})

	repo := users.NewUserRepository()
	us := users.NewUserService(*repo)
	e.GET("/users", users.GetUserHandler(us))

	e.Logger.Fatal(e.Start("127.0.0.1:" + port))
}
