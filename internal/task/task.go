package task

import "time"

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Status      bool       `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"` // опциональное поле (может быть nil)
	// Можно реализовать использования дефолтного значения для тайм, в случае необходимости проверять на .IsZero, хз как будет правильнее
}

func AddTask(id int, firstArg string, status bool, createdAt time.Time, updatedAt *time.Time) *Task {
	return &Task{
		ID:          id,
		Title:       firstArg,
		Status:      status,
		CreatedAt:	createdAt, 
		UpdatedAt: updatedAt,
	}
}
