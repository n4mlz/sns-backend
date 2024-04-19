package utils

import (
	"net/url"
	"path"
)

type Set map[interface{}]struct{}

func NewSet() *Set {
	return &Set{}
}

func (s Set) Add(key interface{}) {
	s[key] = struct{}{}
}

func (s Set) Delete(key interface{}) {
	_, ok := s[key]
	if ok {
		delete(s, key)
	}
}

func (s Set) Contains(key interface{}) bool {
	_, ok := s[key]
	return ok
}

func JoinPath(srcUrl string, paths ...string) string {
	u, _ := url.Parse(srcUrl)

	u.Path = path.Join(append([]string{u.Path}, paths...)...)

	return u.String()
}
