package repository

import "github.com/BogdanBratsky/microblogs-api/internal/model"

type UserRepository interface {
	GetAllUsers() ([]model.UserDTO, error)
	GetUserById(id int) (model.UserDTO, error)
	CreateUser(name, email, password string) error
	UpdateUser(id int, input model.UserUpdate) error
	DeleteUser(id int) error
	// Для проверки:
	ExistsByUsername(email string) (bool, error)
	ExistsByEmail(email string) (bool, error)
	ExistsById(id int) (bool, error)
}

type PostRepository interface {
	GetAllPosts() ([]model.Post, error)
	// GetPostById(id int) (model.Post, error)
	// CreatePost(content string, userId int) error
	// UpdatePost(content string) error
	// DeletePost(id, userId int) error
}
