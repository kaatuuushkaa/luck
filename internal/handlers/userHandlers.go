package handlers

import (
	"context"
	"task4/internal/userService"
	"task4/internal/web/users"
)

type UserHandler struct {
	service userService.UserService
}

func (u *UserHandler) GetUsers(ctx context.Context, request users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := u.service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := users.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		user := users.User{
			Id:       func(u uint) *uint { return &u }(uint(usr.ID)),
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (u *UserHandler) PostUsers(ctx context.Context, request users.PostUsersRequestObject) (users.PostUsersResponseObject, error) {
	userRequest := request.Body
	userToCreate := userService.User{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	err := u.service.CreateUser(&userToCreate)

	if err != nil {
		return nil, err
	}

	response := users.PostUsers201JSONResponse{
		Id:       func(u uint) *uint { return &u }(uint(userToCreate.ID)),
		Email:    &userToCreate.Email,
		Password: &userToCreate.Password,
	}

	return response, nil
}

func (u *UserHandler) DeleteUser(ctx context.Context, request users.DeleteUserRequestObject) (users.DeleteUserResponseObject, error) {
	err := u.service.DeleteUser(request.Id)
	if err != nil {
		errMsg := "Could not delete user"
		return users.DeleteUser500JSONResponse{Error: &errMsg}, nil
	}

	return users.DeleteUser204Response{}, nil
}

func (u *UserHandler) PatchUser(ctx context.Context, request users.PatchUserRequestObject) (users.PatchUserResponseObject, error) {
	if request.Body == nil {
		errMsg := "Request body is required"
		return users.PatchUser400JSONResponse{Error: &errMsg}, nil
	}
	userRequest := request.Body

	user, err := u.service.UpdateUser(request.Id, userRequest.Email, userRequest.Password)
	if err != nil {
		errMsg := "Could not update user"
		return users.PatchUser400JSONResponse{Error: &errMsg}, nil
	}
	response := users.PatchUser200JSONResponse{
		Id:       func(u uint) *uint { return &u }(uint(user.ID)),
		Email:    &user.Email,
		Password: &user.Password,
	}
	return response, nil
}

func NewUserHandler(u userService.UserService) *UserHandler {
	return &UserHandler{service: u}
}
