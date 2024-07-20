package postDomain

import (
	"strings"
	"unicode/utf8"

	"github.com/n4mlz/sns-backend/internal/utils"
)

const MAX_CONTENT_LENGTH = 256

const MAX_CURSOR_PAGINATION_LIMIT = 128

const (
	COMMENT = "comment"
	REPLY   = "reply"
)

type PostId string

func (pid PostId) String() string {
	return string(pid)
}

type Content string

func (c Content) String() string {
	return string(c)
}

func (c Content) IsValid() bool {
	length := utf8.RuneCountInString(c.String())
	return 1 <= length && length <= MAX_CONTENT_LENGTH
}

func (c Content) TrimWordGaps() Content {
	return Content(strings.TrimSpace(utils.TrimWordGaps(c.String())))
}

type CommentId string

func (cid CommentId) String() string {
	return string(cid)
}

type ReplyId string

func (rid ReplyId) String() string {
	return string(rid)
}

type PostNotificationId string

func (pnid PostNotificationId) String() string {
	return string(pnid)
}

type NotificationType string

func (nt NotificationType) String() string {
	return string(nt)
}

func (nt NotificationType) IsValid() bool {
	return nt == COMMENT || nt == REPLY
}
