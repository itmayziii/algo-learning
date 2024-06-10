package linked_list

import (
	"bytes"
	"fmt"
	"strings"
)

type OutOfRangeError struct {
	position int
}

func (e OutOfRangeError) Error() string {
	return fmt.Sprintf("position: %d is not in range of linked list", e.position)
}

type Node struct {
	Next *Node
	Data []byte
}

// Traverse returns the contents of the linked list in a readable format. O(N)
// i.e. A linked list with the head value of "A" and the second node value of "B" would produce "A -> B"
func (n *Node) Traverse() string {
	current := n
	var traversalStr strings.Builder
	for i := 0; current != nil; i++ {
		if i != 0 {
			traversalStr.WriteString(fmt.Sprintf(" -> "))
		}
		traversalStr.WriteString(fmt.Sprintf("%s", current.Data))
		current = current.Next
	}

	return traversalStr.String()
}

// Search returns whether a given value is in the linked list. O(N)
func (n *Node) Search(needle []byte) bool {
	current := n
	for current != nil {
		if bytes.Equal(current.Data, needle) {
			return true
		}
		current = current.Next
	}

	return false
}

// Length returns how many nodes are in the linked list. O(N)
func (n *Node) Length() int {
	current := n
	count := 0
	for current != nil {
		count++
		current = current.Next
	}

	return count
}

// Shift removes the head of a linked list. O(1)
// Returns the head node.
func (n *Node) Shift() *Node {
	if n == nil {
		return nil
	}

	return n.Next
}

// Unshift adds a new node as the head of the linked list and returns the new head. O(1)
// Returns the head node.
func (n *Node) Unshift(newNode *Node) *Node {
	newNode.Next = n
	return newNode
}

// Push adds a new node as the tail of the linked list. O(N)
// Returns the head node.
func (n *Node) Push(newNode *Node) *Node {
	if n == nil {
		return newNode
	}

	previous := n
	current := n
	for current != nil {
		previous = current
		current = current.Next
	}

	previous.Next = newNode

	return n
}

// Pop removes the tail of a linked list. O(N)
// Returns the head node.
func (n *Node) Pop() *Node {
	if n == nil || n.Next == nil {
		return nil
	}

	previous := n
	current := n
	for current != nil {
		if current.Next == nil {
			previous.Next = nil
			return n
		}

		previous = current
		current = current.Next
	}

	return n
}

// Insert inserts a new node at a given position. O(N)
// Return the head node.
func (n *Node) Insert(data []byte, position int) (*Node, error) {
	if position < 0 {
		return nil, OutOfRangeError{position}
	}

	newNode := &Node{Data: data}
	if position == 0 || n == nil {
		return n.Unshift(newNode), nil
	}

	previous := n
	current := n
	for i := 0; i <= position; i++ {
		if i != position {
			if current == nil {
				return nil, OutOfRangeError{position}
			}

			previous = current
			current = current.Next
			continue
		}

		newNode.Next = current
		previous.Next = newNode
		break
	}

	return n, nil
}

// Delete deletes a node at a given position. O(N)
// Returns the head node
func (n *Node) Delete(position int) (*Node, error) {
	if position < 0 {
		return nil, OutOfRangeError{position}
	}

	if position == 0 || n == nil {
		return n.Shift(), nil
	}

	previous := n
	current := n
	for i := 0; i <= position; i++ {
		if i != position {
			if current == nil {
				return nil, OutOfRangeError{position}
			}

			previous = current
			current = current.Next
			continue
		}

		previous.Next = current.Next
		break
	}
	
	return n, nil
}

// String implements the Stringer interface
func (n *Node) String() string {
	return n.Traverse()
}
