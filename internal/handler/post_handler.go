package handler

import (
	"log"
	"net/http"

	"github.com/BogdanBratsky/microblogs-api/internal/service"
	"github.com/BogdanBratsky/microblogs-api/pkg"
)

type PostHandlerImpl struct {
	service service.PostService
}

func NewPostHandler(service service.PostService) *PostHandlerImpl {
	return &PostHandlerImpl{service: service}
}

func (h *PostHandlerImpl) GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request recieved:", r.Method, r.URL)

	posts, err := h.service.GetAllPostsService()
	if err != nil {
		log.Println("Ошибка:", err.Error())
		pkg.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	pkg.WriteSuccess(w, http.StatusOK, posts)
}
