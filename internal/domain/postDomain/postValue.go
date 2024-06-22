package postDomain

import "unicode/utf8"

const MAX_CONTENT_LENGTH = 256

const MAX_CURSOR_PAGENATION_LIMIT = 128

type PostId string

func (pid PostId) String() string {
	return string(pid)
}

type Content string

func (c Content) String() string {
	return string(c)
}

func (c Content) IsValid() bool {
	return utf8.RuneCountInString(c.String()) <= MAX_CONTENT_LENGTH
}

type CommentId string

func (cid CommentId) String() string {
	return string(cid)
}

type ReplyId string

func (rid ReplyId) String() string {
	return string(rid)
}
