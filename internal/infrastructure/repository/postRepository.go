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
	poster, _ := userDomain.Factory.GetUser(userDomain.UserId(gormPost.UserID))

	return &postDomain.Post{
		PostId:    postDomain.PostId(gormPost.ID),
		Poster:    *poster,
		Content:   postDomain.Content(gormPost.Content),
		CreatedAt: gormPost.CreatedAt,
	}
}

func (r *PostRepository) Create(post *postDomain.Post) error {
	gormPost := toGormPost(post)
	gormPost.ID = xid.New().String()

	return query.Post.WithContext(context.Background()).Create(gormPost)
}

func (r *PostRepository) FindById(postId postDomain.PostId) (*postDomain.Post, error) {
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
