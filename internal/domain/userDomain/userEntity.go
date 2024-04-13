package userDomain

import (
	"errors"
	"time"
)

type User struct {
	UserRepository IUserRepository
	UserId         UserId
	UserName       UserName
	DisplayName    DisplayName
	Biography      Biography
	CreatedAt      time.Time
}

func (u *User) SaveUser() error {
	if !u.UserName.IsValid() || !u.DisplayName.IsValid() || !u.Biography.IsValid() {
		return errors.New("invalid profile")
	}
	u.UserRepository.Save(u)
	return nil
}

func (u *User) Follow(user *User) error {
	if u.IsFollowing(user) {
		return errors.New("already following")
	}

	u.UserRepository.Follow(u, user)
	return nil
}

func (u *User) Unfollow(user *User) error {
	if u.UserId == user.UserId {
		return errors.New("cannot unfollow yourself")
	}

	if !u.IsFollowing(user) {
		return errors.New("not following")
	}

	u.UserRepository.Unfollow(u, user)
	return nil
}

func (u *User) IsFollowing(user *User) bool {
	return u.UserRepository.IsFollowing(u, user)
}

func (u *User) IsMutualFollow(user *User) bool {
	followingStatus := u.GetFollowingStatus(user)
	return followingStatus == MUTUAL
}

func (u *User) GetFollowingStatus(user *User) string {
	if u.UserId == user.UserId {
		return OWN
	}

	isFollowing := u.UserRepository.IsFollowing(u, user)

	isFollowed := u.UserRepository.IsFollowing(user, u)

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

func (u *User) FollowingUserList() ([]*User, error) {
	return u.UserRepository.FollowingUserList(u)
}

func (u *User) FollowerUserList() ([]*User, error) {
	return u.UserRepository.FollowerUserList(u)
}

func (u *User) MutualFollowUserList() ([]*User, error) {
	return u.UserRepository.MutualFollowUserList(u)
}

// following and not followed
func (u *User) FollowRequestUserList() ([]*User, error) {
	return u.UserRepository.FollowRequestUserList(u)
}
