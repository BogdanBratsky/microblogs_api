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
	log.Println("Запускаем сервер...")

	r := chi.NewRouter()

	// !!!MECEВО!!!
	userRepo := memory.NewUserMemoryRepo()
	postRepo := memory.NewPostMemoryRepo()

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)
	// !!!MECEВО!!!

	// routes
	r.Route("/api/v1", func(r chi.Router) {
		// users
		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.GetUsersHandler)
			r.Get("/{id}", userHandler.GetUserByIdHandler)
			r.Post("/", userHandler.CreateUserHandler)
			r.Patch("/{id}", userHandler.UpdateUserHandler)
			r.Delete("/{id}", userHandler.DeleteUserHandler)
		})
		// posts
		r.Route("/posts", func(r chi.Router) {
			r.Get("/", postHandler.GetPostsHandler)
		})
		// auth
		// r.Route("/auth", func(r chi.Router) {
		// 	r.Post("/login", authHandler.)
		// 	r.Post("/register", authHandler.)
		// 	r.Post("/refresh", authHandler.)
		// })
	})

	// server
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal("Ошибка:", err.Error())
	}
}
