package postDomain

const MAX_POST_CONTENT_LENGTH = 256

type PostId string

func (pid PostId) String() string {
	return string(pid)
}

type Content string

func (c Content) String() string {
	return string(c)
}

func (c Content) IsValid() bool {
	return len(c) <= MAX_POST_CONTENT_LENGTH
}
