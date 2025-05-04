package main

import (
	"log"

	"github.com/BogdanBratsky/microblogs-api/internal/app"
)

func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		log.Fatal("Ошибка запуска:", err)
	}
}
