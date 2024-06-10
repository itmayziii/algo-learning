package linked_list_test

import (
	"errors"
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"testing"
)

func TestNode_Traverse(t *testing.T) {
	tests := []struct {
		list     *linked_list.Node
		expected string
	}{
		{newLinkedList(), "A -> B -> C"},
		{newLinkedList().Next, "B -> C"},
		{newLinkedList().Next.Next, "C"},
		{nil, ""},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Traverse()
			if actual != tt.expected {
				t.Errorf("actual %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Search(t *testing.T) {
	head := &linked_list.Node{
		Data: []byte("A"),
	}
	second := &linked_list.Node{
		Data: []byte("B"),
	}
	third := &linked_list.Node{
		Data: []byte("C"),
	}
	head.Next = second
	second.Next = third

	tests := []struct {
		list     *linked_list.Node
		needle   []byte
		expected bool
	}{
		{newLinkedList(), []byte("A"), true},
		{newLinkedList(), []byte("B"), true},
		{newLinkedList(), []byte("C"), true},
		{newLinkedList(), []byte("D"), false},
		{newLinkedList().Next, []byte("A"), false},
		{newLinkedList().Next, []byte("B"), true},
		{newLinkedList().Next, []byte("C"), true},
		{newLinkedList().Next.Next, []byte("A"), false},
		{newLinkedList().Next.Next, []byte("B"), false},
		{newLinkedList().Next.Next, []byte("C"), true},
		{newLinkedList(), []byte("AA"), false},
		{newLinkedList(), []byte(""), false},
		{nil, []byte("A"), false},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Search(tt.needle)
			if actual != tt.expected {
				t.Errorf("actual %t, expected %t", actual, tt.expected)
			}
		})
	}
}

func TestNode_Length(t *testing.T) {
	head := &linked_list.Node{
		Data: []byte("A"),
	}
	second := &linked_list.Node{
		Data: []byte("B"),
	}
	third := &linked_list.Node{
		Data: []byte("C"),
	}
	head.Next = second
	second.Next = third

	tests := []struct {
		list     *linked_list.Node
		expected int
	}{
		{newLinkedList(), 3},
		{newLinkedList().Next, 2},
		{newLinkedList().Next.Next, 1},
		{nil, 0},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Length()
			if actual != tt.expected {
				t.Errorf("actual %d, expected %d", actual, tt.expected)
			}
		})
	}
}

func TestNode_Shift(t *testing.T) {
	tests := []struct {
		list     *linked_list.Node
		expected string
	}{
		{newLinkedList(), "B -> C"},
		{newLinkedList().Next, "C"},
		{newLinkedList().Next.Next, ""},
		{nil, ""},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Shift().Traverse()
			if actual != tt.expected {
				t.Errorf("actual: %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Unshift(t *testing.T) {
	tests := []struct {
		list     *linked_list.Node
		newHead  *linked_list.Node
		expected string
	}{
		{newLinkedList(), &linked_list.Node{Data: []byte("0")}, "0 -> A -> B -> C"},
		{newLinkedList().Next, &linked_list.Node{Data: []byte("0")}, "0 -> B -> C"},
		{newLinkedList().Next.Next, &linked_list.Node{Data: []byte("0")}, "0 -> C"},
		{nil, &linked_list.Node{Data: []byte("0")}, "0"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Unshift(tt.newHead).Traverse()
			if actual != tt.expected {
				t.Errorf("actual %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Push(t *testing.T) {
	tests := []struct {
		list     *linked_list.Node
		newTail  *linked_list.Node
		expected string
	}{
		{newLinkedList(), &linked_list.Node{Data: []byte("Z")}, "A -> B -> C -> Z"},
		{newLinkedList().Next, &linked_list.Node{Data: []byte("Z")}, "B -> C -> Z"},
		{newLinkedList().Next.Next, &linked_list.Node{Data: []byte("Z")}, "C -> Z"},
		{nil, &linked_list.Node{Data: []byte("Z")}, "Z"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Push(tt.newTail).Traverse()
			if actual != tt.expected {
				t.Errorf("actual %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	type expected struct {
		value string
		err   error
	}

	tests := []struct {
		list     *linked_list.Node
		data     []byte
		position int
		expected expected
	}{
		{newLinkedList(), []byte("AA"), 0, expected{"AA -> A -> B -> C", nil}},
		{newLinkedList(), []byte("BB"), 1, expected{"A -> BB -> B -> C", nil}},
		{newLinkedList(), []byte("CC"), 2, expected{"A -> B -> CC -> C", nil}},
		{newLinkedList(), []byte("D"), 3, expected{"A -> B -> C -> D", nil}},
		{newLinkedList(), []byte("E"), 4, expected{
			"A -> B -> C",
			errors.New("position: 4 is not in range of linked list"),
		}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			head, err := tt.list.Insert(tt.data, tt.position)
			if tt.expected.err == nil {
				if err != nil {
					t.Errorf("unexpected error %v", err)
				}

				actual := head.Traverse()
				if actual != tt.expected.value {
					t.Errorf("actual: %s, expected %s", actual, tt.expected.value)
				}
			} else {
				if err != nil && tt.expected.err.Error() != err.Error() {
					t.Errorf("actual error: %v, expected error: %v", err, tt.expected.err)
				}
			}
		})
	}
}

// newLinkedList returns a new linked list as some of the linked list methods are mutating and we don't want
// to have shared state between test cases.
func newLinkedList() *linked_list.Node {
	head := &linked_list.Node{
		Data: []byte("A"),
	}
	second := &linked_list.Node{
		Data: []byte("B"),
	}
	third := &linked_list.Node{
		Data: []byte("C"),
	}
	head.Next = second
	second.Next = third

	return head
}
