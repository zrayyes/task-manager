package repositories

import (
	"database/sql"
	"fmt"
	"sync"
	"time"
)

// TaskRepositoryInterface defines the methods for task repository
type TaskRepositoryInterface interface {
	CreateTask(task *Task) error
	GetTask(id string) (*Task, error)
	UpdateTask(id string, updatedTask *Task) error
	DeleteTask(id string) error
	GetAllTasks() []*Task
}

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository struct {
	mu    sync.Mutex
	tasks map[string]*Task
}

// Ensure TaskRepository implements TaskRepositoryInterface
var _ TaskRepositoryInterface = (*TaskRepository)(nil)

func NewTaskRepository() TaskRepositoryInterface {
	return &TaskRepository{
		tasks: make(map[string]*Task),
	}
}

func (r *TaskRepository) CreateTask(task *Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task.ID = fmt.Sprintf("%d", len(r.tasks)+1)
	task.Completed = false
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt
	r.tasks[task.ID] = task
	return nil
}

func (r *TaskRepository) GetTask(id string) (*Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, ErrNotFound
	}

	return task, nil
}

func (r *TaskRepository) UpdateTask(id string, updatedTask *Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, exists := r.tasks[id]
	if !exists {
		return ErrNotFound
	}

	updatedTask.CreatedAt = task.CreatedAt // Preserve original creation time
	updatedTask.UpdatedAt = time.Now()
	r.tasks[id] = updatedTask
	return nil
}

func (r *TaskRepository) DeleteTask(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return ErrNotFound
	}

	delete(r.tasks, id)
	return nil
}

func (r *TaskRepository) GetAllTasks() []*Task {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks := make([]*Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks
}

type TaskRepositoryPSQL struct {
	db *sql.DB
}

func NewTaskRepositoryPSQL(db *sql.DB) TaskRepositoryInterface {
	return &TaskRepositoryPSQL{
		db: db,
	}
}

func (r *TaskRepositoryPSQL) CreateTask(task *Task) error {
	query := `INSERT INTO tasks (title, body, completed, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.db.QueryRow(query, task.Title, task.Body, task.Completed, time.Now(), time.Now()).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepositoryPSQL) GetTask(id string) (*Task, error) {
	query := `SELECT id, title, body, completed, created_at, updated_at FROM tasks WHERE id = $1`
	task := &Task{}
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Title, &task.Body, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return task, nil
}

func (r *TaskRepositoryPSQL) UpdateTask(id string, updatedTask *Task) error {
	query := `UPDATE tasks SET title = $1, body = $2, completed = $3, updated_at = $4 WHERE id = $5`
	result, err := r.db.Exec(query, updatedTask.Title, updatedTask.Body, updatedTask.Completed, time.Now(), id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *TaskRepositoryPSQL) DeleteTask(id string) error {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *TaskRepositoryPSQL) GetAllTasks() []*Task {
	query := `SELECT id, title, body, completed, created_at, updated_at FROM tasks ORDER BY created_at DESC`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	tasks := []*Task{}
	for rows.Next() {
		task := &Task{}
		err := rows.Scan(&task.ID, &task.Title, &task.Body, &task.Completed, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			continue
		}
		tasks = append(tasks, task)
	}
	return tasks
}
