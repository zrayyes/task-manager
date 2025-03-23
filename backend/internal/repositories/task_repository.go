package repositories

import (
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

	if _, exists := r.tasks[task.ID]; exists {
		return ErrAlreadyExists
	}

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
