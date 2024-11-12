package service

import (
	"github.com/arjunmalhotra1/go-template/internal/model"
	"github.com/arjunmalhotra1/go-template/internal/repository"
)

type TaskService struct {
	Repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{Repo: repo}
}

func (s *TaskService) CreateTask(task model.Task) (int, error) {
	return s.Repo.CreateTask(task)
}

func (s *TaskService) GetTasks() ([]model.Task, error) {
	return s.Repo.GetTasks()
}
