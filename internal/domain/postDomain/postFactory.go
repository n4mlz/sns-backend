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

	post, err := (*pf.postRepository).Create(post)
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

	if !post.Poster.IsVisible(sourceUser) {
		return nil, errors.New("permission denied")
	}

	post.PostRepository = pf.postRepository

	return post, nil
}

func (pf *PostFactory) GetPostsByUser(sourceUser *userDomain.User, targetUser *userDomain.User, cursor PostId, limit int) ([]*Post, PostId, error) {
	if !targetUser.IsVisible(sourceUser) {
		return nil, "", errors.New("permission denied")
	}

	if !(1 <= limit && limit <= MAX_CURSOR_PAGENATION_LIMIT) {
		return nil, "", errors.New("invalid limit")
	}

	posts, nextCursor, err := (*pf.postRepository).FindPostsByUserId(targetUser.UserId, cursor, limit)
	if err != nil {
		return nil, "", err
	}

	var result []*Post
	for _, post := range posts {
		post.PostRepository = pf.postRepository
		result = append(result, post)
	}

	return result, nextCursor, nil
}

func (pf *PostFactory) GetPostsByVisibleUsers(sourceUser *userDomain.User, cursor PostId, limit int) ([]*Post, PostId, error) {
	if !(1 <= limit && limit <= MAX_CURSOR_PAGENATION_LIMIT) {
		return nil, "", errors.New("invalid limit")
	}

	visibleUsers, err := sourceUser.VisibleUsers()
	if err != nil {
		return nil, "", err
	}

	var userIds []userDomain.UserId
	for _, user := range visibleUsers {
		userIds = append(userIds, user.UserId)
	}

	posts, nextCursor, err := (*pf.postRepository).FindPostsByUserIds(userIds, cursor, limit)
	if err != nil {
		return nil, "", err
	}

	var result []*Post
	for _, post := range posts {
		post.PostRepository = pf.postRepository
		result = append(result, post)
	}

	return result, nextCursor, nil
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

	if !post.Poster.IsVisible(commenter) {
		return nil, errors.New("permission denied")
	}

	comment := &Comment{
		PostId:    post.PostId,
		Commenter: commenter,
		Content:   content,
	}

	comment, err := (*pf.postRepository).CreateComment(comment)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (pf *PostFactory) GetComment(sourceUser *userDomain.User, commentId CommentId) (*Comment, error) {
	comment, err := (*pf.postRepository).FindCommentById(commentId)
	if err != nil {
		return nil, err
	}

	if !comment.Commenter.IsVisible(sourceUser) {
		return nil, errors.New("permission denied")
	}

	for _, reply := range comment.Replies {
		// fix N+1 problem
		if !reply.Replier.IsVisible(sourceUser) {
			reply.Replier = nil
			reply.Content = ""
			reply.CreatedAt = time.Time{}
		}
	}

	return comment, nil
}

func (pf *PostFactory) GetComments(sourceUser *userDomain.User, post *Post) ([]*Comment, error) {
	comments, err := (*pf.postRepository).FindCommentsByPostId(post.PostId)
	if err != nil {
		return nil, err
	}

	var result []*Comment
	for _, comment := range comments {
		// TODO: fix N+1 problem
		if comment.Commenter.IsVisible(sourceUser) {
			for _, reply := range comment.Replies {
				// TODO: fix N+1 problem
				if !reply.Replier.IsVisible(sourceUser) {
					reply.Replier = nil
					reply.Content = ""
					reply.CreatedAt = time.Time{}
				}
			}
			result = append(result, comment)
		}
	}
	return result, nil
}

func (pf *PostFactory) GetReply(sourceUser *userDomain.User, replyId ReplyId) (*Reply, error) {
	reply, err := (*pf.postRepository).FindReplyById(replyId)
	if err != nil {
		return nil, err
	}

	if !reply.Replier.IsVisible(sourceUser) {
		return nil, errors.New("permission denied")
	}

	return reply, nil
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

	if !comment.Commenter.IsVisible(replier) {
		return nil, errors.New("permission denied")
	}

	reply := &Reply{
		CommentId: comment.CommentId,
		Replier:   replier,
		Content:   content,
	}

	reply, err := (*pf.postRepository).CreateReply(reply)
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
