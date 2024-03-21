package models

import (
	"context"
	"fmt"
	"regexp"

	"github.com/n4mlz/sns-backend/internal/repository/query"
)

var (
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 16
)

func isValidUserName(s string) bool {
	pattern := regexp.MustCompile(fmt.Sprintf(`^[A-Za-z0-9_]{%d,%d}$`, MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH))
	return pattern.MatchString(s)
}

func isValidDisplayName(s string) bool {
	return len(s) != 0
}

func isExistUser(userName string) bool {
	count, _ := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq(userName)).Count()
	return count != 0
}

func userNameToUserId(userName string) string {
	user, _ := query.User.WithContext(context.Background()).Where(query.User.UserName.Eq(userName)).Take()
	return user.ID
}

func userIdToUserName(userId string) string {
	user, _ := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(userId)).Take()
	return user.UserName
}

func isFollowing(fromUserId string, toUserId string) bool {
	count, _ := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(fromUserId)).Where(query.Follow.FollowingUserID.Eq(toUserId)).Count()
	return count != 0
}

func isMutualFollow(userId1 string, userId2 string) bool {
	return isFollowing(userId1, userId2) && isFollowing(userId2, userId1)
}

func getFollowerUserIdList(userId string) ([]string, error) {
	follower, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(userId)).Find()
	if err != nil {
		return nil, err
	}

	var followerUserIdList []string
	for _, follow := range follower {
		followerUserIdList = append(followerUserIdList, follow.FollowerUserID)
	}

	return followerUserIdList, nil
}

func getMutualFollowUserIdList(userId string) ([]string, error) {
	followerUserIdList, err := getFollowerUserIdList(userId)

	if err != nil {
		return nil, err
	}

	mutualFollowing, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(userId)).Where(query.Follow.FollowingUserID.In(followerUserIdList...)).Find()
	if err != nil {
		return nil, err
	}

	var mutualFollowingUserIdList []string
	for _, follow := range mutualFollowing {
		mutualFollowingUserIdList = append(mutualFollowingUserIdList, follow.FollowingUserID)
	}

	return mutualFollowingUserIdList, nil
}
