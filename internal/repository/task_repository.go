package repository

import (
	"database/sql"

	"github.com/arjunmalhotra1/go-template/internal/model"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task model.Task) (int, error) {
	query := `INSERT INTO tasks (title, description, completed, created_at, updated_at)
              VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id`
	var id int
	err := r.DB.QueryRow(query, task.Title, task.Description, task.Completed).Scan(&id)
	return id, err
}

func (r *TaskRepository) GetTasks() ([]model.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, description, completed, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []model.Task{}
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
