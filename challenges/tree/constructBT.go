// Given preorder and inorder traversal of a tree, construct the binary tree.
package tree

import "github.com/abdalla/go_dsa/tree"

func helper(preorder []int, inorder []int, current *int, low, high int) *tree.Tree {
	if low > high {
		return nil
	}

	if low == high {
		root := &tree.Tree{Value: preorder[*current]}
		(*current)++
		return root
	}

	root := &tree.Tree{Value: preorder[*current]}
	rootValue := preorder[*current]
	(*current)++

	var rootIndex int
	for i := low; i <= high; i++ {
		if inorder[i] == rootValue {
			rootIndex = i
		}
	}

	root.Left = helper(preorder, inorder, current, low, rootIndex-1)
	root.Right = helper(preorder, inorder, current, rootIndex+1, high)

	return root
}

func BuildTree(preorder []int, inorder []int) *tree.Tree {

	lenOfTree := len(preorder)

	current := 0
	newRoot := helper(preorder, inorder, &current, 0, lenOfTree-1)

	return newRoot
}
