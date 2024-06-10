package linked_list_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"testing"
)

func TestLinkedList_Traverse(t *testing.T) {
	tests := []struct {
		list     *linked_list.Singly
		expected string
	}{
		{newLinkedList("A", "B", "C"), "A -> B -> C"},
		{newLinkedList("B", "C"), "B -> C"},
		{newLinkedList("C"), "C"},
		{newLinkedList("A", "B", "C", "X", "Y", "Z"), "A -> B -> C -> X -> Y -> Z"},
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

func TestNode_Unshift(t *testing.T) {
	tests := []struct {
		list     *linked_list.Singly
		value    string
		expected string
	}{
		{newLinkedList("A", "B", "C"), "Z", "Z -> A -> B -> C"},
		{newLinkedList("B", "C"), "A", "A -> B -> C"},
		{newLinkedList("C"), "B", "B -> C"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			tt.list.Unshift(tt.value)
			actual := tt.list.Traverse()
			if actual != tt.expected {
				t.Errorf("actual: %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Shift(t *testing.T) {
	type expected struct {
		value     string
		traversal string
	}

	tests := []struct {
		list     *linked_list.Singly
		expected expected
	}{
		{newLinkedList("A", "B", "C"), expected{"A", "B -> C"}},
		{newLinkedList("B", "C"), expected{"B", "C"}},
		{newLinkedList("C"), expected{"C", ""}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Shift()
			if actual != tt.expected.value {
				t.Errorf("actual: %s, expected %s", actual, tt.expected.value)
			}

			actualTraversal := tt.list.Traverse()
			if actualTraversal != tt.expected.traversal {
				t.Errorf("actualTraversal: %s, expectedTraversal %s", actualTraversal, tt.expected.traversal)
			}
		})
	}
}

func TestNode_Push(t *testing.T) {
	tests := []struct {
		list     *linked_list.Singly
		newTail  string
		expected string
	}{
		{newLinkedList("A", "B", "C"), "Z", "A -> B -> C -> Z"},
		{newLinkedList("B", "C"), "Z", "B -> C -> Z"},
		{newLinkedList("C"), "Z", "C -> Z"},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			tt.list.Push(tt.newTail)
			actual := tt.list.Traverse()
			if actual != tt.expected {
				t.Errorf("actual: %s, expected %s", actual, tt.expected)
			}
		})
	}
}

func TestNode_Pop(t *testing.T) {
	type expected struct {
		value     string
		traversal string
	}

	tests := []struct {
		list     *linked_list.Singly
		expected expected
	}{
		{newLinkedList("A", "B", "C"), expected{"C", "A -> B"}},
		{newLinkedList("B", "C"), expected{"C", "B"}},
		{newLinkedList("C"), expected{"C", ""}},
		{newLinkedList("A", "B"), expected{"B", "A"}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Pop()
			if actual != tt.expected.value {
				t.Errorf("actual: %s, expected %s", actual, tt.expected.value)
			}
			actualTraversal := tt.list.Traverse()
			if actualTraversal != tt.expected.traversal {
				t.Errorf("actual traversal: %s, expected traversal %s", actualTraversal, tt.expected.traversal)
			}
		})
	}
}

func TestNode_Insert(t *testing.T) {
	type expected struct {
		value     int
		traversal string
	}

	tests := []struct {
		list     *linked_list.Singly
		value    string
		position int
		expected expected
	}{
		{newLinkedList("A", "B", "C"), "AA", 0, expected{0, "AA -> A -> B -> C"}},
		{newLinkedList("A", "B", "C"), "BB", 1, expected{1, "A -> BB -> B -> C"}},
		{newLinkedList("A", "B", "C"), "CC", 2, expected{2, "A -> B -> CC -> C"}},
		{newLinkedList("A", "B", "C"), "D", 3, expected{3, "A -> B -> C -> D"}},
		{newLinkedList("A", "B", "C"), "E", 4, expected{-1, "A -> B -> C"}},
		{newLinkedList("A", "B", "C"), "X", -3, expected{-1, "A -> B -> C"}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Insert(tt.value, tt.position)
			if actual != tt.expected.value {
				t.Errorf("actual: %d, expected %d", actual, tt.expected.value)
			}
			actualTraversal := tt.list.Traverse()
			if actualTraversal != tt.expected.traversal {
				t.Errorf("actual traversal: %s, expected traversal %s", actualTraversal, tt.expected.traversal)
			}
		})
	}
}

//func TestNode_Search(t *testing.T) {
//	head := &singly_linked_list.Node{
//		Data: []byte("A"),
//	}
//	second := &singly_linked_list.Node{
//		Data: []byte("B"),
//	}
//	third := &singly_linked_list.Node{
//		Data: []byte("C"),
//	}
//	head.Next = second
//	second.Next = third
//
//	tests := []struct {
//		list     *singly_linked_list.Node
//		needle   []byte
//		expected bool
//	}{
//		{newLinkedList(), []byte("A"), true},
//		{newLinkedList(), []byte("B"), true},
//		{newLinkedList(), []byte("C"), true},
//		{newLinkedList(), []byte("D"), false},
//		{newLinkedList().Next, []byte("A"), false},
//		{newLinkedList().Next, []byte("B"), true},
//		{newLinkedList().Next, []byte("C"), true},
//		{newLinkedList().Next.Next, []byte("A"), false},
//		{newLinkedList().Next.Next, []byte("B"), false},
//		{newLinkedList().Next.Next, []byte("C"), true},
//		{newLinkedList(), []byte("AA"), false},
//		{newLinkedList(), []byte(""), false},
//		{nil, []byte("A"), false},
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
//			t.Parallel()
//
//			actual := tt.list.Search(tt.needle)
//			if actual != tt.expected {
//				t.Errorf("actual %t, expected %t", actual, tt.expected)
//			}
//		})
//	}
//}
//
//func TestNode_Length(t *testing.T) {
//	head := &singly_linked_list.Node{
//		Data: []byte("A"),
//	}
//	second := &singly_linked_list.Node{
//		Data: []byte("B"),
//	}
//	third := &singly_linked_list.Node{
//		Data: []byte("C"),
//	}
//	head.Next = second
//	second.Next = third
//
//	tests := []struct {
//		list     *singly_linked_list.Node
//		expected int
//	}{
//		{newLinkedList(), 3},
//		{newLinkedList().Next, 2},
//		{newLinkedList().Next.Next, 1},
//		{nil, 0},
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
//			t.Parallel()
//
//			actual := tt.list.Length()
//			if actual != tt.expected {
//				t.Errorf("actual %d, expected %d", actual, tt.expected)
//			}
//		})
//	}
//}
//
//func TestNode_Shift(t *testing.T) {
//	tests := []struct {
//		list     *singly_linked_list.Node
//		expected string
//	}{
//		{newLinkedList(), "B -> C"},
//		{newLinkedList().Next, "C"},
//		{newLinkedList().Next.Next, ""},
//		{nil, ""},
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
//			t.Parallel()
//
//			actual := tt.list.Shift().Traverse()
//			if actual != tt.expected {
//				t.Errorf("actual: %s, expected %s", actual, tt.expected)
//			}
//		})
//	}
//}

//

//
//func TestNode_Delete(t *testing.T) {
//	type expected struct {
//		value string
//		err   error
//	}
//
//	tests := []struct {
//		list     *singly_linked_list.Node
//		position int
//		expected expected
//	}{
//		//{newLinkedList(), 0, expected{"B -> C", nil}},
//		{newLinkedList(), 1, expected{"A -> C", nil}},
//		//{newLinkedList(), 2, expected{"A -> B", nil}},
//		//{newLinkedList(), 3, expected{"", linked_list.OutOfRangeError{}}},
//		//{newLinkedList(), 4, expected{"", linked_list.OutOfRangeError{}}},
//	}
//
//	for i, tt := range tests {
//		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
//			t.Parallel()
//
//			head, err := tt.list.Delete(tt.position)
//			if tt.expected.err == nil {
//				if err != nil {
//					t.Errorf("unexpected error %v", err)
//				}
//
//				actual := head.Traverse()
//				if actual != tt.expected.value {
//					t.Errorf("actual: %s, expected %s", actual, tt.expected.value)
//				}
//			} else {
//				if err != nil && errors.Is(err, tt.expected.err) {
//					t.Errorf("actual error: %v, expected error: %v", err, tt.expected.err)
//				}
//			}
//		})
//	}
//}

// newLinkedList returns a new linked list as some of the linked list methods are mutating, and we don't want
// to have shared state between test cases.
func newLinkedList(values ...string) *linked_list.Singly {
	ll := &linked_list.Singly{}

	for _, value := range values {
		ll.Push(value)
	}

	return ll
}
