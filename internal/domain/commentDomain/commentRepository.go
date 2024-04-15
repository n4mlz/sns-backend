package commentDomain

import "github.com/n4mlz/sns-backend/internal/domain/postDomain"

type ICommentRepository interface {
	CreateComment(comment *Comment) error
	DeleteComment(comment *Comment) error
	CreateReply(reply *Reply) error
	DeleteReply(reply *Reply) error
	FindByPostId(postId postDomain.PostId) ([]*Comment, error)
	IsExistCommentId(commentId CommentId) bool
	IsExistReplyId(replyId ReplyId) bool
}
