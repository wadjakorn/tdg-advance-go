package main

import (
	"net/http"
	"os"
	"product-api/products"

	"github.com/labstack/echo/v4"
)

func main() {
	/* if use flag instead of env, use this (also import flag) */
	// var port = flag.String("port", "8080", "server")
	// flag.Parse()

	port := os.Getenv("PORT")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Product API!")
	})

	pService := products.NewProductService()
	e.GET("/products", pService.GetAll())

	e.Logger.Fatal(e.Start("127.0.0.1:" + port))

	/* if use flag instead of env, use *port */
	// e.Logger.Fatal(e.Start("127.0.0.1:" + *port))
}
