package restapi

import (
	"errors"
	"test_task/internal/pkg/database"
	"test_task/internal/pkg/repository"
	"test_task/internal/repositories/users"
)

// Repositories описывает список репозиториев
type Repositories struct {
	Users repository.Users
}

func configureRepositories(dbName string) (Repositories, error) {
	// Подключение к БД.
	db, err := database.Open(dbName)
	if err != nil {
		return Repositories{}, errors.New("Ошибка получения конфигурации репозиториев: " + dbName)
	}

	return Repositories{
		Users: users.New(db),
	}, nil
}
