package commentDomain

const MAX_COMMENT_CONTENT_LENGTH = 256

type CommentId string

func (cid CommentId) String() string {
	return string(cid)
}

type Content string

func (c Content) String() string {
	return string(c)
}

func (c Content) IsValid() bool {
	return len(c) <= MAX_COMMENT_CONTENT_LENGTH
}

type ReplyId string

func (rid ReplyId) String() string {
	return string(rid)
}

type Sequence int

func (s Sequence) Int32() int32 {
	return int32(s)
}
