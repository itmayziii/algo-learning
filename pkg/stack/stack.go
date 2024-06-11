package stack

import "github.com/itmayziii/fm_last_algo_course/pkg/linked_list"

// Stack implements a Queue data structure which implements the LIFO principle. This implementation uses
// the beginning of a linked list instead of the end to represent the next item the stack should grab. This
// is because we have a singly linked list and the Shift and Unshift operations are O(1) where as the Pop operation is
// O(N) and the Push operation is only O(1) because we are keeping a reference to the tail pointer.
type Stack struct {
	list *linked_list.Singly
}

func NewStack(list *linked_list.Singly) *Stack {
	return &Stack{list: list}
}

func (s *Stack) Push(value string) {
	s.list.Unshift(value)
}

func (s *Stack) Pop() string {
	return s.list.Shift()
}

func (s *Stack) Peek() string {
	return s.list.First()
}

func (s *Stack) String() string {
	return s.list.String()
}
