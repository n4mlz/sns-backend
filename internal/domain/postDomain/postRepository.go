package postDomain

import "github.com/n4mlz/sns-backend/internal/domain/userDomain"

type IPostRepository interface {
	Create(*Post) error
	FindById(PostId) (*Post, error)
	IsExistPostId(PostId) bool
	Like(*Post, *userDomain.User) error
	Unlike(*Post, *userDomain.User) error
	IsLiked(*Post, *userDomain.User) bool
	GetLikeCount(*Post) (int, error)
	GetLikers(*Post) ([]*userDomain.User, error)
}
