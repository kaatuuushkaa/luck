package taskService

type TaskService interface {
	CreateTask(task *Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id string) (Task, error)
	UpdateTask(id, task string) (Task, error)
	DeleteTask(id string) error
	GetTasksForUser(userID uint) ([]Task, error)
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (ts *taskService) CreateTask(task *Task) error {
	return ts.repo.CreateTask(task)
}

func (ts *taskService) GetAllTasks() ([]Task, error) {
	return ts.repo.GetAllTasks()
}

func (ts *taskService) GetTaskByID(id string) (Task, error) {
	return ts.repo.GetTaskByID(id)
}

func (ts *taskService) UpdateTask(id, task string) (Task, error) {
	existing, err := ts.repo.GetTaskByID(id)
	if err != nil {
		return Task{}, err
	}

	existing.Task = task

	if err := ts.repo.UpdateTask(existing); err != nil {
		return Task{}, err
	}

	return existing, nil
}

func (ts *taskService) DeleteTask(id string) error {
	return ts.repo.DeleteTask(id)
}
func (ts *taskService) GetTasksForUser(userID uint) ([]Task, error) {
	return ts.repo.GetTasksForUser(userID)
}
