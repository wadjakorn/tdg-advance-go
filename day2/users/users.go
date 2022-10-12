package users

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
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

type IRepository interface {
	GetSth() (string, error)
}

type UserRepository struct {
	*mongo.Client
}

func NewUserRepository(client *mongo.Client) UserRepository {
	return UserRepository{Client: client}
}

func (r *UserRepository) GetSth() (string, error) {
	return "TODO next", fmt.Errorf("TODO next")
}

type UserService struct {
	repo IRepository
}

func NewUserService(repo IRepository) *UserService {
	return &UserService{repo: repo}
}

func (us *UserService) GetAll() string {
	result, _ := us.repo.GetSth()
	return result
}
