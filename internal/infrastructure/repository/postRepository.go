package repository

import (
	"context"
	"errors"

	"github.com/n4mlz/sns-backend/internal/domain/postDomain"
	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/model"
	"github.com/n4mlz/sns-backend/internal/infrastructure/repository/query"
	"github.com/rs/xid"
)

type PostRepository struct{}

func toGormPost(post *postDomain.Post) *model.Post {
	return &model.Post{
		ID:        post.PostId.String(),
		UserID:    post.Poster.UserId.String(),
		Content:   post.Content.String(),
		CreatedAt: post.CreatedAt,
	}
}

func toPost(gormPost *model.Post) *postDomain.Post {
	// TODO: fix N+1 problem
	poster, _ := userDomain.Factory.GetUser(userDomain.UserId(gormPost.UserID))

	return &postDomain.Post{
		PostId:    postDomain.PostId(gormPost.ID),
		Poster:    poster,
		Content:   postDomain.Content(gormPost.Content),
		CreatedAt: gormPost.CreatedAt,
	}
}

func toGormComment(comment *postDomain.Comment) *model.Comment {
	return &model.Comment{
		ID:        comment.CommentId.String(),
		PostID:    comment.PostId.String(),
		UserID:    comment.Commenter.UserId.String(),
		Content:   comment.Content.String(),
		CreatedAt: comment.CreatedAt,
	}
}

func toComment(gormComment *model.Comment) *postDomain.Comment {
	// TODO: fix N+1 problem
	commenter, _ := userDomain.Factory.GetUser(userDomain.UserId(gormComment.UserID))

	// TODO: fix N+1 problem
	gormReplies, _ := query.Reply.WithContext(context.Background()).Where(query.Reply.CommentID.Eq(gormComment.ID)).Find()

	var replies []*postDomain.Reply
	for _, gormReply := range gormReplies {
		// TODO: fix N+1 problem
		replies = append(replies, toReply(gormReply))
	}

	replies = postDomain.Service.SortReplies(replies)

	return &postDomain.Comment{
		CommentId: postDomain.CommentId(gormComment.ID),
		PostId:    postDomain.PostId(gormComment.PostID),
		Commenter: commenter,
		Content:   postDomain.Content(gormComment.Content),
		Replies:   replies,
		CreatedAt: gormComment.CreatedAt,
	}
}

func toGormReply(reply *postDomain.Reply) *model.Reply {
	return &model.Reply{
		ID:        reply.ReplyId.String(),
		CommentID: reply.CommentId.String(),
		UserID:    reply.Replier.UserId.String(),
		Content:   reply.Content.String(),
		CreatedAt: reply.CreatedAt,
	}
}

func toReply(gormReply *model.Reply) *postDomain.Reply {
	// TODO: fix N+1 problem
	replier, _ := userDomain.Factory.GetUser(userDomain.UserId(gormReply.UserID))

	return &postDomain.Reply{
		ReplyId:   postDomain.ReplyId(gormReply.ID),
		CommentId: postDomain.CommentId(gormReply.CommentID),
		Replier:   replier,
		Content:   postDomain.Content(gormReply.Content),
		CreatedAt: gormReply.CreatedAt,
	}
}

func (r *PostRepository) Create(post *postDomain.Post) (*postDomain.Post, error) {
	gormPost := toGormPost(post)
	gormPost.ID = xid.New().String()

	err := query.Post.WithContext(context.Background()).Create(gormPost)
	return toPost(gormPost), err
}

func (r *PostRepository) Delete(post *postDomain.Post) error {
	_, err := query.Post.WithContext(context.Background()).Where(query.Post.ID.Eq(post.PostId.String())).Delete()
	return err
}

func (r *PostRepository) FindPostById(postId postDomain.PostId) (*postDomain.Post, error) {
	if !r.IsExistPostId(postId) {
		return nil, errors.New("post not found")
	}

	gormPost, err := query.Post.WithContext(context.Background()).Where(query.Post.ID.Eq(postId.String())).Take()
	return toPost(gormPost), err
}

func (r *PostRepository) IsExistPostId(postId postDomain.PostId) bool {
	count, _ := query.Post.WithContext(context.Background()).Where(query.Post.ID.Eq(postId.String())).Count()
	return count != 0
}

