package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

func CreatePost(ctx *gin.Context) {
	var request PostContentDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content := postDomain.Content(request.Content)

	post, err := postDomain.Factory.CreatePostToRepository(user, content)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poster := UserDisplayDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
	}

	response := PostDto{
		PostId:    post.PostId.String(),
		Poster:    poster,
		Content:   post.Content.String(),
		Likes:     0,
		Liked:     false,
		Comments:  0,
		CreatedAt: post.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func DeletePost(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(ctx.Param("postId"))

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postDomain.Factory.DeletePostFromRepository(user, post)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func LikePost(ctx *gin.Context) {
	var request PostIdDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(request.PostId)

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = post.Like(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := PostIdDto{
		PostId: post.PostId.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func UnlikePost(ctx *gin.Context) {
	var request PostIdDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(request.PostId)

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = post.Unlike(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := PostIdDto{
		PostId: post.PostId.String(),
	}

	ctx.JSON(http.StatusOK, response)
}

func Likes(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(ctx.Param("postId"))

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	likers, err := post.GetLikers(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []UserDto
	for _, liker := range likers {
		followingStatus := userDomain.MUTUAL
		if liker.UserId == user.UserId {
			followingStatus = userDomain.OWN
		}

		response = append(response, UserDto{
			UserName:        liker.UserName.String(),
			DisplayName:     liker.DisplayName.String(),
			Biography:       liker.Biography.String(),
			CreatedAt:       liker.CreatedAt,
			FollowingStatus: followingStatus,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

func GetPost(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(ctx.Param("postId"))

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	liked := post.IsLiked(user)
	likeCount, err := post.GetLikeCount()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	poster := UserDisplayDto{
		UserName:    post.Poster.UserName.String(),
		DisplayName: post.Poster.DisplayName.String(),
	}

	commentObjects, err := post.GetComments(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var comments []CommentDto
	for _, comment := range commentObjects {
		commenter := UserDisplayDto{
			UserName:    comment.Commenter.UserName.String(),
			DisplayName: comment.Commenter.DisplayName.String(),
		}

		var replies []ReplyDto
		for _, reply := range comment.Replies {
			replier := UserDisplayDto{
				UserName:    reply.Replier.UserName.String(),
				DisplayName: reply.Replier.DisplayName.String(),
			}

			replies = append(replies, ReplyDto{
				ReplyId:   reply.ReplyId.String(),
				Replier:   replier,
				Content:   reply.Content.String(),
				CreatedAt: reply.CreatedAt,
			})
		}

		comments = append(comments, CommentDto{
			CommentId: comment.CommentId.String(),
			Commenter: commenter,
			Content:   comment.Content.String(),
			Replies:   replies,
			CreatedAt: comment.CreatedAt,
		})
	}

	response := PostDetailDto{
		PostId:    post.PostId.String(),
		Poster:    poster,
		Content:   post.Content.String(),
		Likes:     likeCount,
		Liked:     liked,
		Comments:  comments,
		CreatedAt: post.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func CreateComment(ctx *gin.Context) {
	var request CreateCommentDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	postId := postDomain.PostId(request.PostId)

	post, err := postDomain.Factory.GetPost(user, postId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content := postDomain.Content(request.Content)

	comment, err := postDomain.Factory.CreateCommentToRepository(post, user, content)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commenter := UserDisplayDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
	}

	response := CommentDto{
		CommentId: comment.CommentId.String(),
		Commenter: commenter,
		Content:   comment.Content.String(),
		Replies:   []ReplyDto{},
		CreatedAt: comment.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func DeleteComment(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentId := postDomain.CommentId(ctx.Param("commentId"))

	comment, err := postDomain.Factory.GetComment(user, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postDomain.Factory.DeleteCommentFromRepository(user, comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func CreateReply(ctx *gin.Context) {
	var request CreateReplyDto

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentId := postDomain.CommentId(request.CommentId)

	comment, err := postDomain.Factory.GetComment(user, commentId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	content := postDomain.Content(request.Content)

	reply, err := postDomain.Factory.CreateReplyToRepository(comment, user, content)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replier := UserDisplayDto{
		UserName:    user.UserName.String(),
		DisplayName: user.DisplayName.String(),
	}

	response := ReplyDto{
		ReplyId:   reply.ReplyId.String(),
		Replier:   replier,
		Content:   reply.Content.String(),
		CreatedAt: reply.CreatedAt,
	}

	ctx.JSON(http.StatusOK, response)
}

func DeleteReply(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replyId := postDomain.ReplyId(ctx.Param("replyId"))

	reply, err := postDomain.Factory.GetReply(user, replyId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = postDomain.Factory.DeleteReplyFromRepository(user, reply)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func Timeline(ctx *gin.Context) {
	user, err := userDomain.Factory.GetUser(userDomain.UserId(ctx.GetString("userId")))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	timeline, err := postDomain.Factory.GetPostsByVisibleUsers(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response []PostDto
	for _, post := range timeline {
		// TODO: fix N+1 problem

		liked := post.IsLiked(user)

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

		poster := UserDisplayDto{
			UserName:    post.Poster.UserName.String(),
			DisplayName: post.Poster.DisplayName.String(),
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
