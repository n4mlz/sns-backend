package utils

import (
	"net/url"
	"strings"
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

type TypedSet[T comparable] map[T]T

func NewTypedSet[T comparable]() *TypedSet[T] {
	return &TypedSet[T]{}
}

func (s *TypedSet[T]) Add(key T) {
	(*s)[key] = key
}

func (s *TypedSet[T]) Delete(key T) {
	_, ok := (*s)[key]
	if ok {
		delete(*s, key)
	}
}

func (s *TypedSet[T]) Contains(key T) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *TypedSet[T]) Values() []T {
	values := make([]T, 0, len(*s))
	for k := range *s {
		values = append(values, k)
	}
	return values
}

func MapValues[T comparable, U any](m map[T]U) []U {
	values := make([]U, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

func IsUrl(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func TrimWordGaps(s string) string {
	f := func(r rune) bool {
		return r == ' ' || r == '\t' || r == '\r'
	}

	s = strings.Join(strings.FieldsFunc(s, f), " ")

	f = func(r rune) bool {
		return r == '\n'
	}

	return strings.Join(strings.FieldsFunc(s, f), "\n")
}
