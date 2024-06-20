package tree

func PostOrderSearch(root *BinaryNode) []int {
	return postOrderWalk(root, []int{})
}

func postOrderWalk(curr *BinaryNode, path []int) []int {
	if curr == nil {
		return path
	}

	// pre

	// recurse
	path = postOrderWalk(curr.left, path)
	path = postOrderWalk(curr.right, path)
	path = append(path, curr.value)

	// post
	return path
}
