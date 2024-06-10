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

func TestSingly_Search(t *testing.T) {
	tests := []struct {
		list     *linked_list.Singly
		needle   string
		expected bool
	}{
		{newLinkedList("A", "B", "C"), "A", true},
		{newLinkedList("A", "B", "C"), "B", true},
		{newLinkedList("A", "B", "C"), "C", true},
		{newLinkedList("A", "B", "C"), "D", false},
		{newLinkedList("A", "B", "C"), "Z", false},
		{newLinkedList(), "A", false},
		{newLinkedList(), "B", false},
		{newLinkedList(), "C", false},
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

func TestSingly_Length(t *testing.T) {
	tests := []struct {
		list     *linked_list.Singly
		expected int
	}{
		{newLinkedList(), 0},
		{newLinkedList("A"), 1},
		{newLinkedList("A", "B"), 2},
		{newLinkedList("A", "B", "B"), 3},
		{newLinkedList("A", "B", "B", "B"), 4},
		{newLinkedList("A", "B", "B", "B", "Z"), 5},
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

func TestSingly_Unshift(t *testing.T) {
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

func TestSingly_Shift(t *testing.T) {
	type expected struct {
		value     string
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Singly
		expected expected
	}{
		{newLinkedList("A", "B", "C"), expected{"A", "B -> C", 2}},
		{newLinkedList("B", "C"), expected{"B", "C", 1}},
		{newLinkedList("C"), expected{"C", "", 0}},
		{newLinkedList(), expected{"", "", 0}},
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

			actualLength := tt.list.Length()
			if actualLength != tt.expected.length {
				t.Errorf("actual length: %d, expected length %d", actualLength, tt.expected.length)
			}
		})
	}
}

func TestSingly_Push(t *testing.T) {
	type expected struct {
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Singly
		newTail  string
		expected expected
	}{
		{newLinkedList("A", "B", "C"), "Z", expected{"A -> B -> C -> Z", 4}},
		{newLinkedList("B", "C"), "Z", expected{"B -> C -> Z", 3}},
		{newLinkedList("C"), "Z", expected{"C -> Z", 2}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			tt.list.Push(tt.newTail)
			actual := tt.list.Traverse()
			if actual != tt.expected.traversal {
				t.Errorf("actual: %s, expected %s", actual, tt.expected.traversal)
			}

			actualLength := tt.list.Length()
			if actualLength != tt.expected.length {
				t.Errorf("actual length: %d, expected length %d", actualLength, tt.expected.length)
			}
		})
	}
}

func TestSingly_Pop(t *testing.T) {
	type expected struct {
		value     string
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Singly
		expected expected
	}{
		{newLinkedList("A", "B", "C"), expected{"C", "A -> B", 2}},
		{newLinkedList("B", "C"), expected{"C", "B", 1}},
		{newLinkedList("C"), expected{"C", "", 0}},
		{newLinkedList("A", "B"), expected{"B", "A", 1}},
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

			actualLength := tt.list.Length()
			if actualLength != tt.expected.length {
				t.Errorf("actual length: %d, expected length %d", actualLength, tt.expected.length)
			}
		})
	}
}

func TestSingly_Insert(t *testing.T) {
	type expected struct {
		value     int
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Singly
		value    string
		position int
		expected expected
	}{
		{newLinkedList("A", "B", "C"), "AA", 0, expected{
			0,
			"AA -> A -> B -> C",
			4,
		}},
		{newLinkedList("A", "B", "C"), "BB", 1, expected{
			1,
			"A -> BB -> B -> C",
			4,
		}},
		{newLinkedList("A", "B", "C"), "CC", 2, expected{
			2,
			"A -> B -> CC -> C",
			4,
		}},
		{newLinkedList("A", "B", "C"), "D", 3, expected{
			3,
			"A -> B -> C -> D",
			4,
		}},
		{newLinkedList("A", "B", "C"), "E", 4, expected{
			-1,
			"A -> B -> C",
			3,
		}},
		{newLinkedList("A", "B", "C"), "X", -3, expected{
			-1,
			"A -> B -> C",
			3,
		}},
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

			actualLength := tt.list.Length()
			if actualLength != tt.expected.length {
				t.Errorf("actual length: %d, expected length %d", actualLength, tt.expected.length)
			}
		})
	}
}

func TestSingly_Delete(t *testing.T) {
	type expected struct {
		value     int
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Singly
		position int
		expected expected
	}{
		{newLinkedList("A", "B", "C"), 0, expected{0, "B -> C", 2}},
		{newLinkedList("A", "B", "C"), 1, expected{1, "A -> C", 2}},
		{newLinkedList("A", "B", "C"), 2, expected{2, "A -> B", 2}},
		{newLinkedList("A", "B", "C"), 3, expected{-1, "A -> B -> C", 3}},
		{newLinkedList("X", "Y", "Z"), 4, expected{-1, "X -> Y -> Z", 3}},
		{newLinkedList(), 4, expected{-1, "", 0}},
		{newLinkedList(), 0, expected{-1, "", 0}},
		{newLinkedList("A"), 0, expected{0, "", 0}},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("test case %d", i), func(t *testing.T) {
			t.Parallel()

			actual := tt.list.Delete(tt.position)
			if actual != tt.expected.value {
				t.Errorf("actual: %d, expected %d", actual, tt.expected.value)
			}

			actualTraversal := tt.list.Traverse()
			if actualTraversal != tt.expected.traversal {
				t.Errorf("actual traversal: %s, expected traversal %s", actualTraversal, tt.expected.traversal)
			}

			actualLength := tt.list.Length()
			if actualLength != tt.expected.length {
				t.Errorf("actual length: %d, expected length %d", actualLength, tt.expected.length)
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
