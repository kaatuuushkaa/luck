package handlers

import (
	"context"
	"task4/internal/taskService"
	"task4/internal/web/tasks"
)

type TaskHandler struct {
	service taskService.TaskService
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     func(u uint) *uint { return &u }(uint(tsk.ID)),
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	err := h.service.CreateTask(&taskToCreate)

	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     func(u uint) *uint { return &u }(uint(taskToCreate.ID)),
		Task:   &taskToCreate.Task,
		IsDone: &taskToCreate.IsDone,
	}

	return response, nil
}

func NewTaskHandler(s taskService.TaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) PatchTask(ctx context.Context, request tasks.PatchTaskRequestObject) (tasks.PatchTaskResponseObject, error) {
	if request.Body == nil {
		errMsg := "Request body is required"
		return tasks.PatchTask400JSONResponse{Error: &errMsg}, nil
	}
	taskRequest := request.Body

	task, err := h.service.UpdateTask(request.Id, taskRequest.Task)
	if err != nil {
		errMsg := "Could not update task"
		return tasks.PatchTask400JSONResponse{Error: &errMsg}, err
	}
	return tasks.PatchTask200JSONResponse{Task: task.Task}, nil
}

func (h *TaskHandler) DeleteTask(ctx context.Context, request tasks.DeleteTaskRequestObject) (tasks.DeleteTaskResponseObject, error) {
	err := h.service.DeleteTask(request.Id)
	if err != nil {
		errMsg := "Could not delete task"
		return tasks.DeleteTask500JSONResponse{Error: &errMsg}, nil
	}

	return tasks.DeleteTask204Response{}, nil
}
