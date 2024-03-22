package models

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/n4mlz/sns-backend/internal/repository/query"
)

var (
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 16
)

var MUTUAL = "mutual"
var FOLLOWING = "following"
var FOLLOWED = "followed"
var NONE = "none"
var OWN = "own"

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
	followers, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(userId)).Find()
	if err != nil {
		return nil, err
	}

	var followerUserIdList []string
	for _, follower := range followers {
		followerUserIdList = append(followerUserIdList, follower.FollowerUserID)
	}

	return followerUserIdList, nil
}

func getFollowingUserIdList(userId string) ([]string, error) {
	followings, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(userId)).Find()
	if err != nil {
		return nil, err
	}

	var followingUserIdList []string
	for _, following := range followings {
		followingUserIdList = append(followingUserIdList, following.FollowingUserID)
	}

	return followingUserIdList, nil
}

func getMutualFollowUserIdList(userId string) ([]string, error) {
	followerUserIdList, err := getFollowerUserIdList(userId)
	if err != nil {
		return nil, err
	}

	mutualFollowings, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowerUserID.Eq(userId)).Where(query.Follow.FollowingUserID.In(followerUserIdList...)).Find()
	if err != nil {
		return nil, err
	}

	var mutualFollowingUserIdList []string
	for _, follow := range mutualFollowings {
		mutualFollowingUserIdList = append(mutualFollowingUserIdList, follow.FollowingUserID)
	}

	return mutualFollowingUserIdList, nil
}

func getFollowRequestUserIdList(userId string) ([]string, error) {
	followingUserIdList, err := getFollowingUserIdList(userId)
	if err != nil {
		return nil, err
	}

	followRequests, err := query.Follow.WithContext(context.Background()).Where(query.Follow.FollowingUserID.Eq(userId)).Where(query.Follow.FollowerUserID.NotIn(followingUserIdList...)).Find()
	if err != nil {
		return nil, err
	}

	var followRequestUserIdList []string
	for _, follow := range followRequests {
		followRequestUserIdList = append(followRequestUserIdList, follow.FollowerUserID)
	}

	log.Print(followingUserIdList)
	log.Print(followRequestUserIdList)

	return followRequestUserIdList, nil
}

func getFollowingStatus(fromUserId string, toUserId string) string {
	if fromUserId == toUserId {
		return OWN
	}

	following := isFollowing(fromUserId, toUserId)
	followed := isFollowing(toUserId, fromUserId)

	if following && followed {
		return MUTUAL
	} else if following {
		return FOLLOWING
	} else if followed {
		return FOLLOWED
	} else {
		return NONE
	}
}
