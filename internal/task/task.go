package task

type Task struct {
    ID          int
    Title       string
    Description string
    Status      string
}

func NewTask(id int, title, description, status string) *Task {
    return &Task{
        ID:          id,
        Title:       title,
        Description: description,
        Status:      status,
    }
}

func (t *Task) UpdateTitle(title string) {
    t.Title = title
}

func (t *Task) UpdateDescription(description string) {
    t.Description = description
}

func (t *Task) UpdateStatus(status string) {
    t.Status = status
}