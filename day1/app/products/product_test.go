// go:build e2e

package products_test

import (
	"net/http"
	"net/http/httptest"
	"product-api/products"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSuccess(t *testing.T) {
	e := echo.New()
	service := products.NewProductService()
	e.GET("/products", products.GetAll(service))
	httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	assert.Equal(t, http.StatusOK, rec.Code)
}
