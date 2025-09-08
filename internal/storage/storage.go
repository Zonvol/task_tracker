package storage

import "github.com/yourusername/task-tracker/internal/task"


// TaskStorage defines the interface for task storage operations.
type TaskStorage interface {
    Save(task *task.Task) error
    GetByID(id string) (*task.Task, error)
    GetAll() ([]*task.Task, error)
    Delete(id string) error
}

// InMemoryStorage is an in-memory implementation of TaskStorage.
type InMemoryStorage struct {
    tasks map[string]*task.Task
}

// NewInMemoryStorage creates a new instance of InMemoryStorage.
func NewInMemoryStorage() *InMemoryStorage {
    return &InMemoryStorage{
        tasks: make(map[string]*task.Task),
    }
}

// Save saves a task to the in-memory storage.
func (s *InMemoryStorage) Save(task *task.Task) error {
    s.tasks[task.ID] = task
    return nil
}

// GetByID retrieves a task by its ID from the in-memory storage.
func (s *InMemoryStorage) GetByID(id string) (*task.Task, error) {
    task, exists := s.tasks[id]
    if !exists {
        return nil, nil // or return an error
    }
    return task, nil
}

// GetAll retrieves all tasks from the in-memory storage.
func (s *InMemoryStorage) GetAll() ([]*task.Task, error) {
    var allTasks []*task.Task
    for _, task := range s.tasks {
        allTasks = append(allTasks, task)
    }
    return allTasks, nil
}

// Delete removes a task by its ID from the in-memory storage.
func (s *InMemoryStorage) Delete(id string) error {
    delete(s.tasks, id)
    return nil
}