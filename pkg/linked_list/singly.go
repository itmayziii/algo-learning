// Package linked_list implements different kinds of linked lists.
package linked_list

import (
	"strings"
)

// Singly represents a singly linked list.
type Singly struct {
	head   *singlyNode
	tail   *singlyNode
	length int
}

// singlyNode represents a node in a singly linked list.
type singlyNode struct {
	next  *singlyNode
	value string
}

// Traverse returns the contents of the linked list in a readable format. O(N)
// i.e. A linked list with the head value of "A" and the second singlyNode value of "B" would produce "A -> B"
func (ll *Singly) Traverse() string {
	current := ll.head
	var traversalStr strings.Builder
	for i := 0; current != nil; i++ {
		if i != 0 {
			traversalStr.WriteString(" -> ")
		}

		traversalStr.WriteString(current.value)
		current = current.next
	}

	return traversalStr.String()
}

// Search returns whether a given value is in the linked list. O(N)
func (ll *Singly) Search(needle string) bool {
	current := ll.head
	for current != nil {
		if current.value == needle {
			return true
		}
		current = current.next
	}

	return false
}

// Length returns how many nodes are in the linked list.
//
// O(1) - would be O(N) if we did not maintain the length internally and would instead have to walk the N sized
// list to count the values.
func (ll *Singly) Length() int {
	return ll.length
}

// Unshift adds a new node as the head of the linked list.
//
// O(1)
func (ll *Singly) Unshift(value string) {
	newNode := &singlyNode{value: value, next: ll.head}
	ll.head = newNode
	ll.length++
	if ll.head.next == nil {
		ll.tail = ll.head
	}
}

// Shift removes the head of a linked list and returns the value of the removed node.
//
// O(1)
func (ll *Singly) Shift() string {
	if ll.head == nil {
		return ""
	}

	oldHead := ll.head
	ll.head = oldHead.next
	ll.length--

	if ll.head != nil && ll.head.next == nil {
		ll.tail = ll.head
	}

	return oldHead.value
}

// Push adds a new node as the tail of the linked list.
//
// O(1) - It would be O(N) if we did not maintain a reference to the current tail node as we would have to loop
// through the N sized list to find the tail.
func (ll *Singly) Push(value string) {
	newNode := &singlyNode{value: value}

	if ll.tail != nil {
		ll.tail.next = newNode
	}
	ll.tail = newNode
	ll.length++

	if ll.head == nil {
		ll.head = newNode
	}
}

// Pop removes the tail of a linked list and returns its value.
//
// O(N) - In a doubly linked list this operation could be O(1) because we could jump to the tail and set the previous
// node.next = nil, but in a singly linked list we must walk the N sized list so that we can get penultimate node.
func (ll *Singly) Pop() string {
	if ll.tail == nil {
		return ""
	}

	if ll.head.next == nil {
		value := ll.head.value
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return value
	}

	previous := ll.head
	current := ll.head
	for current != nil {
		if current.next == nil {
			previous.next = nil
			ll.tail = previous
			ll.length--
			return current.value
		}

		previous = current
		current = current.next
	}

	return ""
}

// Insert inserts a new node at a given position and returns the position the value was inserted at or -1 if no
// insertion occurred.
//
// O(N)
func (ll *Singly) Insert(value string, position int) int {
	if position < 0 {
		return -1
	}

	if position == 0 {
		ll.Unshift(value)
		return 0
	}

	previous := ll.head
	current := ll.head
	for i := 0; i <= position; i++ {
		// We are 2 past the tail node, time to give up
		if previous == nil {
			return -1
		}

		if i != position {
			previous = current
			if current != nil {
				current = current.next
			}
			continue
		}

		newNode := &singlyNode{value: value, next: current}
		if current == nil {
			ll.tail = newNode
		}
		previous.next = newNode
		ll.length++
		return i
	}

	return -1
}

// Delete deletes a node and returns the position of the deleted node or -1 if no deletion occurred.
//
// O(N)
func (ll *Singly) Delete(position int) int {
	if position < 0 || ll.length == 0 {
		return -1
	}

	if position == 0 {
		ll.Shift()
		return 0
	}

	previous := ll.head
	current := ll.head
	for i := 0; i <= position && current != nil; i++ {
		if i != position {
			previous = current
			current = current.next
			continue
		}

		previous.next = current.next
		ll.length--
		if previous.next == nil {
			ll.tail = previous
		}
		return i
	}

	return -1
}

// First returns the value of the first node in the list.
//
// O(1)
func (ll *Singly) First() string {
	if ll.head == nil {
		return ""
	}

	return ll.head.value
}

// Last returns the value of the last node in the list.
//
// O(1)
func (ll *Singly) Last() string {
	if ll.tail == nil {
		return ""
	}

	return ll.tail.value
}

// String implements the Stringer interface
func (ll *Singly) String() string {
	return ll.Traverse()
}

// String implements the Stringer interface
func (n *singlyNode) String() string {
	return n.value
}
