package main

import (
	"log"
	"net/http"

	"github.com/BogdanBratsky/microblogs-api/internal/handler"
	"github.com/BogdanBratsky/microblogs-api/internal/repository/memory"
	"github.com/BogdanBratsky/microblogs-api/internal/service"
	"github.com/go-chi/chi"
)

func main() {
	log.Println("Start server...")

	r := chi.NewRouter()

	// Месево...

	// Создаём хранилище юзеров в карте
	repo := memory.NewUserMemoryRepo()
	// Создаём экземпляр бизнес-логики, работающей с картой юзеров (repo)
	userService := service.NewUserService(repo)
	// Создаём обработчик запросов для юзеров, подключаем к нему userService
	userHandler := handler.NewUserHandler(userService)

	// routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.GetUsersHandler)
			r.Get("/{id}", userHandler.GetUserByIdHandler)
			r.Post("/", userHandler.CreateUserHandler)
		})
	})

	// server
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Ошибка:", err.Error())
	}
}
