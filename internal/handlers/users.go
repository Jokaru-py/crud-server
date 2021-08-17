package handlers

import (
	"test_task/internal/pkg/repository"
	"test_task/models"
	"test_task/restapi/operations/users"

	"github.com/go-openapi/runtime/middleware"
)

// Получить всех пользователей
func GetAllUsers(repo repository.Users) middleware.Responder {
	res, err := repo.GetAllUsers()
	if err != nil {
		message := err.Error()
		return users.NewGetUsersInternalServerError().WithPayload(&models.Error{Message: &message})
	}

	return users.NewGetUsersOK().WithPayload(res)
}

// Добавить нового пользователя
func AddNewUser(repo repository.Users, params users.PostUsersParams) middleware.Responder {
	res, err := repo.AddNewUser(params)
	if err != nil {
		message := err.Error()
		return users.NewPostUsersInternalServerError().WithPayload(&models.Error{Message: &message})
	}

	return users.NewPostUsersOK().WithPayload(res)
}

// Удалить пользователя
func DeleteUser(repo repository.Users, params users.DeleteUsersParams) middleware.Responder {
	err := repo.DeleteUser(params)
	if err != nil {
		message := err.Error()
		return users.NewDeleteUsersInternalServerError().WithPayload(&models.Error{Message: &message})
	}

	return users.NewDeleteUsersOK()
}

// Обновить данные пользователя
func UpdateUser(repo repository.Users, params users.PatchUsersParams) middleware.Responder {
	res, err := repo.UpdateUser(params)
	if err != nil {
		message := err.Error()
		return users.NewPatchUsersInternalServerError().WithPayload(&models.Error{Message: &message})
	}

	return users.NewPatchUsersOK().WithPayload(res)
}
