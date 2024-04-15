package userDomain

import (
	"errors"
	"time"
)

type User struct {
	UserRepository *IUserRepository
	UserId         UserId
	UserName       UserName
	DisplayName    DisplayName
	Biography      Biography
	CreatedAt      time.Time
}

func (u *User) Follow(user *User) error {
	if u.IsFollowing(user) {
		return errors.New("already following")
	}

	(*u.UserRepository).Follow(u, user)
	return nil
}

func (u *User) Unfollow(user *User) error {
	if u.UserId == user.UserId {
		return errors.New("cannot unfollow yourself")
	}

	if !u.IsFollowing(user) {
		return errors.New("not following")
	}

	(*u.UserRepository).Unfollow(u, user)
	return nil
}

func (u *User) IsFollowing(user *User) bool {
	return (*u.UserRepository).IsFollowing(u, user)
}

func (u *User) IsVisible(user *User) bool {
	followingStatus := u.GetFollowingStatus(user)
	return followingStatus == OWN || followingStatus == MUTUAL
}

func (u *User) GetFollowingStatus(user *User) string {
	if u.UserId == user.UserId {
		return OWN
	}

	isFollowing := (*u.UserRepository).IsFollowing(u, user)

	isFollowed := (*u.UserRepository).IsFollowing(user, u)

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
	return (*u.UserRepository).FollowingUserList(u)
}

func (u *User) Followers() ([]*User, error) {
	return (*u.UserRepository).FollowerUserList(u)
}

func (u *User) MutualFollows() ([]*User, error) {
	return (*u.UserRepository).MutualFollowUserList(u)
}

// following and not followed
func (u *User) FollowRequests() ([]*User, error) {
	return (*u.UserRepository).FollowRequestUserList(u)
}
