package repository

import (
	"database/sql"

	"github.com/tunardev/auth-user/pkg/entity"
)

type TaskRepository interface {
	All(id int64) ([]entity.Task, error)
	Create(task entity.Task) (int64, error)
	Get(id string, userId int64) (entity.Task, error)
	Delete(id string, userId int64) error
	Update(id string, task entity.Task) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return taskRepository{db}
}

func (repo taskRepository) All(id int64) ([]entity.Task, error) {
	rows, err := repo.db.Query("SELECT * FROM tasks WHERE user_id = $1", id)
	if err != nil {
		return []entity.Task{}, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.UserId, &task.Status)
		if err != nil {
			return []entity.Task{}, err
		}
		tasks = append(tasks, task)
	}

	err = rows.Err()
	if err != nil {
		return []entity.Task{}, err
	}

	return tasks, nil
}

func (repo taskRepository) Get(id string, userId int64) (entity.Task, error) {
	var task entity.Task
	err := repo.db.QueryRow("SELECT * FROM tasks WHERE id = $1 AND user_id = $2", id, userId).Scan(&task.ID, &task.Title, &task.Content, &task.UserId, &task.Status)
	if err != nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (repo taskRepository) Create(task entity.Task) (int64, error) {
	var id int64
	err := repo.db.QueryRow("INSERT INTO tasks (title, content, status, user_id) VALUES ($1, $2, $3, $4) RETURNING id", task.Title, task.Content, task.Status, task.UserId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (repo taskRepository) Delete(id string, userId int64) error {
	_, err := repo.db.Exec("DELETE FROM tasks WHERE id = $1 AND user_id = $2", id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (repo taskRepository) Update(id string, task entity.Task) error {
	_, err := repo.db.Exec("UPDATE tasks SET status = $1, title = $2, content = $3 WHERE id = $4 AND user_id = $5", task.Status, task.Title, task.Content, id, task.UserId)
	if err != nil {
		return err
	}

	return nil

}