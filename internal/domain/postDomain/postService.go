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

func (ps *PostService) SortPosts(posts []*Post) []*Post {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].CreatedAt.Before(posts[j].CreatedAt)
	})
	return posts
}

func (ps *PostService) SortReplies(replies []*Reply) []*Reply {
	sort.Slice(replies, func(i, j int) bool {
		return replies[i].CreatedAt.Before(replies[j].CreatedAt)
	})
	return replies
}
