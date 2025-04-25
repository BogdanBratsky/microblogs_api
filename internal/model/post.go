package model

import "time"

type Post struct {
	Id           int       `json:"id"`
	Content      string    `json:"content"`
	ParentPostID *int      `json:"parentPostId,omitempty"`
	UserId       int       `json:"userId"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
