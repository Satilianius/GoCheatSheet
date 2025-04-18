package main

type Set[T comparable] struct {
	items map[T]struct{}
}

func NewSet[T comparable]() *Set[T]  {
	return &Set[T]{
		items: make(map[T]struct{}),
	}
}

func (s *Set[T]) Add(item T) {
	s.items[item] = struct{}{}
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Contains(item T) bool {
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Empty() bool {
	return s.Size() == 0
}

func (s *Set[T]) Clear() {
	s.items = make(map[T]struct{})
}

func (s *Set[T]) Items() []T {
	items := make([]T, 0, s.Size())
	for item := range s.items {
		items = append(items, item)
	}
	return items
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		result.Add(item)
	}
	for item := range other.items {
		result.Add(item)
	}
	return result
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	result := NewSet[T]()

	// Determine which set is smaller
	smaller, larger := s, other
	if len(s.items) > len(other.items) {
		smaller, larger = other, s
	}

	// Iterate over the smaller set for better performance
	for item := range smaller.items {
		if larger.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T]()
	for item := range s.items {
		if !other.Contains(item) {
			result.Add(item)
		}
	}
	return result
}
