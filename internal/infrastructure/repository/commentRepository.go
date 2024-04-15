package repository

import (
	"context"

	"github.com/n4mlz/sns-backend/internal/domain/commentDomain"
	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/model"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/query"
)

type CommentRepository struct{}

func toGormComment(comment *commentDomain.Comment) *model.Comment {
	return &model.Comment{
		ID:        comment.CommentId.String(),
		PostID:    comment.PostId.String(),
		UserID:    comment.Commenter.UserId.String(),
		Content:   comment.Content.String(),
		CreatedAt: comment.CreatedAt,
	}
}

func toComment(gormComment *model.Comment) *commentDomain.Comment {
	// TODO: fix N+1 problem
	commenter, _ := userDomain.Factory.GetUser(userDomain.UserId(gormComment.UserID))

	// TODO: fix N+1 problem
	gormReplies, _ := query.Reply.WithContext(context.Background()).Where(query.Reply.CommentID.Eq(gormComment.ID)).Find()

	var replies []*commentDomain.Reply
	for _, gormReply := range gormReplies {
		// TODO: fix N+1 problem
		replies = append(replies, toReply(gormReply))
	}

	return &commentDomain.Comment{
		CommentId: commentDomain.CommentId(gormComment.ID),
		PostId:    postDomain.PostId(gormComment.PostID),
		Commenter: commenter,
		Content:   commentDomain.Content(gormComment.Content),
		Replies:   replies,
		CreatedAt: gormComment.CreatedAt,
	}
}

func toGormReply(reply *commentDomain.Reply) *model.Reply {
	return &model.Reply{
		ID:        reply.ReplyId.String(),
		CommentID: reply.CommentId.String(),
		UserID:    reply.Replier.UserId.String(),
		Sequence:  reply.Sequence.Int32(),
		Content:   reply.Content.String(),
		CreatedAt: reply.CreatedAt,
	}
}

func toReply(gormReply *model.Reply) *commentDomain.Reply {
	// TODO: fix N+1 problem
	replier, _ := userDomain.Factory.GetUser(userDomain.UserId(gormReply.UserID))

	return &commentDomain.Reply{
		ReplyId:   commentDomain.ReplyId(gormReply.ID),
		CommentId: commentDomain.CommentId(gormReply.CommentID),
		Replier:   replier,
		Sequence:  commentDomain.Sequence(gormReply.Sequence),
		Content:   commentDomain.Content(gormReply.Content),
		CreatedAt: gormReply.CreatedAt,
	}
}

func (r *CommentRepository) CreateComment(comment *commentDomain.Comment) error {
	gormComment := toGormComment(comment)
	return query.Comment.WithContext(context.Background()).Save(gormComment)
}

func (r *CommentRepository) DeleteComment(comment *commentDomain.Comment) error {
	_, err := query.Comment.WithContext(context.Background()).Where(query.Comment.ID.Eq(comment.CommentId.String())).Delete()
	return err
}

func CreateReply(reply *commentDomain.Reply) error {
	gormReply := toGormReply(reply)
	return query.Reply.WithContext(context.Background()).Save(gormReply)
}

func DeleteReply(reply *commentDomain.Reply) error {
	_, err := query.Reply.WithContext(context.Background()).Where(query.Reply.ID.Eq(reply.ReplyId.String())).Delete()
	return err
}

func (r *CommentRepository) FindByPostId(postId postDomain.PostId) ([]*commentDomain.Comment, error) {
	gormComments, err := query.Comment.WithContext(context.Background()).Where(query.Comment.PostID.Eq(postId.String())).Find()
	if err != nil {
		return nil, err
	}

	var comments []*commentDomain.Comment
	for _, gormComment := range gormComments {
		// TODO: fix N+1 problem
		comments = append(comments, toComment(gormComment))
	}

	return comments, nil
}

func (r *CommentRepository) IsExistCommentId(commentId commentDomain.CommentId) bool {
	count, _ := query.Comment.WithContext(context.Background()).Where(query.Comment.ID.Eq(commentId.String())).Count()
	return count != 0
}

func (r *CommentRepository) IsExistReplyId(replyId commentDomain.ReplyId) bool {
	count, _ := query.Reply.WithContext(context.Background()).Where(query.Reply.ID.Eq(replyId.String())).Count()
	return count != 0
}
