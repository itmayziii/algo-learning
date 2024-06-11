package queue

import "github.com/itmayziii/fm_last_algo_course/pkg/linked_list"

// Queue implements a Queue data structure which implements the FIFO principle.
type Queue struct {
	list *linked_list.Singly
}

func NewQueue(list *linked_list.Singly) *Queue {
	return &Queue{list: list}
}

// Enqueue adds item to the queue at the end of the list.
func (q *Queue) Enqueue(value string) {
	q.list.Push(value)
}

// Deque returns the first value in the queue and removes it from the list.
func (q *Queue) Deque() string {
	return q.list.Shift()
}

// Peek shows the next value in the queue without removing it from the list.
func (q *Queue) Peek() string {
	return q.list.First()
}

func (q *Queue) String() string {
	return q.list.String()
}
