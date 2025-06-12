//GetUsers (выводит всех пользователей),
//PostUser (создать нового пользователя)
//PatchUserByID (Отредактировать поля user по его ID) и
//DeleteUserByID

package userService

import (
	"gorm.io/gorm"
)

// work update
type UserRepository interface {
	CreateUser(user *User) error
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(user User) error
	DeleteUser(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user *User) error {
	return u.db.Create(&user).Error
}

func (u *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userRepository) GetUserByID(id string) (User, error) {
	var user User
	err := u.db.First(&user, "id = ?", id).Error
	return user, err
}

func (u *userRepository) UpdateUser(user User) error {
	return u.db.Save(&user).Error
}

func (u *userRepository) DeleteUser(id string) error {
	return u.db.Delete(&User{}, "id = ?", id).Error
}
