package userService

import "task4/internal/taskService"

type User struct {
	ID       int                `gorm:"primaryKey" json:"id"`
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Tasks    []taskService.Task `json:"tasks"`
}
