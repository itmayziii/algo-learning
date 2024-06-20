package tree

type BinaryNode struct {
	value int
	left  *BinaryNode
	right *BinaryNode
}

func NewBinaryNode(value int, left, right *BinaryNode) *BinaryNode {
	return &BinaryNode{value, left, right}
}
