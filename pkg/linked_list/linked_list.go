package linked_list

import (
	"bytes"
	"fmt"
	"strings"
)

type Node struct {
	Next *Node
	Data []byte
}

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
