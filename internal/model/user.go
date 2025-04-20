package model

import "time"

type User struct {
	Id           int       `json:"id"`
	Username     string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password"`
	AvatarURL    string    `json:"avatarURL"`
	Status       string    `json:"status"`
	About        string    `json:"about"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (u User) ToUserDTO() UserDTO {
	return UserDTO{
		Id:        u.Id,
		Username:  u.Username,
		AvatarURL: u.AvatarURL,
		Status:    u.Status,
		About:     u.About,
	}
}

type UserDTO struct {
	Id        int    `json:"id"`
	Username  string `json:"name"`
	AvatarURL string `json:"avatarURL"`
	Status    string `json:"status"`
	About     string `json:"about"`
}
