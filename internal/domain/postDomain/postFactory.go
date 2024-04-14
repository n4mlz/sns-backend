package postDomain

import (
	"errors"

	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

var Factory *PostFactory

type PostFactory struct {
	postRepository IPostRepository
}

func NewPostFactory(postRepository IPostRepository) *PostFactory {
	return &PostFactory{
		postRepository: postRepository,
	}
}

func SetDefaultPostFactory(postFactory *PostFactory) {
	Factory = postFactory
}

func (pf *PostFactory) CreatePostToRepository(poster *userDomain.User, content Content) (*Post, error) {
	if !content.IsValid() {
		return nil, errors.New("invalid content")
	}

	post := &Post{
		PostRepository: pf.postRepository,
		Poster:         poster,
		Content:        content,
	}

	err := pf.postRepository.Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pf *PostFactory) GetPost(postId PostId) (*Post, error) {
	post, err := pf.postRepository.FindById(postId)
	if err != nil {
		return nil, err
	}

	return &Post{
		PostRepository: pf.postRepository,
		PostId:         post.PostId,
		Poster:         post.Poster,
		Content:        post.Content,
		CreatedAt:      post.CreatedAt,
	}, nil
}

func (pf *PostFactory) DeletePostFromRepository(sourceUser *userDomain.User, post *Post) error {
	if sourceUser.UserId != post.Poster.UserId {
		return errors.New("not permitted")
	}
	return pf.postRepository.Delete(post)
}
