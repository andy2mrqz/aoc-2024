package set

type Set[T comparable] map[T]bool

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T])
	for _, item := range items {
		s.Add(item)
	}
	return s
}

func (s Set[T]) Add(items ...T) {
	for _, item := range items {
		s[item] = true
	}
}

func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Slice() []T {
	ns := make([]T, 0, len(s))
	for item := range s {
		ns = append(ns, item)
	}
	return ns
}

func (s1 Set[T]) Intersection(s2 Set[T]) Set[T] {
	x := make(Set[T])
	for i := range s1 {
		if s2.Has(i) {
			x.Add(i)
		}
	}
	return x
}
