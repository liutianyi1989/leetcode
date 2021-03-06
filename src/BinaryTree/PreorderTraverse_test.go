package BinaryTree

import "testing"

func TestPreorderTraverse(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(preorderTraversal(root))
}

func TestPreorderTraverseIter(t *testing.T) {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}

	root.Left.Left = &TreeNode{Val:4}

	t.Log(preorderTraversalIter(root))
}
