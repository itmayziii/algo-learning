package linked_list_test

import (
	"fmt"
	"github.com/itmayziii/fm_last_algo_course/pkg/linked_list"
	"testing"
)

func TestDoubly_Traverse(t *testing.T) {
	tests := []struct {
		list     *linked_list.Doubly
		expected string
	}{
		{newDoubly("A", "B", "C"), "A -> B -> C"},
		{newDoubly("B", "C"), "B -> C"},
		{newDoubly("C"), "C"},
		{newDoubly("A", "B", "C", "X", "Y", "Z"), "A -> B -> C -> X -> Y -> Z"},
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

func TestDoubly_Search(t *testing.T) {
	tests := []struct {
		list     *linked_list.Doubly
		needle   string
		expected bool
	}{
		{newDoubly("A", "B", "C"), "A", true},
		{newDoubly("A", "B", "C"), "B", true},
		{newDoubly("A", "B", "C"), "C", true},
		{newDoubly("A", "B", "C"), "D", false},
		{newDoubly("A", "B", "C"), "Z", false},
		{newDoubly(), "A", false},
		{newDoubly(), "B", false},
		{newDoubly(), "C", false},
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

func TestDoubly_Length(t *testing.T) {
	tests := []struct {
		list     *linked_list.Doubly
		expected int
	}{
		{newDoubly(), 0},
		{newDoubly("A"), 1},
		{newDoubly("A", "B"), 2},
		{newDoubly("A", "B", "B"), 3},
		{newDoubly("A", "B", "B", "B"), 4},
		{newDoubly("A", "B", "B", "B", "Z"), 5},
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

func TestDoubly_Unshift(t *testing.T) {
	tests := []struct {
		list     *linked_list.Doubly
		value    string
		expected string
	}{
		{newDoubly("A", "B", "C"), "Z", "Z -> A -> B -> C"},
		{newDoubly("B", "C"), "A", "A -> B -> C"},
		{newDoubly("C"), "B", "B -> C"},
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

func TestDoubly_Shift(t *testing.T) {
	type expected struct {
		value     string
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Doubly
		expected expected
	}{
		{newDoubly("A", "B", "C"), expected{"A", "B -> C", 2}},
		{newDoubly("B", "C"), expected{"B", "C", 1}},
		{newDoubly("C"), expected{"C", "", 0}},
		{newDoubly(), expected{"", "", 0}},
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

func TestDoubly_Push(t *testing.T) {
	type expected struct {
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Doubly
		newTail  string
		expected expected
	}{
		{newDoubly("A", "B", "C"), "Z", expected{"A -> B -> C -> Z", 4}},
		{newDoubly("B", "C"), "Z", expected{"B -> C -> Z", 3}},
		{newDoubly("C"), "Z", expected{"C -> Z", 2}},
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

func TestDoubly_Pop(t *testing.T) {
	type expected struct {
		value     string
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Doubly
		expected expected
	}{
		{newDoubly("A", "B", "C"), expected{"C", "A -> B", 2}},
		{newDoubly("B", "C"), expected{"C", "B", 1}},
		{newDoubly("C"), expected{"C", "", 0}},
		{newDoubly("A", "B"), expected{"B", "A", 1}},
		{newDoubly(""), expected{"", "", 0}},
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

func TestDoubly_Insert(t *testing.T) {
	type expected struct {
		value     int
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Doubly
		value    string
		position int
		expected expected
	}{
		{newDoubly("A", "B", "C"), "AA", 0, expected{
			0,
			"AA -> A -> B -> C",
			4,
		}},
		{newDoubly("A", "B", "C"), "BB", 1, expected{
			1,
			"A -> BB -> B -> C",
			4,
		}},
		{newDoubly("A", "B", "C"), "CC", 2, expected{
			2,
			"A -> B -> CC -> C",
			4,
		}},
		{newDoubly("A", "B", "C"), "D", 3, expected{
			3,
			"A -> B -> C -> D",
			4,
		}},
		{newDoubly("A", "B", "C"), "E", 4, expected{
			-1,
			"A -> B -> C",
			3,
		}},
		{newDoubly("A", "B", "C"), "X", -3, expected{
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

func TestDoubly_Delete(t *testing.T) {
	type expected struct {
		value     int
		traversal string
		length    int
	}

	tests := []struct {
		list     *linked_list.Doubly
		position int
		expected expected
	}{
		{newDoubly("A", "B", "C"), 0, expected{0, "B -> C", 2}},
		{newDoubly("A", "B", "C"), 1, expected{1, "A -> C", 2}},
		{newDoubly("A", "B", "C"), 2, expected{2, "A -> B", 2}},
		{newDoubly("A", "B", "C"), 3, expected{-1, "A -> B -> C", 3}},
		{newDoubly("X", "Y", "Z"), 4, expected{-1, "X -> Y -> Z", 3}},
		{newDoubly(), 4, expected{-1, "", 0}},
		{newDoubly(), 0, expected{-1, "", 0}},
		{newDoubly("A"), 0, expected{0, "", 0}},
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

// I don't generally believe in tests like this but there is enough internal state being kept by the doubly
// linked list that it is important to test the result of operations after other operations have occurred.
func TestDoubly_Integration(t *testing.T) {
	t.Parallel()

	type expected struct {
		traversal string
		length    int
		first     string
		last      string
	}

	type change struct {
		operation func(ll *linked_list.Doubly)
		expected  expected
	}

	test := struct {
		list    *linked_list.Doubly
		changes []change
	}{
		newDoubly("A", "B", "C"),
		[]change{
			{
				operation: nil,
				expected: expected{
					traversal: "A -> B -> C",
					length:    3,
					first:     "A",
					last:      "C",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Shift()
				},
				expected: expected{
					traversal: "B -> C",
					length:    2,
					first:     "B",
					last:      "C",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Pop()
				},
				expected: expected{
					traversal: "B",
					length:    1,
					first:     "B",
					last:      "B",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Unshift("Z")
				},
				expected: expected{
					traversal: "Z -> B",
					length:    2,
					first:     "Z",
					last:      "B",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Delete(4)
				},
				expected: expected{
					traversal: "Z -> B",
					length:    2,
					first:     "Z",
					last:      "B",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Delete(1)
				},
				expected: expected{
					traversal: "Z",
					length:    1,
					first:     "Z",
					last:      "Z",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Pop()
				},
				expected: expected{
					traversal: "",
					length:    0,
					first:     "",
					last:      "",
				},
			},
			{
				operation: func(ll *linked_list.Doubly) {
					ll.Insert("X", 0)
				},
				expected: expected{
					traversal: "X",
					length:    1,
					first:     "X",
					last:      "X",
				},
			},
		},
	}

	for _, c := range test.changes {
		if c.operation != nil {
			c.operation(test.list)
		}

		actualTraversal := test.list.Traverse()
		if actualTraversal != c.expected.traversal {
			t.Errorf("actual traveral: %s, expected traversal: %s", actualTraversal, c.expected.traversal)
		}

		actualLength := test.list.Length()
		if actualLength != c.expected.length {
			t.Errorf("actual length: %d, expected length: %d", actualLength, c.expected.length)
		}

		actualFirst := test.list.First()
		if actualFirst != c.expected.first {
			t.Errorf("actual first: %s, expected first: %s", actualFirst, c.expected.first)
		}

		actualLast := test.list.Last()
		if actualLast != c.expected.last {
			t.Errorf("actual last: %s, expected last: %s", actualLast, c.expected.last)
		}
	}
}

// newDoubly returns a new linked list as some of the linked list methods are mutating, and we don't want
// to have shared state between test cases.
func newDoubly(values ...string) *linked_list.Doubly {
	ll := &linked_list.Doubly{}

	for _, value := range values {
		ll.Push(value)
	}

	return ll
}
