package models

type ProfileSchema struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
}

type UserNameSchema struct {
	UserName string `json:"userName"`
}
