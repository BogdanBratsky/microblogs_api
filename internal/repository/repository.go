package repository

import "github.com/BogdanBratsky/microblogs-api/internal/model"

type UserRepository interface {
	GetAllUsers() ([]model.UserDTO, error)
	GetUserById(id int) (model.UserDTO, error)
	CreateUser(name, email, password string) error
	// Для проверки:
	ExistsByUsername(email string) (bool, error)
	ExistsByEmail(email string) (bool, error)
}
