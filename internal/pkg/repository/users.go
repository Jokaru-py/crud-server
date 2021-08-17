package repository

import (
	"test_task/models"
	"test_task/restapi/operations/users"
)

type Users interface {
	GetAllUsers() (*models.AllUserResponse, error)
	AddNewUser(users.PostUsersParams) (*models.UserUploadResponse, error)
	DeleteUser(users.DeleteUsersParams) error
	UpdateUser(users.PatchUsersParams) (*models.User, error)
}
