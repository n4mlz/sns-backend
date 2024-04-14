package usecases

import "time"

type UserNameDto struct {
	UserName string `json:"userName"`
}

type UserDisplayDto struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
}

type ProfileDto struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
}

type UserDto struct {
	UserName        string    `json:"userName"`
	DisplayName     string    `json:"displayName"`
	Biography       string    `json:"biography"`
	CreatedAt       time.Time `json:"createdAt"`
	FollowingStatus string    `json:"followingStatus"`
}

type PostIdDto struct {
	PostId string `json:"postId"`
}

type PostContentDto struct {
	Content string `json:"content"`
}

type PostDto struct {
	PostId    string         `json:"postId"`
	Poster    UserDisplayDto `json:"poster"`
	Content   string         `json:"content"`
	Likes     int            `json:"likes"`
	Liked     bool           `json:"liked"`
	Comments  int            `json:"comments"`
	CreatedAt time.Time      `json:"createdAt"`
}
