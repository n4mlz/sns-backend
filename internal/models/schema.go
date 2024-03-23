package models

import "time"

type UserNameSchema struct {
	UserName string `json:"userName"`
}

type UserDisplaySchema struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
}

type ProfileSchema struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
}

type UserDataSchema struct {
	UserName        string    `json:"userName"`
	DisplayName     string    `json:"displayName"`
	Biography       string    `json:"biography"`
	CreatedAt       time.Time `json:"createdAt"`
	FollowingStatus string    `json:"followingStatus"`
}

type PostContentSchema struct {
	Content string `json:"content"`
}

type PostSchema struct {
	PostId    string            `json:"postId"`
	Poster    UserDisplaySchema `json:"poster"`
	Content   string            `json:"content"`
	Likes     int               `json:"likes"`
	Liked     bool              `json:"liked"`
	Comments  int               `json:"comments"`
	CreatedAt time.Time         `json:"createdAt"`
}
