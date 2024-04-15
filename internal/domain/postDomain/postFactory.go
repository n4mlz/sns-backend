package postDomain

import (
	"errors"
	"time"

	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
)

var Factory *PostFactory

type PostFactory struct {
	postRepository *IPostRepository
}

func NewPostFactory(postRepository IPostRepository) *PostFactory {
	return &PostFactory{
		postRepository: &postRepository,
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

	err := (*pf.postRepository).Create(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (pf *PostFactory) GetPost(sourceUser *userDomain.User, postId PostId) (*Post, error) {
	post, err := (*pf.postRepository).FindPostById(postId)
	if err != nil {
		return nil, err
	}

	if !post.Poster.IsMutualFollow(sourceUser) {
		return nil, errors.New("permission denied")
	}

	post.PostRepository = pf.postRepository

	return post, nil
}

func (pf *PostFactory) DeletePostFromRepository(sourceUser *userDomain.User, post *Post) error {
	if sourceUser.UserId != post.Poster.UserId {
		return errors.New("not permitted")
	}
	return (*pf.postRepository).Delete(post)
}

func (pf *PostFactory) CreateCommentToRepository(post *Post, commenter *userDomain.User, content Content) (*Comment, error) {
	if !content.IsValid() {
		return nil, errors.New("invalid content")
	}

	if !post.Poster.IsMutualFollow(commenter) {
		return nil, errors.New("permission denied")
	}

	comment := &Comment{
		PostId:    post.PostId,
		Commenter: commenter,
		Content:   content,
	}

	err := (*pf.postRepository).CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (pf *PostFactory) GetComments(sourceUser *userDomain.User, post *Post) ([]*Comment, error) {
	comments, err := (*pf.postRepository).FindCommentByPostId(post.PostId)
	if err != nil {
		return nil, err
	}

	var result []*Comment
	for _, comment := range comments {
		if comment.Commenter.IsMutualFollow(sourceUser) {
			for _, reply := range comment.Replies {
				if !reply.Replier.IsMutualFollow(sourceUser) {
					reply.Replier = nil
					reply.Content = ""
					reply.CreatedAt = time.Time{}
				}
				result = append(result, comment)
			}
		}
	}
	return result, nil
}

func (pf *PostFactory) DeleteCommentFromRepository(sourceUser *userDomain.User, comment *Comment) error {
	if sourceUser.UserId != comment.Commenter.UserId {
		return errors.New("permission denied")
	}

	return (*pf.postRepository).DeleteComment(comment)
}

func (pf *PostFactory) CreateReplyToRepository(comment *Comment, replier *userDomain.User, content Content) (*Reply, error) {
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

	err := (*pf.postRepository).CreateReply(reply)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (pf *PostFactory) DeleteReplyFromRepository(sourceUser *userDomain.User, reply *Reply) error {
	if sourceUser.UserId != reply.Replier.UserId {
		return errors.New("permission denied")
	}

	return (*pf.postRepository).DeleteReply(reply)
}
