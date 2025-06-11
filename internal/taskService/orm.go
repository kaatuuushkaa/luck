package taskService

type Task struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
}
type TaskRequest struct {
	Task string `json:"task"`
}
