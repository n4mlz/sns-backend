package commentDomain

import (
	"time"

	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

type Comment struct {
	CommentRepository *ICommentRepository
	CommentId         CommentId
	PostId            postDomain.PostId
	Commenter         *userDomain.User
	Content           Content
	Replies           []*Reply
	CreatedAt         time.Time
}

type Reply struct {
	ReplyId   ReplyId
	CommentId CommentId
	Replier   *userDomain.User
	Sequence  Sequence
	Content   Content
	CreatedAt time.Time
}
