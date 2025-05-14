package service

import (
	"todo-api/terminal/models"
	"todo-api/terminal/repository"
)

type UserService interface {
	CreateUser(user *models.User) (*models.User, error)
	GetAllUsers() ([]*models.User, error)
	GetUserByID(id string) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(id string) error
}
type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}
func (s *userService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.CreateUser(user)
}
func (s *userService) GetAllUsers() ([]*models.User, error) {
	return s.repo.GetAllUsers()
}
func (s *userService) GetUserByID(id string) (*models.User, error) {
	return s.repo.GetUserByID(id)
}
func (s *userService) UpdateUser(user *models.User) (*models.User, error) {
	return s.repo.UpdateUser(user)
}
func (s *userService) DeleteUser(id string) error {
	return s.repo.DeleteUser(id)
}
