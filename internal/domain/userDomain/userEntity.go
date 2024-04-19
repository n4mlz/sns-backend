package userDomain

import (
	"errors"
	"os"
	"time"

	"github.com/n4mlz/sns-backend/internal/utils"
)

type User struct {
	userRepository      *IUserRepository
	userImageRepository *IUserImageRepository
	UserId              UserId
	UserName            UserName
	DisplayName         DisplayName
	Biography           Biography
	CreatedAt           time.Time
}

func (u *User) Follow(user *User) error {
	if u.IsFollowing(user) {
		return errors.New("already following")
	}

	(*u.userRepository).Follow(u, user)
	return nil
}

func (u *User) Unfollow(user *User) error {
	if u.UserId == user.UserId {
		return errors.New("cannot unfollow yourself")
	}

	if !u.IsFollowing(user) {
		return errors.New("not following")
	}

	(*u.userRepository).Unfollow(u, user)
	return nil
}

func (u *User) IsFollowing(user *User) bool {
	return (*u.userRepository).IsFollowing(u, user)
}

func (u *User) IsVisible(user *User) bool {
	followingStatus := u.GetFollowingStatus(user)
	return followingStatus == OWN || followingStatus == MUTUAL
}

func (u *User) GetFollowingStatus(user *User) string {
	if u.UserId == user.UserId {
		return OWN
	}

	isFollowing := (*u.userRepository).IsFollowing(u, user)

	isFollowed := (*u.userRepository).IsFollowing(user, u)

	if isFollowing && isFollowed {
		return MUTUAL
	} else if isFollowing {
		return FOLLOWING
	} else if isFollowed {
		return FOLLOWED
	} else {
		return NONE
	}
}

func (u *User) Followings() ([]*User, error) {
	return (*u.userRepository).FollowingUserList(u)
}

func (u *User) Followers() ([]*User, error) {
	return (*u.userRepository).FollowerUserList(u)
}

func (u *User) VisibleUsers() ([]*User, error) {
	return (*u.userRepository).VisibleUserList(u)
}

// following and not followed
func (u *User) FollowRequests() ([]*User, error) {
	return (*u.userRepository).FollowRequestUserList(u)
}

func (u *User) UserIconImageUrl() string {
	return utils.JoinPath(IMAGE_REPOSITORY_URL, "images", "user", u.UserName.String(), "icon.png")
}

func (u *User) UserBgImageUrl() string {
	return utils.JoinPath(IMAGE_REPOSITORY_URL, "images", "user", u.UserName.String(), "background.png")
}

func (u *User) UpdateIconImage(file *os.File) error {
	objectKey := u.UserIconImageUrl()
	return (*u.userImageRepository).SaveImage(objectKey, file)
}

func (u *User) UpdateBgImage(file *os.File) error {
	objectKey := u.UserBgImageUrl()
	return (*u.userImageRepository).SaveImage(objectKey, file)
}

func (u *User) DeleteIconImage() error {
	objectKey := u.UserIconImageUrl()
	return (*u.userImageRepository).Delete(objectKey)
}

func (u *User) DeleteBgImage() error {
	objectKey := u.UserBgImageUrl()
	return (*u.userImageRepository).Delete(objectKey)
}
