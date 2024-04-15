package postDomain

import "sort"

var Service *PostService

type PostService struct {
	PostRepository *IPostRepository
}

func NewPostService(postRepository IPostRepository) *PostService {
	return &PostService{PostRepository: &postRepository}
}

func SetDefaultPostService(postService *PostService) {
	Service = postService
}

func (ps *PostService) SortReplies(replies []*Reply) []*Reply {
	sort.Slice(replies, func(i, j int) bool {
		return replies[i].CreatedAt.Before(replies[j].CreatedAt)
	})

	return replies
}
