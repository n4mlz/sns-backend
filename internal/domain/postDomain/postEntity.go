package postDomain

import (
	"errors"
	"time"

	"github.com/n4mlz/sns-backend/internal/domain/userDomain"
	"github.com/n4mlz/sns-backend/internal/utils"
)

type Post struct {
	PostRepository *IPostRepository
	PostId         PostId
	Poster         *userDomain.User
	Content        Content
	CreatedAt      time.Time
}

func (p *Post) Like(user *userDomain.User) error {
	if !p.Poster.IsVisible(user) {
		return errors.New("permission denied")
	}

	if p.IsLiked(user) {
		return errors.New("already liked")
	}

	(*p.PostRepository).Like(p, user)
	return nil
}

func (p *Post) Unlike(user *userDomain.User) error {
	if !p.Poster.IsVisible(user) {
		return errors.New("permission denied")
	}

	if !p.IsLiked(user) {
		return errors.New("not liked")
	}

	(*p.PostRepository).Unlike(p, user)
	return nil
}

func (p *Post) IsLiked(user *userDomain.User) bool {
	return (*p.PostRepository).IsLiked(p, user)
}

func (p *Post) GetLikeCount() (int, error) {
	return (*p.PostRepository).GetLikeCount(p)
}

func (p *Post) GetLikers(sourceUser *userDomain.User) ([]*userDomain.User, error) {
	likers, err := (*p.PostRepository).GetLikers(p)
	if err != nil {
		return nil, err
	}

	visibleUsers, err := sourceUser.VisibleUsers()
	if err != nil {
		return nil, err
	}

	VisibleSet := utils.NewSet()
	for _, visibleUser := range visibleUsers {
		VisibleSet.Add(visibleUser.UserId)
	}

	var visibleLikers []*userDomain.User
	for _, liker := range likers {
		if VisibleSet.Contains(liker.UserId) {
			visibleLikers = append(visibleLikers, liker)
		}
	}

	return visibleLikers, nil
}

func (p *Post) GetComments(sourceUser *userDomain.User) ([]*Comment, error) {
	return Factory.GetComments(sourceUser, p)
}

type Comment struct {
	CommentId CommentId
	PostId    PostId
	Commenter *userDomain.User
	Content   Content
	Replies   []*Reply
	CreatedAt time.Time
}

type Reply struct {
	ReplyId   ReplyId
	CommentId CommentId
	Replier   *userDomain.User
	Content   Content
	CreatedAt time.Time
}
