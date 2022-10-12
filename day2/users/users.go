package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserResponse struct {
	Message string `json:"message"`
}

func GetUserHandler(service *UserService) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := UserResponse{service.GetAll()}
		return c.JSON(http.StatusOK, r)
	}
}

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetSth() (string, error) {
	return "Call get user", fmt.Errorf("TODO next")
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetAll() string {
	result, _ := us.repo.GetSth()
	return result
}
