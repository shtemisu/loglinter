package stack

import (
	"errors"
)

type Stack[T any] struct {
	Items []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		Items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.Items = append(s.Items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}

	lastIndex := len(s.Items) - 1
	item := s.Items[lastIndex]

	s.Items[lastIndex] = zero
	s.Items = s.Items[:lastIndex]

	return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("стек пуст")
	}

	return s.Items[len(s.Items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.Items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.Items)
}

func (s *Stack[T]) Clear() {
	s.Items = make([]T, 0)
}

func (s *Stack[T]) Values() []T {
	result := make([]T, len(s.Items))
	copy(result, s.Items)
	return result
}
