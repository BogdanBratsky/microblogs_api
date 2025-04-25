package service

import (
	"github.com/BogdanBratsky/microblogs-api/internal/model"
	"github.com/BogdanBratsky/microblogs-api/internal/repository"
)

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) *postService {
	return &postService{repo: repo}
}

func (s *postService) GetAllPostsService() ([]model.Post, error) {
	return s.repo.GetAllPosts()
}
