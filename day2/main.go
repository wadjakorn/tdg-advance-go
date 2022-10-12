package main

import (
	"flag"
	"net/http"
	"os"
	"restapi2/lib"
	"restapi2/users"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var url = flag.String("url", "mongodb://root:example@localhost:27017/", "MongoDB Host")

func main() {
	port := os.Getenv("PORT")

	e := echo.New()
	e.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, User API!")
	})

	mongoClient := lib.NewMongoClient(*url)
	repo := users.NewUserRepository(mongoClient)
	us := users.NewUserService(&repo)
	e.GET("/users", users.GetUserHandler(us))

	e.Logger.Fatal(e.Start("127.0.0.1:" + port))
}
