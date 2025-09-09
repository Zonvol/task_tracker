package task

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func AddTask(id int, firstArg string, secondArg string, status bool) *Task {
	return &Task{
		ID:          id,
		Title:       firstArg,
		Description: secondArg,
		Status:      status,
	}
}
func DeleteTask() {

}
func UpdateTask() {

}
func ListTask() {

}
