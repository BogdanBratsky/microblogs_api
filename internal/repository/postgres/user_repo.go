package postgres

import (
	"database/sql"

	"github.com/BogdanBratsky/microblogs-api/internal/model"
)

type UserPostgresRepo struct {
	db *sql.DB
}

func NewUserPostgresRepo(db *sql.DB) *UserPostgresRepo {
	return &UserPostgresRepo{db: db}
}

func (u *UserPostgresRepo) GetAllUsers() ([]model.UserDTO, error) {
	return []model.UserDTO{}, nil
}

func (u *UserPostgresRepo) GetUserById(id int) (model.UserDTO, error) {
	return model.UserDTO{}, nil
}

func (u *UserPostgresRepo) CreateUser(name, email, password string) error {
	return nil
}

func (u *UserPostgresRepo) UpdateUser(id int, input model.UserUpdate) error {
	return nil
}

func (u *UserPostgresRepo) DeleteUser(id int) error {
	return nil
}

// Для проверки:
func (u *UserPostgresRepo) ExistsByUsername(email string) (bool, error) {
	return true, nil
}

func (u *UserPostgresRepo) ExistsByEmail(email string) (bool, error) {
	return true, nil
}

func (u *UserPostgresRepo) ExistsById(id int) (bool, error) {
	return true, nil
}
