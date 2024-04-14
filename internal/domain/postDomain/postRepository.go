package postDomain

import "github.com/n4mlz/sns-backend/internal/domain/userDomain"

type IPostRepository interface {
	Create(*Post) error
	FindById(PostId) (*Post, error)
	Like(*Post, userDomain.User) error
	Unlike(*Post, userDomain.User) error
	IsLiked(*Post, userDomain.User) bool
	GetLikeCount(*Post) int
	GetLikers(*Post) ([]userDomain.User, error)
}
