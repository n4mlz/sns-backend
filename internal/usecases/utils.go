package usecases

// TODO: change directory

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
