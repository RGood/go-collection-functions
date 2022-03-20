package collections

type Set[K comparable] struct {
	data map[K]struct{}
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{data: map[K]struct{}{}}
}

func (s *Set[K]) Add(value K) {
	s.data[value] = struct{}{}
}

func (s *Set[K]) Remove(value K) {
	delete(s.data, value)
}

func (s *Set[K]) Contains(value K) bool {
	_, ok := s.data[value]

	return ok
}
