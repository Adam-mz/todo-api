package service

import (
	"todo-api/terminal/models"
	"todo-api/terminal/repository"
)

type TaskService interface {
	PostTasks(task *models.Task) (*models.Task, error)
	GetAllTasks() ([]models.Task, error)
	GetTaskByID(id string) (*models.Task, error)
	PatchTasksId(task *models.Task) (*models.Task, error)
	DeleteTask(id string) error
	GetTasksByUserID(userID uint) ([]models.Task, error)
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (s *taskService) PostTasks(task *models.Task) (*models.Task, error) {
	return s.repo.PostTasks(task)
}
func (s *taskService) GetTasksByUserID(userID uint) ([]models.Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
func (s *taskService) GetAllTasks() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetTaskByID(id string) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) PatchTasksId(task *models.Task) (*models.Task, error) {
	return s.repo.PatchTasksId(task)
}

func (s *taskService) DeleteTask(id string) error {
	task, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(task)
}
