package service

import "github.com/BogdanBratsky/microblogs-api/internal/model"

type UserService interface {
	GetAllUsersService() ([]model.UserDTO, error)
	GetUserByIdService(id int) (model.UserDTO, error)
	CreateUserService(name, email, password string) error
}
