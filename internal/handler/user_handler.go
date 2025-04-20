package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/BogdanBratsky/microblogs-api/internal/model"
	"github.com/BogdanBratsky/microblogs-api/internal/service"
	"github.com/BogdanBratsky/microblogs-api/pkg"
	"github.com/go-chi/chi"
)

// Структура обработчика пользователей
type UserHandlerImpl struct {
	service service.UserService
}

// Функция для создания экземпляра обработчика пользователей
func NewUserHandler(service service.UserService) *UserHandlerImpl {
	return &UserHandlerImpl{service: service}
}

// Обработчик для получения всех пользователей
func (h *UserHandlerImpl) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved:", r.Method, r.URL)

	users, _ := h.service.GetAllUsersService()

	pkg.WriteSuccess(w, http.StatusOK, users)
}

// Обработчик для получения конкретного пользователя по его id
func (h *UserHandlerImpl) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved:", r.Method, r.URL)

	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Println("Ошибка:", err.Error())
		return
	}

	user, err := h.service.GetUserByIdService(id)

	if err != nil {
		pkg.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	pkg.WriteSuccess(w, http.StatusOK, user)
}

// Обработчик для создания пользователя
func (h *UserHandlerImpl) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved:", r.Method, r.URL)

	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Ошибка:", err.Error())
		return
	}

	err := h.service.CreateUserService(user.Username, user.Email, user.PasswordHash)

	if err != nil {
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	pkg.WriteSuccess(w, http.StatusCreated, map[string]string{
		"message": "пользователь успешно создан",
	})
}
