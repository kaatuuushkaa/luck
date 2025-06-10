package taskService

type TaskService interface {
	CreateTask(task Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	UpdateTask(task Task) (Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskService{repo: r}
}

func (ts *taskService) CreateTask(task Task) error {
	return ts.repo.CreateTask(task)
}

func (ts *taskService) GetAllTasks() ([]Task, error) {
	return ts.repo.GetAllTasks()
}

func (ts *taskService) GetTaskByID(id int) (Task, error) {
	return ts.repo.GetTaskByID(id)
}

func (ts *taskService) UpdateTask(task Task) (Task, error) {
	existing, err := ts.repo.GetTaskByID(task.ID)
	if err != nil {
		return Task{}, err
	}

	existing.Task = task.Task
	existing.IsDone = task.IsDone

	if err := ts.repo.UpdateTask(existing); err != nil {
		return Task{}, err
	}

	return existing, nil
}

func (ts *taskService) DeleteTask(id int) error {
	return ts.repo.DeleteTask(id)
}
