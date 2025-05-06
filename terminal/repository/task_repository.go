package repository

import (
	"todo-api/terminal/models"

	"gorm.io/gorm"
)

type TaskRepository interface {
	PostTasks(task *models.Task) (*models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id string) (*models.Task, error)
	PatchTasksId(task *models.Task) (*models.Task, error)
	Delete(task *models.Task) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) PostTasks(task *models.Task) (*models.Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(id string) (*models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepository) PatchTasksId(task *models.Task) (*models.Task, error) {
	if err := r.db.Save(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *taskRepository) Delete(task *models.Task) error {
	return r.db.Delete(task).Error
}
