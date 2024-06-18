// Package linked_list implements different kinds of linked lists.
package linked_list

import (
	"strings"
)

// Doubly represents a doubly linked list.
type Doubly struct {
	head   *doublyNode
	tail   *doublyNode
	length int
}

// doublyNode represents a node in a doubly linked list.
type doublyNode struct {
	next  *doublyNode
	prev  *doublyNode
	value string
}

// Traverse returns the contents of the linked list in a readable format. O(N)
// i.e. A linked list with the head value of "A" and the second singlyNode value of "B" would produce "A -> B"
func (ll *Doubly) Traverse() string {
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
func (ll *Doubly) Search(needle string) bool {
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
func (ll *Doubly) Length() int {
	return ll.length
}

// Unshift adds a new node as the head of the linked list.
//
// O(1)
func (ll *Doubly) Unshift(value string) {
	newNode := &doublyNode{value: value, next: ll.head}

	if ll.head != nil {
		ll.head.prev = newNode
	}

	ll.head = newNode
	ll.length++
	if ll.head.next == nil {
		ll.tail = ll.head
	}
}

// Shift removes the head of a linked list and returns the value of the removed node.
//
// O(1)
func (ll *Doubly) Shift() string {
	if ll.head == nil {
		return ""
	}

	oldHead := ll.head
	ll.head = oldHead.next
	ll.length--

	// free memory
	oldHead.next = nil
	oldHead.prev = nil

	if ll.head != nil {
		ll.head.prev = nil

		if ll.head.next == nil {
			ll.tail = ll.head
		}
	}

	return oldHead.value
}

// Push adds a new node as the tail of the linked list.
//
// O(1) - It would be O(N) if we did not maintain a reference to the current tail node as we would have to loop
// through the N sized list to find the tail.
func (ll *Doubly) Push(value string) {
	newNode := &doublyNode{value: value, prev: ll.tail}

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
// O(1) - In a singly linked list this operation would be O(N) because we could not jump to the tail and set
// the previous.next = nil
func (ll *Doubly) Pop() string {
	if ll.tail == nil {
		return ""
	}

	if ll.head.next == nil || ll.tail.prev == nil {
		value := ll.head.value
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return value
	}

	oldTail := ll.tail
	ll.tail = oldTail.prev
	ll.tail.next = nil
	ll.length--

	// free memory
	oldTail.next = nil
	oldTail.prev = nil

	return oldTail.value
}

// Insert inserts a new node at a given position and returns the position the value was inserted at or -1 if no
// insertion occurred.
//
// O(N)
func (ll *Doubly) Insert(value string, position int) int {
	if position < 0 || position > ll.length {
		return -1
	}

	if position == 0 {
		ll.Unshift(value)
		return 0
	}

	if position == ll.length {
		ll.Push(value)
		return position
	}

	current := ll.head
	for i := 0; i <= position && current != nil; i++ {
		if i != position {
			current = current.next
			continue
		}

		newNode := &doublyNode{value: value, next: current, prev: current.prev}
		current.prev.next = newNode
		current.prev = newNode

		ll.length++
		return i
	}

	return -1
}

// Delete deletes a node and returns the position of the deleted node or -1 if no deletion occurred.
//
// O(N)
func (ll *Doubly) Delete(position int) int {
	if position < 0 || ll.length == 0 || position >= ll.length {
		return -1
	}

	if position == 0 {
		ll.Shift()
		return 0
	}

	if position == ll.length-1 {
		ll.Pop()
		return position
	}

	current := ll.head
	for i := 0; i <= position && current != nil; i++ {
		if i != position {
			current = current.next
			continue
		}

		current.prev.next = current.next
		if current.next != nil {
			current.next.prev = current.prev
		}
		ll.length--

		// free memory
		current.next = nil
		current.prev = nil

		return i
	}

	return -1
}

// First returns the value of the first node in the list.
//
// O(1)
func (ll *Doubly) First() string {
	if ll.head == nil {
		return ""
	}

	return ll.head.value
}

// Last returns the value of the last node in the list.
//
// O(1)
func (ll *Doubly) Last() string {
	if ll.tail == nil {
		return ""
	}

	return ll.tail.value
}

// String implements the Stringer interface
func (ll *Doubly) String() string {
	return ll.Traverse()
}

// String implements the Stringer interface
func (n *doublyNode) String() string {
	return n.value
}
