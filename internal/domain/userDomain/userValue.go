package userDomain

import (
	"fmt"
	"regexp"
)

const (
	MIN_USERNAME_LENGTH = 4
	MAX_USERNAME_LENGTH = 16
)

const MAX_BIOGRAPHY_LENGTH = 256

const (
	MUTUAL    = "mutual"
	FOLLOWING = "following"
	FOLLOWED  = "followed"
	NONE      = "none"
	OWN       = "own"
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
	pattern := regexp.MustCompile(fmt.Sprintf(`^[A-Za-z0-9_]{%d,%d}$`, MIN_USERNAME_LENGTH, MAX_USERNAME_LENGTH))
	return pattern.MatchString(u.String())
}

type DisplayName string

func (d DisplayName) String() string {
	return string(d)
}

func (d DisplayName) IsValid() bool {
	return len(d) != 0
}

type Biography string

func (b Biography) String() string {
	return string(b)
}

func (b Biography) IsValid() bool {
	return len(b) <= MAX_BIOGRAPHY_LENGTH
}
