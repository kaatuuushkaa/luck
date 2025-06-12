package userService

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
