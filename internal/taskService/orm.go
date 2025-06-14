package taskService

type Task struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Task   string `json:"task"`
	IsDone bool   `json:"isDone"`
	UserID int    `json:"userId"`
}
type TaskRequest struct {
	Task string `json:"task"`
}
