package queue_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"github.com/itmayziii/fm_last_algo_course/pkg/queue"
	"testing"
)

func TestQueue_Enqueue(t *testing.T) {
	type change struct {
		enqueue  string
		expected string
	}

	tests := []struct {
		queue   *queue.Queue
		changes []change
	}{
		{
			queue.NewQueue(newLinkedList("A", "B", "C")),
			[]change{
				{enqueue: "D", expected: "A -> B -> C -> D"},
				{enqueue: "E", expected: "A -> B -> C -> D -> E"},
				{enqueue: "Z", expected: "A -> B -> C -> D -> E -> Z"},
			},
		},
		{
			queue.NewQueue(newLinkedList()),
			[]change{
				{enqueue: "X", expected: "X"},
				{enqueue: "Y", expected: "X -> Y"},
				{enqueue: "Z", expected: "X -> Y -> Z"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			for _, c := range tt.changes {
				tt.queue.Enqueue(c.enqueue)
				actual := tt.queue.String()
				if actual != c.expected {
					t.Errorf("actual: %s, expected %s", actual, c.expected)
				}
			}
		})
	}
}

func TestQueue_Deque(t *testing.T) {
	type expected struct {
		value     string
		traversal string
	}

	type change struct {
		expected expected
	}

	tests := []struct {
		queue   *queue.Queue
		changes []change
	}{
		{
			queue.NewQueue(newLinkedList("A", "B", "C", "X", "Y", "Z")),
			[]change{
				{expected{"A", "B -> C -> X -> Y -> Z"}},
				{expected{"B", "C -> X -> Y -> Z"}},
				{expected{"C", "X -> Y -> Z"}},
				{expected{"X", "Y -> Z"}},
				{expected{"Y", "Z"}},
				{expected{"Z", ""}},
				{expected{"", ""}},
			},
		},
		{
			queue.NewQueue(newLinkedList()),
			[]change{
				{expected{"", ""}},
				{expected{"", ""}},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			for _, c := range tt.changes {
				actual := tt.queue.Deque()
				if actual != c.expected.value {
					t.Errorf("actual: %s, expected %s", actual, c.expected.value)
				}
				actualTraversal := tt.queue.String()
				if actualTraversal != c.expected.traversal {
					t.Errorf("actual: %s, expected %s", actualTraversal, c.expected.traversal)
				}
			}
		})
	}
}

// newLinkedList returns a new linked list as some of the linked list methods are mutating, and we don't want
// to have shared state between test cases.
func newLinkedList(values ...string) *linked_list.Singly {
	ll := &linked_list.Singly{}

	for _, value := range values {
		ll.Push(value)
	}

	return ll
}
