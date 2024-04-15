package postDomain

const MAX_CONTENT_LENGTH = 256

type PostId string

func (pid PostId) String() string {
	return string(pid)
}

type Content string

func (c Content) String() string {
	return string(c)
}

func (c Content) IsValid() bool {
	return len(c) <= MAX_CONTENT_LENGTH
}

type CommentId string

func (cid CommentId) String() string {
	return string(cid)
}

type ReplyId string

func (rid ReplyId) String() string {
	return string(rid)
}

type Sequence int

func (s Sequence) Int32() int32 {
	return int32(s)
}
