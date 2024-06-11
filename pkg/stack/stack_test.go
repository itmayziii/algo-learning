package stack_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"github.com/itmayziii/fm_last_algo_course/pkg/stack"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	type change struct {
		push     string
		expected string
	}

	tests := []struct {
		stack   *stack.Stack
		changes []change
	}{
		{
			stack.NewStack(newLinkedList("A", "B", "C")),
			[]change{
				{push: "D", expected: "D -> A -> B -> C"},
				{push: "E", expected: "E -> D -> A -> B -> C"},
				{push: "Z", expected: "Z -> E -> D -> A -> B -> C"},
			},
		},
		{
			stack.NewStack(newLinkedList()),
			[]change{
				{push: "X", expected: "X"},
				{push: "Y", expected: "Y -> X"},
				{push: "Z", expected: "Z -> Y -> X"},
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			for _, c := range tt.changes {
				tt.stack.Push(c.push)
				actual := tt.stack.String()
				if actual != c.expected {
					t.Errorf("actual: %s, expected %s", actual, c.expected)
				}
			}
		})
	}
}

func TestQueue_Pop(t *testing.T) {
	type expected struct {
		value     string
		traversal string
	}

	type change struct {
		expected expected
	}

	tests := []struct {
		stack   *stack.Stack
		changes []change
	}{
		{
			stack.NewStack(newLinkedList("A", "B", "C", "X", "Y", "Z")),
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
			stack.NewStack(newLinkedList()),
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
				actual := tt.stack.Pop()
				if actual != c.expected.value {
					t.Errorf("actual: %s, expected %s", actual, c.expected.value)
				}
				actualTraversal := tt.stack.String()
				if actualTraversal != c.expected.traversal {
					t.Errorf("actual traversal: %s, expected traversal %s", actualTraversal, c.expected.traversal)
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
