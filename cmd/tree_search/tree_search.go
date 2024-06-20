package main

import (
	"github.com/itmayziii/fm_last_algo_course/pkg/tree"
	"log"
)

func main() {
	binaryTree := tree.NewBinaryNode(2,
		tree.NewBinaryNode(7,
			tree.NewBinaryNode(2, nil, nil),
			tree.NewBinaryNode(6,
				tree.NewBinaryNode(5, nil, nil),
				tree.NewBinaryNode(11, nil, nil),
			),
		),
		tree.NewBinaryNode(5,
			nil,
			tree.NewBinaryNode(9,
				nil,
				tree.NewBinaryNode(4, nil, nil),
			),
		),
	)

	preOrder := tree.PreOrderSearch(binaryTree)
	log.Printf("pre order search %v", preOrder)

	inOrder := tree.InOrderSearch(binaryTree)
	log.Printf("in order search %v", inOrder)

	postOrder := tree.PostOrderSearch(binaryTree)
	log.Printf("post order search %v", postOrder)
}
