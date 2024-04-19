package userDomain

import "io"

type IUserRepository interface {
	Save(user *User) error
	FindById(id UserId) (*User, error)
	FindByIds(ids []UserId) ([]*User, error)
	FindByUserName(userName UserName) (*User, error)
	IsExistUserId(userId UserId) bool
	IsExistUserName(userName UserName) bool
	Follow(sourceUser *User, targetUser *User) error
	Unfollow(sourceUser *User, targetUser *User) error
	IsFollowing(sourceUser *User, targetUser *User) bool
	FollowingUserList(user *User) ([]*User, error)
	FollowerUserList(user *User) ([]*User, error)
	VisibleUserList(user *User) ([]*User, error)
	FollowRequestUserList(user *User) ([]*User, error)
}

type IUserImageRepository interface {
	SaveIcon(objectKey string, file io.Reader) error
	SaveBgImage(objectKey string, file io.Reader) error
	SaveBinary(objectKey string, fileBytes []byte) error
	Delete(objectKey string) error
	Move(sourceObjectKey string, targetObjectKey string) error
}
