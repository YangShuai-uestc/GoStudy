package main

func main() {

}

// https://www.cnblogs.com/mafeng/p/10331572.html

func (s *Set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = struct{}{}
	}
}

func (s *Set) Size() int {
	return len(s.m)
}

func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

func (s *Set) Equal(b *Set) bool {
	if s.Size() != b.Size() {
		return false
	}

	for item := range s.m {
		if !s.Contains(item) {
			return false
		}
	}

	return true
}

func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

func New(items ...interface{}) *Set {
	s := &Set{
		m: make(map[interface{}]struct{}),
	}
	s.Add(items)
	return s
}

// go set 实现
type Set struct {
	m map[interface{}]struct{}
}
