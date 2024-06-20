package tree

func PreOrderSearch(root *BinaryNode) []int {
	return preOrderWalk(root, []int{})
}

func preOrderWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// pre
	path = append(path, curr.value)

	// recurse
	path = preOrderWalk(curr.left, path)
	path = preOrderWalk(curr.right, path)

	// post
	return path
}
