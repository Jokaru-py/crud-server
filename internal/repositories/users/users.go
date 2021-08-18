package users

import (
	"database/sql"
	"errors"
	"test_task/internal/pkg/validation"
	"test_task/models"
	"test_task/restapi/operations/users"

	"github.com/kisielk/sqlstruct"
)

type Repository struct {
	db *sql.DB
}

// New создает новый экземпляр репозитория БД
func New(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// Возвращает всех пользователей
func (repo *Repository) GetAllUsers() (*models.AllUserResponse, error) {
	rows, err := repo.db.Query(`select * from public.t_users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*models.User
	var user models.User
	for rows.Next() {
		// Делает нормальный скан на структуру.
		err := sqlstruct.Scan(&user, rows)
		if err != nil {
			return nil, err
		}

		// Для передачи значения, а не указателя.
		user1 := user
		users = append(users, &user1)
	}

	return &models.AllUserResponse{Users: users}, nil
}

// Добавление нового пользователя в БД
func (repo *Repository) AddNewUser(params users.PostUsersParams) (*models.UserUploadResponse, error) {
	ok := validation.Email(params.LoadUser.Email, false)
	if !ok {
		return nil, errors.New("Не корректный email")
	}

	var id string
	// Вызов функции в БД с передачей данных о пользователе.
	err := repo.db.QueryRow(`select * from public.fn_user_ins($1,$2,$3)`, params.LoadUser.Name, params.LoadUser.Age, params.LoadUser.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &models.UserUploadResponse{ID: &id}, nil
}

// Удаление пользователя
func (repo *Repository) DeleteUser(params users.DeleteUsersParams) error {
	_, err := repo.db.Exec(`select * from public.fn_user_del($1)`, params.IDUSer.ID)
	if err != nil {
		return err
	}

	return nil
}

// Обновить данные пользователя
func (repo *Repository) UpdateUser(params users.PatchUsersParams) (*models.User, error) {
	ok := validation.Email(params.IDUSer.Email, false)
	if !ok {
		return nil, errors.New("Не корректный email")
	}

	rows, err := repo.db.Query(`select * from public.fn_user_upd($1,$2,$3,$4)`, params.IDUSer.Name, params.IDUSer.Age, params.IDUSer.Email, params.IDUSer.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		// sqlstruct для корректной обработки за структуру
		err := sqlstruct.Scan(&user, rows)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}
