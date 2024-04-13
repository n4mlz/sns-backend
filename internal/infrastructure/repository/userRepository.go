package repository

import (
	"context"
	"errors"

	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/model"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/query"
	"github.com/rs/xid"
)

type UserRepository struct{}

func toGormUser(user *userDomain.User) *model.User {
	return &model.User{
		ID:          user.UserId.String(),
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
		Biography:   user.Biography,
		CreatedAt:   user.CreatedAt,
	}
}

func toUser(gormUser *model.User) *userDomain.User {
	return &userDomain.User{
		UserId:      userDomain.UserId(gormUser.ID),
		UserName:    userDomain.UserName(gormUser.UserName),
		DisplayName: userDomain.DisplayName(gormUser.DisplayName),
		Biography:   gormUser.Biography,
		CreatedAt:   gormUser.CreatedAt,
	}
}

func (r *UserRepository) Save(user *userDomain.User) {
	gormUser := toGormUser(user)

	query.User.WithContext(context.Background()).Save(gormUser)

	if !user.IsFollowing(user) {
		newFollow := &model.Follow{
			ID:              xid.New().String(),
			FollowerUserID:  gormUser.ID,
			FollowingUserID: gormUser.ID,
		}

		query.Follow.WithContext(context.Background()).Save(newFollow)
	}
}

func (r *UserRepository) FindById(userId userDomain.UserId) (*userDomain.User, error) {
	if !r.IsExistUserId(userId) {
		return nil, errors.New("user not found")
	}

	gormUser, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(userId.String())).Take()
	return toUser(gormUser), err
}

func (r *UserRepository) FindByIds(ids []userDomain.UserId) ([]*userDomain.User, error) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, id.String())
	}

	gormUsers, err := query.User.WithContext(context.Background()).Where(query.User.ID.In(idsStr...)).Find()
	if err != nil {
		return nil, err
	}

	var users []*userDomain.User
	for _, gormUser := range gormUsers {
		users = append(users, toUser(gormUser))
	}

	return users, nil
}

func (r *UserRepository) FindByUserName(userName userDomain.UserName) (*userDomain.User, error) {
	if !r.IsExistUserName(userName) {
		return nil, errors.New("user not found")
	}

	gormUser, err := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq(userName.String())).Take()
	return toUser(gormUser), err
}

func (r *UserRepository) IsExistUserId(userId userDomain.UserId) bool {
	count, _ := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(userId.String())).Count()
	return count != 0
}

func (r *UserRepository) IsExistUserName(userName userDomain.UserName) bool {
	count, _ := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq(userName.String())).Count()
	return count != 0
}

func (r *UserRepository) Follow(sourceUser *userDomain.User, targetUser *userDomain.User) error {
	newFollow := &model.Follow{
		ID:              xid.New().String(),
		FollowerUserID:  sourceUser.UserId.String(),
		FollowingUserID: targetUser.UserId.String(),
	}

	return query.Follow.WithContext(context.Background()).Save(newFollow)
}

func (r *UserRepository) Unfollow(sourceUser *userDomain.User, targetUser *userDomain.User) error {
	_, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(sourceUser.UserId.String())).Where(query.Follow.FollowingUserID.Eq(targetUser.UserId.String())).Delete()
	return err
}

func (r *UserRepository) IsFollowing(sourceUser *userDomain.User, targetUser *userDomain.User) bool {
	count, _ := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(sourceUser.UserId.String())).Where(query.Follow.FollowingUserID.Eq(targetUser.UserId.String())).Count()
	return count != 0
}

func (r *UserRepository) FollowingUserList(user *userDomain.User) ([]*userDomain.User, error) {
	followings, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(user.UserId.String())).Find()
	if err != nil {
		return nil, err
	}

	var followingUserIdList []userDomain.UserId
	for _, following := range followings {
		followingUserIdList = append(followingUserIdList, userDomain.UserId(following.FollowingUserID))
	}

	return r.FindByIds(followingUserIdList)
}

func (r *UserRepository) FollowerUserList(user *userDomain.User) ([]*userDomain.User, error) {
	followers, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(user.UserId.String())).Find()
	if err != nil {
		return nil, err
	}

	var followerUserIdList []userDomain.UserId
	for _, follower := range followers {
		followerUserIdList = append(followerUserIdList, userDomain.UserId(follower.FollowerUserID))
	}

	return r.FindByIds(followerUserIdList)
}

func (r *UserRepository) MutualFollowUserList(user *userDomain.User) ([]*userDomain.User, error) {
	followers, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(user.UserId.String())).Find()
	if err != nil {
		return nil, err
	}

	var followerUserIdList []string
	for _, follower := range followers {
		followerUserIdList = append(followerUserIdList, follower.FollowerUserID)
	}

	mutualFollowings, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(user.UserId.String())).Where(query.Follow.FollowingUserID.In(followerUserIdList...)).Find()
	if err != nil {
		return nil, err
	}

	var mutualFollowUserIdList []userDomain.UserId
	for _, mutualFollowing := range mutualFollowings {
		mutualFollowUserIdList = append(mutualFollowUserIdList, userDomain.UserId(mutualFollowing.FollowingUserID))
	}

	return r.FindByIds(mutualFollowUserIdList)
}

func (r *UserRepository) FollowRequestUserList(user *userDomain.User) ([]*userDomain.User, error) {
	followings, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(user.UserId.String())).Find()
	if err != nil {
		return nil, err
	}

	var followingUserIdList []string
	for _, following := range followings {
		followingUserIdList = append(followingUserIdList, following.FollowingUserID)
	}

	followRequests, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(user.UserId.String())).Where(query.Follow.FollowerUserID.NotIn(followingUserIdList...)).Find()
	if err != nil {
		return nil, err
	}

	var followRequestUserIdList []userDomain.UserId
	for _, followRequest := range followRequests {
		followRequestUserIdList = append(followRequestUserIdList, userDomain.UserId(followRequest.FollowerUserID))
	}

	return r.FindByIds(followRequestUserIdList)
}
