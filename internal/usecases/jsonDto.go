package usecases

import "time"

type UserNameDto struct {
	UserName string `json:"userName"`
}

type UserDisplayDto struct {
	UserName    string `json:"userName"`
	DisplayName string `json:"displayName"`
}

type UserSettingsDto struct {
	DisplayName string `json:"displayName"`
	Biography   string `json:"biography"`
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

type PostDetailDto struct {
	PostId    string         `json:"postId"`
	Poster    UserDisplayDto `json:"poster"`
	Content   string         `json:"content"`
	Likes     int            `json:"likes"`
	Liked     bool           `json:"liked"`
	Comments  []CommentDto   `json:"comments"`
	CreatedAt time.Time      `json:"createdAt"`
}

type CreateCommentDto struct {
	PostId  string `json:"postId"`
	Content string `json:"content"`
}

type CommentDto struct {
	CommentId string         `json:"commentId"`
	Commenter UserDisplayDto `json:"commenter"`
	Content   string         `json:"content"`
	Replies   []ReplyDto     `json:"replies"`
	CreatedAt time.Time      `json:"createdAt"`
}

type CreateReplyDto struct {
	CommentId string `json:"commentId"`
	Content   string `json:"content"`
}

type ReplyDto struct {
	ReplyId   string         `json:"replyId"`
	Replier   UserDisplayDto `json:"replier"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"createdAt"`
}
