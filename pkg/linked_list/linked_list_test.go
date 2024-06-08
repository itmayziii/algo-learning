package linked_list_test

import (
	"fmt"
	"github.com/itmayziii/last_algo_course/pkg/linked_list"
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
		{
			head,
			"A -> B -> C",
		},
		{
			second,
			"B -> C",
		},
		{
			third,
			"C",
		},
		{
			nil,
			"",
		},
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
		{
			head,
			[]byte("A"),
			true,
		},
		{
			head,
			[]byte("B"),
			true,
		},
		{
			head,
			[]byte("C"),
			true,
		},
		{
			head,
			[]byte("D"),
			false,
		},
		{
			second,
			[]byte("A"),
			false,
		},
		{
			second,
			[]byte("B"),
			true,
		},
		{
			second,
			[]byte("C"),
			true,
		},
		{
			third,
			[]byte("A"),
			false,
		},
		{
			third,
			[]byte("B"),
			false,
		},
		{
			third,
			[]byte("C"),
			true,
		},
		{
			head,
			[]byte("AA"),
			false,
		},
		{
			head,
			[]byte(""),
			false,
		},
		{
			nil,
			[]byte("A"),
			false,
		},
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
