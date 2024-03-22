package models

import "time"

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

type UserNameSchema struct {
	UserName string `json:"userName"`
}
