package memory

import (
	"fmt"
	"sort"
	"time"

	"github.com/BogdanBratsky/microblogs-api/internal/model"
	"github.com/BogdanBratsky/microblogs-api/internal/repository"
)

type UserMemoryRepo struct {
	data map[int]model.User
}

func NewUserMemoryRepo() repository.UserRepository {
	return &UserMemoryRepo{
		// Создание карты пользователей
		data: func(count int) map[int]model.User {
			var users = make(map[int]model.User)
			for i := 0; i < count; i++ {
				users[i] = model.User{
					Id:           i,
					Username:     fmt.Sprintf("user%d", i),
					Email:        fmt.Sprintf("user%d@email.com", i),
					PasswordHash: fmt.Sprintf("user%dpassword", i),
					AvatarURL:    "",
					Status:       "",
					About:        "",
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}
			}
			return users
		}(10),
	}
}

// Получение всех пользователей из карты
func (u *UserMemoryRepo) GetAllUsers() ([]model.UserDTO, error) {
	var users []model.UserDTO
	for _, user := range u.data {
		users = append(users, user.ToUserDTO())
	}
	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})
	return users, nil
}

// Получения пользователя по id из карты
func (u *UserMemoryRepo) GetUserById(id int) (model.UserDTO, error) {
	user := u.data[id]
	return user.ToUserDTO(), nil
}

// Добавление нового пользователя в карту
func (u *UserMemoryRepo) CreateUser(name, email, password string) error {
	dataLen := len(u.data)
	u.data[dataLen] = model.User{
		Id:           dataLen,
		Username:     name,
		Email:        email,
		PasswordHash: password,
		AvatarURL:    "",
		Status:       "",
		About:        "",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	return nil
}

// Удаление пользователя
// НАДО БУДЕТ ПЕРЕПИСАТЬ, НАВЕРНОЕ
func (u *UserMemoryRepo) DeleteUser(id int) error {
	if _, exists := u.data[id]; !exists {
		return fmt.Errorf("Пользователя с таким %d не существует", id)
	}
	delete(u.data, id)
	return nil
}

// Проверка существования пользователя с таким email
func (u *UserMemoryRepo) ExistsByUsername(name string) (bool, error) {
	for _, user := range u.data {
		if name == user.Username {
			return true, nil
		}
	}
	return false, nil
}

// Проверка существования пользователя с таким username
func (u *UserMemoryRepo) ExistsByEmail(email string) (bool, error) {
	for _, user := range u.data {
		if email == user.Email {
			return true, nil
		}
	}
	return false, nil
}
