package service

import (
	"errors"
	"log"

	"github.com/BogdanBratsky/microblogs-api/internal/model"
	"github.com/BogdanBratsky/microblogs-api/internal/repository"
)

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) GetAllUsersService() ([]model.UserDTO, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByIdService(id int) (model.UserDTO, error) {
	return s.repo.GetUserById(id)
}

func (s *userService) CreateUserService(name, email, password string) error {
	// Проверяем, занята ли почта
	ok, err := s.repo.ExistsByEmail(email)
	if err != nil {
		log.Println("Ошибка:", err.Error())
		return errors.New("ошибка при создании пользователя")
	}
	if ok {
		return errors.New("пользователь с такой почтой уже существует")
	}
	// Проверяем, занят ли логин
	ok, err = s.repo.ExistsByUsername(name)
	if err != nil {
		log.Println("Ошибка:", err.Error())
		return errors.New("ошибка при создании пользователя")
	}
	if ok {
		return errors.New("пользователь с таким логином уже существует")
	}
	// Создание пользователя на уровне доступа к данным
	return s.repo.CreateUser(name, email, password)
}
