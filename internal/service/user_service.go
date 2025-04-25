package service

import (
	"errors"
	"fmt"
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
	exists, err := s.repo.ExistsById(id)
	if err != nil {
		log.Println("Ошибка:", err.Error())
		return model.UserDTO{}, errors.New("ошибка при поиске пользователя")
	}
	if !exists {
		return model.UserDTO{}, fmt.Errorf("пользователь с id=%d не найден", id)
	}
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

func (s *userService) UpdateUserService(id int, user model.UserUpdate) error {
	log.Printf("Данные для обновления %v\n", user)
	return s.repo.UpdateUser(id, user)
}

func (s *userService) DeleteUserService(id int) error {
	exists, err := s.repo.ExistsById(id)
	if err != nil {
		log.Println(err.Error())
	}
	if !exists {
		return fmt.Errorf("пользователя с id=%d не существует", id)
	}
	return s.repo.DeleteUser(id)
}
