package memory

import (
	"fmt"
	"sort"
	"time"

	"github.com/BogdanBratsky/microblogs-api/internal/model"
)

type PostMemoryRepo struct {
	data map[int]*model.Post
}

func NewPostMemoryRepo() *PostMemoryRepo {
	return &PostMemoryRepo{
		data: func(count int) map[int]*model.Post {
			var posts = make(map[int]*model.Post)
			for i := 0; i < count; i++ {
				posts[i] = &model.Post{
					Id:           i,
					Content:      fmt.Sprintf("Text for id %d", i),
					ParentPostID: nil,
					UserId:       1,
					CreatedAt:    time.Now(),
					UpdatedAt:    time.Now(),
				}
			}
			return posts
		}(20),
	}
}

func (p *PostMemoryRepo) GetAllPosts() ([]model.Post, error) {
	var posts []model.Post
	for _, post := range p.data {
		posts = append(posts, *post)
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Id > posts[j].Id
	})
	return posts, nil
}
