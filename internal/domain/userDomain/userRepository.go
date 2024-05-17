package userDomain

import "io"

type IUserRepository interface {
	Save(user *User) error
	FindById(id UserId) (*User, error)
	FindByIds(ids []UserId) ([]*User, error)
	FindByUserName(userName UserName) (*User, error)
	IsExistUserId(userId UserId) bool
	IsExistUserName(userName UserName) bool
	GetVisibleUserCount(user *User) (int, error)
	Follow(sourceUser *User, targetUser *User) error
	Unfollow(sourceUser *User, targetUser *User) error
	IsFollowing(sourceUser *User, targetUser *User) bool
	FollowingUserList(user *User) ([]*User, error)
	FollowerUserList(user *User) ([]*User, error)
	VisibleUserList(user *User) ([]*User, error)
	FollowRequestUserList(user *User) ([]*User, error)
}

type IUserImageRepository interface {
	SaveIcon(user *User, file io.Reader) (ImageUrl, error)
	SaveBgImage(user *User, file io.Reader) (ImageUrl, error)
	DeleteIcon(user *User) error
	DeleteBgImage(user *User) error
	MoveResources(sourceUser *User, targetUser *User) (ImageUrl, ImageUrl, error)
}
