package postDomain

import (
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

type IPostRepository interface {
	Create(*Post) error
	Delete(*Post) error
	FindPostById(PostId) (*Post, error)
	IsExistPostId(PostId) bool
	Like(*Post, *userDomain.User) error
	Unlike(*Post, *userDomain.User) error
	IsLiked(*Post, *userDomain.User) bool
	GetLikeCount(*Post) (int, error)
	GetLikers(*Post) ([]*userDomain.User, error)
	CreateComment(comment *Comment) error
	DeleteComment(comment *Comment) error
	CreateReply(reply *Reply) error
	DeleteReply(reply *Reply) error
	FindCommentByPostId(postId PostId) ([]*Comment, error)
	IsExistCommentId(commentId CommentId) bool
	IsExistReplyId(replyId ReplyId) bool
}