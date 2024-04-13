package usecases

import "time"

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

type UserNameDto struct {
	UserName string `json:"userName"`
}
