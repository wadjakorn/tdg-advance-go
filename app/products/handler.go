package products

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (service *ProductService) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, service.products())
	}
}
