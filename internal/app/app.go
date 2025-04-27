package app

import (
	"log"
	"net/http"

	"github.com/BogdanBratsky/microblogs-api/internal/handler"
	"github.com/BogdanBratsky/microblogs-api/internal/repository/memory"
	"github.com/BogdanBratsky/microblogs-api/internal/service"
	"github.com/go-chi/chi"
)

type App struct {
	router http.Handler
}

func NewApp() *App {
	userRepo := memory.NewUserMemoryRepo()
	postRepo := memory.NewPostMemoryRepo()

	userService := service.NewUserService(userRepo)
	postService := service.NewPostService(postRepo)

	userHandler := handler.NewUserHandler(userService)
	postHandler := handler.NewPostHandler(postService)

	// Router
	r := chi.NewRouter()

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

	return &App{router: r}
}

func (a *App) Run(addr string) error {
	log.Println("Запускаем сервер на порту: ", addr)
	return http.ListenAndServe(addr, a.router)
}
