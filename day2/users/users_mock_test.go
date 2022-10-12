// go:build integration
package users_test

import (
	"net/http"
	"net/http/httptest"
	"restapi2/users"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type MockUserRepository struct {
}

func (r *MockUserRepository) GetSth() (string, error) {
	return "xxx", nil
}

func TestSuccessWithGetAndMockRepo(t *testing.T) {
	// Start server
	e := echo.New()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	c := e.NewContext(req, rec)

	mockRepo := MockUserRepository{}
	service := users.NewUserService(&mockRepo)
	users.GetUserHandler(service)(c)

	// Assert
	assert.Equal(t, rec.Code, 200)
	assert.Contains(t, rec.Body.String(), "xxx")
}
