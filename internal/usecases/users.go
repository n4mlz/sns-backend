package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/utils"
)

func User(ctx *gin.Context) {
	sourceUserId := userDomain.UserId(ctx.GetString("userId"))
	sourceUser, err := userDomain.Factory.GetUser(sourceUserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetUserName := userDomain.UserName(ctx.Param("userName"))
	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	visibleUserCount, err := targetUser.GetVisibleUserCount()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mutualCount := visibleUserCount - 1

	response := UserDetailDto{
		UserName:        targetUser.UserName.String(),
		DisplayName:     targetUser.DisplayName.String(),
		Biography:       targetUser.Biography.String(),
		CreatedAt:       targetUser.CreatedAt,
		Mutuals:         mutualCount,
		FollowingStatus: sourceUser.GetFollowingStatus(targetUser),
	}

	ctx.JSON(http.StatusOK, response)
}

func MutualFollow(ctx *gin.Context) {
	targetUserName := userDomain.UserName(ctx.Param("userName"))
	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceUserId := userDomain.UserId(ctx.GetString("userId"))
	sourceUser, err := userDomain.Factory.GetUser(sourceUserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceFollowingList, err := sourceUser.Followings()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceFollowerList, err := sourceUser.Followers()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	targetMutualList, err := targetUser.VisibleUsers()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followingSet := utils.NewSet()
	for _, user := range sourceFollowingList {
		followingSet.Add(user.UserId)
	}

	followerSet := utils.NewSet()
	for _, user := range sourceFollowerList {
		followerSet.Add(user.UserId)
	}

	var response []UserDto
	for _, user := range targetMutualList {
		followingStatus := userDomain.NONE

		if user.UserId == targetUser.UserId {
			continue
		}

		if user.UserId == sourceUser.UserId {
			followingStatus = userDomain.OWN
		} else {
			isFollowing := followingSet.Contains(user.UserId)

			isFollowed := followerSet.Contains(user.UserId)

			if isFollowing && isFollowed {
				followingStatus = userDomain.MUTUAL
			} else if isFollowing {
				followingStatus = userDomain.FOLLOWING
			} else if isFollowed {
				followingStatus = userDomain.FOLLOWED
			} else {
				followingStatus = userDomain.NONE
			}
		}

		response = append(response, UserDto{
			UserName:        user.UserName.String(),
			DisplayName:     user.DisplayName.String(),
			Biography:       user.Biography.String(),
			CreatedAt:       user.CreatedAt,
			FollowingStatus: followingStatus,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func UserPosts(ctx *gin.Context) {
	targetUserName := userDomain.UserName(ctx.Param("userName"))
	targetUser, err := userDomain.Factory.GetUserByUserName(targetUserName)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sourceUserId := userDomain.UserId(ctx.GetString("userId"))
	sourceUser, err := userDomain.Factory.GetUser(sourceUserId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	posts, err := postDomain.Factory.GetPostsByUser(sourceUser, targetUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poster := UserDisplayDto{
		UserName:    targetUser.UserName.String(),
		DisplayName: targetUser.DisplayName.String(),
	}

	var response []PostDto
	for _, post := range posts {
		// TODO: fix N+1 problem

		liked := post.IsLiked(sourceUser)

		likeCount, err := post.GetLikeCount()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		commentCount, err := post.GetCommentCount()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		response = append(response, PostDto{
			PostId:    post.PostId.String(),
			Poster:    poster,
			Content:   post.Content.String(),
			Likes:     likeCount,
			Liked:     liked,
			Comments:  commentCount,
			CreatedAt: post.CreatedAt,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
