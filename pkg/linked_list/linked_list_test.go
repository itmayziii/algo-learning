package linked_list_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"testing"
)

func TestNode_Traverse(t *testing.T) {
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
		expected string
	}{
		{head, "A -> B -> C"},
		{second, "B -> C"},
		{third, "C"},
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
		{head, []byte("A"), true},
		{head, []byte("B"), true},
		{head, []byte("C"), true},
		{head, []byte("D"), false},
		{second, []byte("A"), false},
		{second, []byte("B"), true},
		{second, []byte("C"), true},
		{third, []byte("A"), false},
		{third, []byte("B"), false},
		{third, []byte("C"), true},
		{head, []byte("AA"), false},
		{head, []byte(""), false},
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
		{head, 3},
		{second, 2},
		{third, 1},
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

func TestNode_Unshift(t *testing.T) {
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

	newHead := &linked_list.Node{
		Data: []byte("0"),
	}

	tests := []struct {
		list    *linked_list.Node
		newHead *linked_list.Node
	}{
		{head, newHead},
		{second, newHead},
		{third, newHead},
		{nil, newHead},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Unshift(tt.newHead)
			if actual != tt.newHead {
				t.Errorf("returned node: %s is not the expected new node: %s", actual, tt.newHead)
			}
			if actual.Next != tt.list {
				t.Errorf("newly added first node: %s does not refrence the old head as the next node: %s",
					actual,
					tt.list,
				)
			}
		})
	}
}

func TestNode_Push(t *testing.T) {
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

	newTail := &linked_list.Node{
		Data: []byte("Z"),
	}

	tests := []struct {
		list    *linked_list.Node
		newTail *linked_list.Node
	}{
		{head, newTail},
		{second, newTail},
		{third, newTail},
		{nil, newTail},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Push(tt.newTail)
			if actual != tt.newTail {
				t.Errorf("returned node: %s, is not the expected new node: %s", actual, tt.newTail)
			}
			if actual.Next != nil {
				t.Errorf("new tail node: %s, is not suppose to have any nodes following it", tt.newTail)
			}

			previous := tt.list
			current := tt.list
			for current != nil {
				previous = current
				current = current.Next
			}
			if previous != nil && previous != actual {
				t.Errorf("last node in list: %s, is not the expected new tail node:", current)
			}
		})
	}
}
