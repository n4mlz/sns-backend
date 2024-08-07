package userDomain

import (
	"fmt"
	"regexp"
	"time"
	"unicode/utf8"

	"github.com/n4mlz/sns-backend/internal/utils"
)

const (
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 16
)

const (
	MIN_DISPLAY_NAME_LENGTH = 1
	MAX_DISPLAY_NAME_LENGTH = 32
)

const MAX_BIOGRAPHY_LENGTH = 256

const (
	MUTUAL    = "mutual"
	FOLLOWING = "following"
	FOLLOWED  = "followed"
	NONE      = "none"
	OWN       = "own"
)

var (
	NonVisibleUserId      = UserId("")
	NonVisibleUserName    = UserName("")
	NonVisibleDisplayName = DisplayName("相互フォローでないユーザー")
	NonVisibleBiography   = Biography("")
	NonVisibleIconUrl     = ImageUrl("")
	NonVisibleBgImageUrl  = ImageUrl("")
	NonVisibleCreatedAt   = time.Time{}
)

type UserId string

func (u UserId) String() string {
	return string(u)
}

type UserName string

func (u UserName) String() string {
	return string(u)
}

func (u UserName) IsValid() bool {
	pattern := regexp.MustCompile(fmt.Sprintf(`^[a-zA-Z0-9_]{%d,%d}$`, MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH))
	return pattern.MatchString(u.String())
}

type DisplayName string

func (d DisplayName) String() string {
	return string(d)
}

func (d DisplayName) IsValid() bool {
	return MIN_DISPLAY_NAME_LENGTH <= utf8.RuneCountInString(d.String()) && utf8.RuneCountInString(d.String()) <= MAX_DISPLAY_NAME_LENGTH
}

type Biography string

func (b Biography) String() string {
	return string(b)
}

func (b Biography) IsValid() bool {
	return utf8.RuneCountInString(b.String()) <= MAX_BIOGRAPHY_LENGTH
}

type ImageUrl string

func (i ImageUrl) String() string {
	return string(i)
}

func (i ImageUrl) IsValid() bool {
	return utils.IsUrl(i.String())
}
