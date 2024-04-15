package commentDomain

import (
	"errors"

	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

var Factory *CommentFactory

type CommentFactory struct {
	commentRepository *ICommentRepository
}

func NewCommentFactory(commentRepository ICommentRepository) *CommentFactory {
	return &CommentFactory{
		commentRepository: &commentRepository,
	}
}

func SetDefaultCommentFactory(commentFactory *CommentFactory) {
	Factory = commentFactory
}

func (cf *CommentFactory) CreateCommentToRepository(post *postDomain.Post, commenter *userDomain.User, content Content) (*Comment, error) {
	if !content.IsValid() {
		return nil, errors.New("invalid content")
	}

	if !post.Poster.IsMutualFollow(commenter) {
		return nil, errors.New("permission denied")
	}

	comment := &Comment{
		CommentRepository: cf.commentRepository,
		PostId:            post.PostId,
		Commenter:         commenter,
		Content:           content,
	}

	err := (*cf.commentRepository).CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (cf *CommentFactory) GetComments(postId postDomain.PostId) ([]*Comment, error) {
	comments, err := (*cf.commentRepository).FindByPostId(postId)
	if err != nil {
		return nil, err
	}

	for _, comment := range comments {
		comment.CommentRepository = cf.commentRepository
	}
	return comments, nil
}

func (cf *CommentFactory) DeleteCommentFromRepository(sourceUser *userDomain.User, comment *Comment) error {
	if sourceUser.UserId != comment.Commenter.UserId {
		return errors.New("permission denied")
	}

	return (*cf.commentRepository).DeleteComment(comment)
}

func (cf *CommentFactory) CreateReplyToRepository(comment *Comment, replier *userDomain.User, content Content) (*Reply, error) {
	if !content.IsValid() {
		return nil, errors.New("invalid content")
	}

	if !comment.Commenter.IsMutualFollow(replier) {
		return nil, errors.New("permission denied")
	}

	reply := &Reply{
		CommentId: comment.CommentId,
		Replier:   replier,
		Content:   content,
	}

	err := (*cf.commentRepository).CreateReply(reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (cf *CommentFactory) DeleteReplyFromRepository(sourceUser *userDomain.User, reply *Reply) error {
	if sourceUser.UserId != reply.Replier.UserId {
		return errors.New("permission denied")
	}

	return (*cf.commentRepository).DeleteReply(reply)
}
