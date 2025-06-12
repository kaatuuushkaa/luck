package userService

// work update
type UserService interface {
	CreateUser(user *User) error
	GetAllUsers() ([]User, error)
	GetUserByID(id string) (User, error)
	UpdateUser(id, email, password string) (User, error)
	DeleteUser(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(u UserRepository) UserService {
	return &userService{repo: u}
}

func (us *userService) CreateUser(user *User) error {
	return us.repo.CreateUser(user)
}

func (us *userService) GetAllUsers() ([]User, error) {
	return us.repo.GetAllUsers()
}

func (us *userService) GetUserByID(id string) (User, error) {
	return us.repo.GetUserByID(id)
}

func (us *userService) UpdateUser(id, email, password string) (User, error) {
	existing, err := us.repo.GetUserByID(id)
	if err != nil {
		return User{}, err
	}

	existing.Email = email
	existing.Password = password

	if err := us.repo.UpdateUser(existing); err != nil {
		return User{}, err
	}

	return existing, nil
}

func (us *userService) DeleteUser(id string) error {
	return us.repo.DeleteUser(id)
}
