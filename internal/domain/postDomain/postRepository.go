package postDomain

import (
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

type IPostRepository interface {
	Create(*Post) (*Post, error)
	Delete(*Post) error
	FindPostById(PostId) (*Post, error)
	FindLatestPostByUserId(userDomain.UserId) (*Post, error)
	FindPostsByUserId(userDomain.UserId, PostId, int) ([]*Post, PostId, error)
	FindPostsByUserIds([]userDomain.UserId, PostId, int) ([]*Post, PostId, error)
	IsExistPostId(PostId) bool
	Like(*Post, *userDomain.User) error
	Unlike(*Post, *userDomain.User) error
	IsLiked(*Post, *userDomain.User) bool
	GetLikeCount(*Post) (int, error)
	GetLikers(*Post) ([]*userDomain.User, error)
	GetCommentCount(*Post) (int, error)
	CreateComment(comment *Comment) (*Comment, error)
	DeleteComment(comment *Comment) error
	CreateReply(reply *Reply) (*Reply, error)
	DeleteReply(reply *Reply) error
	FindCommentById(commentId CommentId) (*Comment, error)
	FindCommentsByPostId(postId PostId) ([]*Comment, error)
	FindReplyById(replyId ReplyId) (*Reply, error)
	IsExistCommentId(commentId CommentId) bool
	IsExistReplyId(replyId ReplyId) bool
	CreatePostNotifications([]*PostNotification) ([]*PostNotification, error)
}
