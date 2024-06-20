package tree

func InOrderSearch(root *BinaryNode) []int {
	return inOrderWalk(root, []int{})
}

func inOrderWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// pre

	// recurse
	path = inOrderWalk(curr.left, path)
	path = append(path, curr.value)
	path = inOrderWalk(curr.right, path)

	// post
	return path
}
