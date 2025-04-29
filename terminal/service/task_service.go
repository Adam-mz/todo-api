package service

import (
	"todo-api/terminal/models"
	"todo-api/terminal/repository"
)

type TaskService interface {
	CreateTask(task *models.Task) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	UpdateTask(task *models.Task) (*models.Task, error)
	DeleteTask(id string) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) CreateTask(task *models.Task) (*models.Task, error) {
	return s.repo.Create(task)
}

func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTaskByID(id string) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) UpdateTask(task *models.Task) (*models.Task, error) {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(task)
}