func (r *PostRepository) Like(post *postDomain.Post, user *userDomain.User) error {
	like := &model.Like{
		ID:     xid.New().String(),
		UserID: user.UserId.String(),
		PostID: post.PostId.String(),
	}

	return query.Like.WithContext(context.Background()).Save(like)
}

func (r *PostRepository) Unlike(post *postDomain.Post, user *userDomain.User) error {
	_, err := query.Like.WithContext(context.Background()).Where(query.Like.PostID.Eq(post.PostId.String())).Where(query.Like.UserID.Eq(user.UserId.String())).Delete()
	return err
}

func (r *PostRepository) IsLiked(post *postDomain.Post, user *userDomain.User) bool {
	count, _ := query.Like.WithContext(context.Background()).Where(query.Like.PostID.Eq(post.PostId.String())).Where(query.Like.UserID.Eq(user.UserId.String())).Count()
	return count != 0
}

func (r *PostRepository) GetLikeCount(post *postDomain.Post) (int, error) {
	count, err := query.Like.WithContext(context.Background()).Where(query.Like.PostID.Eq(post.PostId.String())).Count()
	return int(count), err
}

func (r *PostRepository) GetLikers(post *postDomain.Post) ([]*userDomain.User, error) {
	likes, err := query.Like.WithContext(context.Background()).Where(query.Like.PostID.Eq(post.PostId.String())).Find()
	if err != nil {
		return nil, err
	}

	var likerUserIdList []userDomain.UserId
	for _, like := range likes {
		likerUserIdList = append(likerUserIdList, userDomain.UserId(like.UserID))
	}

	return userDomain.Factory.GetUsers(likerUserIdList)
}

func (r *PostRepository) CreateComment(comment *postDomain.Comment) (*postDomain.Comment, error) {
	gormComment := toGormComment(comment)
	gormComment.ID = xid.New().String()

	err := query.Comment.WithContext(context.Background()).Save(gormComment)
	return toComment(gormComment), err
}

func (r *PostRepository) DeleteComment(comment *postDomain.Comment) error {
	_, err := query.Comment.WithContext(context.Background()).Where(query.Comment.ID.Eq(comment.CommentId.String())).Delete()
	return err
}

func (r *PostRepository) CreateReply(reply *postDomain.Reply) (*postDomain.Reply, error) {
	gormReply := toGormReply(reply)
	gormReply.ID = xid.New().String()

	err := query.Reply.WithContext(context.Background()).Save(gormReply)
	return toReply(gormReply), err
}

func (r *PostRepository) DeleteReply(reply *postDomain.Reply) error {
	_, err := query.Reply.WithContext(context.Background()).Where(query.Reply.ID.Eq(reply.ReplyId.String())).Delete()
	return err
}

func (r *PostRepository) FindCommentById(commentId postDomain.CommentId) (*postDomain.Comment, error) {
	if !r.IsExistCommentId(commentId) {
		return nil, errors.New("comment not found")
	}

	gormComment, err := query.Comment.WithContext(context.Background()).Where(query.Comment.ID.Eq(commentId.String())).Take()
	return toComment(gormComment), err
}

func (r *PostRepository) FindCommentsByPostId(postId postDomain.PostId) ([]*postDomain.Comment, error) {
	gormComments, err := query.Comment.WithContext(context.Background()).Where(query.Comment.PostID.Eq(postId.String())).Find()
	if err != nil {
		return nil, err
	}

	var comments []*postDomain.Comment
	for _, gormComment := range gormComments {
		// TODO: fix N+1 problem
		comments = append(comments, toComment(gormComment))
	}

	return comments, nil
}

func (r *PostRepository) FindReplyById(replyId postDomain.ReplyId) (*postDomain.Reply, error) {
	if !r.IsExistReplyId(replyId) {
		return nil, errors.New("reply not found")
	}

	gormReply, err := query.Reply.WithContext(context.Background()).Where(query.Reply.ID.Eq(replyId.String())).Take()
	return toReply(gormReply), err
}

func (r *PostRepository) IsExistCommentId(commentId postDomain.CommentId) bool {
	count, _ := query.Comment.WithContext(context.Background()).Where(query.Comment.ID.Eq(commentId.String())).Count()
	return count != 0
}

func (r *PostRepository) IsExistReplyId(replyId postDomain.ReplyId) bool {
	count, _ := query.Reply.WithContext(context.Background()).Where(query.Reply.ID.Eq(replyId.String())).Count()
	return count != 0
}
